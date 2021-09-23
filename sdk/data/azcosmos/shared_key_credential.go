// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync/atomic"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// NewSharedKeyCredential creates an immutable SharedKeyCredential containing the
// account's primary or secondary key.
func NewSharedKeyCredential(accountKey string) (*SharedKeyCredential, error) {
	c := SharedKeyCredential{}
	if err := c.SetAccountKey(accountKey); err != nil {
		return nil, err
	}
	return &c, nil
}

// SharedKeyCredential contains an account's name and its primary or secondary key.
// It is immutable making it shareable and goroutine-safe.
type SharedKeyCredential struct {
	// Only the NewSharedKeyCredential method should set these; all other methods should treat them as read-only
	accountKey atomic.Value // []byte
}

// SetAccountKey replaces the existing account key with the specified account key.
func (c *SharedKeyCredential) SetAccountKey(accountKey string) error {
	bytes, err := base64.StdEncoding.DecodeString(accountKey)
	if err != nil {
		return fmt.Errorf("decode account key: %w", err)
	}
	c.accountKey.Store(bytes)
	return nil
}

// computeHMACSHA256 generates a hash signature for an HTTP request
func (c *SharedKeyCredential) computeHMACSHA256(s string) (base64String string) {
	h := hmac.New(sha256.New, c.accountKey.Load().([]byte))
	_, _ = h.Write([]byte(s))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *SharedKeyCredential) buildCanonicalizedAuthHeaderFromRequest(req *azcore.Request) (string, error) {
	var opValues cosmosOperationContext
	value := ""

	if req.OperationValue(&opValues) {
		resourceTypePath, err := getResourcePath(opValues.resourceType)

		if err != nil {
			return "", err
		}

		resourceAddress := opValues.resourceAddress
		if opValues.isRidBased {
			resourceAddress = strings.ToLower(resourceAddress)
		}

		value = c.buildCanonicalizedAuthHeader(req.Method, resourceTypePath, resourceAddress, req.Request.Header.Get(azcore.HeaderXmsDate), "master", "1.0")
	}

	return value, nil
}

//where date is like time.RFC1123 but hard-codes GMT as the time zone
func (c *SharedKeyCredential) buildCanonicalizedAuthHeader(method, resourceType, resourceAddress, xmsDate, tokenType, version string) string {
	if method == "" || resourceType == "" {
		return ""
	}

	// https://docs.microsoft.com/en-us/rest/api/cosmos-db/access-control-on-cosmosdb-resources#constructkeytoken
	stringToSign := join(strings.ToLower(method), "\n", strings.ToLower(resourceType), "\n", resourceAddress, "\n", strings.ToLower(xmsDate), "\n", "", "\n")
	signature := c.computeHMACSHA256(stringToSign)

	return url.QueryEscape(join("type=" + tokenType + "&ver=" + version + "&sig=" + signature))
}

// AuthenticationPolicy implements the Credential interface on SharedKeyCredential.
func (c *SharedKeyCredential) AuthenticationPolicy(azcore.AuthenticationPolicyOptions) azcore.Policy {
	return azcore.PolicyFunc(func(req *azcore.Request) (*azcore.Response, error) {

		// Add a x-ms-date header if it doesn't already exist
		if d := req.Request.Header.Get(azcore.HeaderXmsDate); d == "" {
			req.Request.Header.Set(azcore.HeaderXmsDate, time.Now().UTC().Format(http.TimeFormat))
		}

		authHeader, err := c.buildCanonicalizedAuthHeaderFromRequest(req)
		if err != nil {
			return nil, err
		}

		if authHeader != "" {
			req.Request.Header.Set(azcore.HeaderAuthorization, authHeader)
		}

		response, err := req.Next()
		if err != nil && response != nil && response.StatusCode == http.StatusForbidden {
			// Service failed to authenticate request, log it
			azcore.Log().Write(azcore.LogResponse, "===== HTTP Forbidden status, Authorization:\n"+authHeader+"\n=====\n")
		}
		return response, err
	})
}

func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}