//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateLinkServiceResourceOperationResultsClient contains the methods for the PrivateLinkServiceResourceOperationResults group.
// Don't use this type directly, use NewPrivateLinkServiceResourceOperationResultsClient() instead.
type PrivateLinkServiceResourceOperationResultsClient struct {
	host           string
	subscriptionID string
	operationID    string
	pl             runtime.Pipeline
}

// NewPrivateLinkServiceResourceOperationResultsClient creates a new instance of PrivateLinkServiceResourceOperationResultsClient with the specified values.
// subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
// operationID - The id of Azure async operation.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPrivateLinkServiceResourceOperationResultsClient(subscriptionID string, operationID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateLinkServiceResourceOperationResultsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &PrivateLinkServiceResourceOperationResultsClient{
		subscriptionID: subscriptionID,
		operationID:    operationID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginGet - Gets operation result of Private Link Service Resources for Power BI.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PrivateLinkServiceResourceOperationResultsClientBeginGetOptions contains the optional parameters for the PrivateLinkServiceResourceOperationResultsClient.BeginGet
// method.
func (client *PrivateLinkServiceResourceOperationResultsClient) BeginGet(ctx context.Context, options *PrivateLinkServiceResourceOperationResultsClientBeginGetOptions) (PrivateLinkServiceResourceOperationResultsClientGetPollerResponse, error) {
	resp, err := client.get(ctx, options)
	if err != nil {
		return PrivateLinkServiceResourceOperationResultsClientGetPollerResponse{}, err
	}
	result := PrivateLinkServiceResourceOperationResultsClientGetPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateLinkServiceResourceOperationResultsClient.Get", "azure-async-operation", resp, client.pl)
	if err != nil {
		return PrivateLinkServiceResourceOperationResultsClientGetPollerResponse{}, err
	}
	result.Poller = &PrivateLinkServiceResourceOperationResultsClientGetPoller{
		pt: pt,
	}
	return result, nil
}

// Get - Gets operation result of Private Link Service Resources for Power BI.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PrivateLinkServiceResourceOperationResultsClient) get(ctx context.Context, options *PrivateLinkServiceResourceOperationResultsClientBeginGetOptions) (*http.Response, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkServiceResourceOperationResultsClient) getCreateRequest(ctx context.Context, options *PrivateLinkServiceResourceOperationResultsClientBeginGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PowerBI/operationResults/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.operationID == "" {
		return nil, errors.New("parameter client.operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(client.operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
