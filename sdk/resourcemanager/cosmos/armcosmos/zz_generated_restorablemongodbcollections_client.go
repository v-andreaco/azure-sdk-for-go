//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// RestorableMongodbCollectionsClient contains the methods for the RestorableMongodbCollections group.
// Don't use this type directly, use NewRestorableMongodbCollectionsClient() instead.
type RestorableMongodbCollectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewRestorableMongodbCollectionsClient creates a new instance of RestorableMongodbCollectionsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewRestorableMongodbCollectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *RestorableMongodbCollectionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &RestorableMongodbCollectionsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Show the event feed of all mutations done on all the Azure Cosmos DB MongoDB collections under a specific database.
// This helps in scenario where container was accidentally deleted. This API requires
// 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission
// If the operation fails it returns an *azcore.ResponseError type.
// location - Cosmos DB region, with spaces between words and each word capitalized.
// instanceID - The instanceId GUID of a restorable database account.
// options - RestorableMongodbCollectionsClientListOptions contains the optional parameters for the RestorableMongodbCollectionsClient.List
// method.
func (client *RestorableMongodbCollectionsClient) List(ctx context.Context, location string, instanceID string, options *RestorableMongodbCollectionsClientListOptions) (RestorableMongodbCollectionsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, instanceID, options)
	if err != nil {
		return RestorableMongodbCollectionsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableMongodbCollectionsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableMongodbCollectionsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *RestorableMongodbCollectionsClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableMongodbCollectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbCollections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-15")
	if options != nil && options.RestorableMongodbDatabaseRid != nil {
		reqQP.Set("restorableMongodbDatabaseRid", *options.RestorableMongodbDatabaseRid)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableMongodbCollectionsClient) listHandleResponse(resp *http.Response) (RestorableMongodbCollectionsClientListResponse, error) {
	result := RestorableMongodbCollectionsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableMongodbCollectionsListResult); err != nil {
		return RestorableMongodbCollectionsClientListResponse{}, err
	}
	return result, nil
}
