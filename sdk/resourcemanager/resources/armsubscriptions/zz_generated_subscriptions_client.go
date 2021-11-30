//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsubscriptions

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SubscriptionsClient contains the methods for the Subscriptions group.
// Don't use this type directly, use NewSubscriptionsClient() instead.
type SubscriptionsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewSubscriptionsClient creates a new instance of SubscriptionsClient with the specified values.
func NewSubscriptionsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *SubscriptionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SubscriptionsClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets details about a specified subscription.
// If the operation fails it returns the *CloudError error type.
func (client *SubscriptionsClient) Get(ctx context.Context, subscriptionID string, options *SubscriptionsGetOptions) (SubscriptionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return SubscriptionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SubscriptionsClient) getCreateRequest(ctx context.Context, subscriptionID string, options *SubscriptionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SubscriptionsClient) getHandleResponse(resp *http.Response) (SubscriptionsGetResponse, error) {
	result := SubscriptionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Subscription); err != nil {
		return SubscriptionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SubscriptionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets all subscriptions for a tenant.
// If the operation fails it returns the *CloudError error type.
func (client *SubscriptionsClient) List(options *SubscriptionsListOptions) *SubscriptionsListPager {
	return &SubscriptionsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SubscriptionsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SubscriptionListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SubscriptionsClient) listCreateRequest(ctx context.Context, options *SubscriptionsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SubscriptionsClient) listHandleResponse(resp *http.Response) (SubscriptionsListResponse, error) {
	result := SubscriptionsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SubscriptionListResult); err != nil {
		return SubscriptionsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SubscriptionsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListLocations - This operation provides all the locations that are available for resource providers; however, each resource provider may support a subset
// of this list.
// If the operation fails it returns the *CloudError error type.
func (client *SubscriptionsClient) ListLocations(ctx context.Context, subscriptionID string, options *SubscriptionsListLocationsOptions) (SubscriptionsListLocationsResponse, error) {
	req, err := client.listLocationsCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return SubscriptionsListLocationsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionsListLocationsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionsListLocationsResponse{}, client.listLocationsHandleError(resp)
	}
	return client.listLocationsHandleResponse(resp)
}

// listLocationsCreateRequest creates the ListLocations request.
func (client *SubscriptionsClient) listLocationsCreateRequest(ctx context.Context, subscriptionID string, options *SubscriptionsListLocationsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/locations"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	if options != nil && options.IncludeExtendedLocations != nil {
		reqQP.Set("includeExtendedLocations", strconv.FormatBool(*options.IncludeExtendedLocations))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listLocationsHandleResponse handles the ListLocations response.
func (client *SubscriptionsClient) listLocationsHandleResponse(resp *http.Response) (SubscriptionsListLocationsResponse, error) {
	result := SubscriptionsListLocationsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LocationListResult); err != nil {
		return SubscriptionsListLocationsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listLocationsHandleError handles the ListLocations error response.
func (client *SubscriptionsClient) listLocationsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}