//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

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
	"strings"
)

// ApplicationDefinitionsClient contains the methods for the ApplicationDefinitions group.
// Don't use this type directly, use NewApplicationDefinitionsClient() instead.
type ApplicationDefinitionsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewApplicationDefinitionsClient creates a new instance of ApplicationDefinitionsClient with the specified values.
func NewApplicationDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ApplicationDefinitionsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ApplicationDefinitionsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Creates a new managed application definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationDefinitionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsBeginCreateOrUpdateOptions) (ApplicationDefinitionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
	if err != nil {
		return ApplicationDefinitionsCreateOrUpdatePollerResponse{}, err
	}
	result := ApplicationDefinitionsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ApplicationDefinitionsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ApplicationDefinitionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ApplicationDefinitionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a new managed application definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationDefinitionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ApplicationDefinitionsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreateOrUpdateByID - Creates a new managed application definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationDefinitionsClient) BeginCreateOrUpdateByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsBeginCreateOrUpdateByIDOptions) (ApplicationDefinitionsCreateOrUpdateByIDPollerResponse, error) {
	resp, err := client.createOrUpdateByID(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
	if err != nil {
		return ApplicationDefinitionsCreateOrUpdateByIDPollerResponse{}, err
	}
	result := ApplicationDefinitionsCreateOrUpdateByIDPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ApplicationDefinitionsClient.CreateOrUpdateByID", "", resp, client.pl, client.createOrUpdateByIDHandleError)
	if err != nil {
		return ApplicationDefinitionsCreateOrUpdateByIDPollerResponse{}, err
	}
	result.Poller = &ApplicationDefinitionsCreateOrUpdateByIDPoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdateByID - Creates a new managed application definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationDefinitionsClient) createOrUpdateByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsBeginCreateOrUpdateByIDOptions) (*http.Response, error) {
	req, err := client.createOrUpdateByIDCreateRequest(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateByIDHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateByIDCreateRequest creates the CreateOrUpdateByID request.
func (client *ApplicationDefinitionsClient) createOrUpdateByIDCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsBeginCreateOrUpdateByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateByIDHandleError handles the CreateOrUpdateByID error response.
func (client *ApplicationDefinitionsClient) createOrUpdateByIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the managed application definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationDefinitionsClient) BeginDelete(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsBeginDeleteOptions) (ApplicationDefinitionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return ApplicationDefinitionsDeletePollerResponse{}, err
	}
	result := ApplicationDefinitionsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ApplicationDefinitionsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ApplicationDefinitionsDeletePollerResponse{}, err
	}
	result.Poller = &ApplicationDefinitionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the managed application definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationDefinitionsClient) deleteOperation(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationDefinitionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ApplicationDefinitionsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDeleteByID - Deletes the managed application definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationDefinitionsClient) BeginDeleteByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsBeginDeleteByIDOptions) (ApplicationDefinitionsDeleteByIDPollerResponse, error) {
	resp, err := client.deleteByID(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return ApplicationDefinitionsDeleteByIDPollerResponse{}, err
	}
	result := ApplicationDefinitionsDeleteByIDPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ApplicationDefinitionsClient.DeleteByID", "", resp, client.pl, client.deleteByIDHandleError)
	if err != nil {
		return ApplicationDefinitionsDeleteByIDPollerResponse{}, err
	}
	result.Poller = &ApplicationDefinitionsDeleteByIDPoller{
		pt: pt,
	}
	return result, nil
}

// DeleteByID - Deletes the managed application definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationDefinitionsClient) deleteByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsBeginDeleteByIDOptions) (*http.Response, error) {
	req, err := client.deleteByIDCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteByIDHandleError(resp)
	}
	return resp, nil
}

// deleteByIDCreateRequest creates the DeleteByID request.
func (client *ApplicationDefinitionsClient) deleteByIDCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsBeginDeleteByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteByIDHandleError handles the DeleteByID error response.
func (client *ApplicationDefinitionsClient) deleteByIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the managed application definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationDefinitionsClient) Get(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsGetOptions) (ApplicationDefinitionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return ApplicationDefinitionsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationDefinitionsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return ApplicationDefinitionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ApplicationDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationDefinitionsClient) getHandleResponse(resp *http.Response) (ApplicationDefinitionsGetResponse, error) {
	result := ApplicationDefinitionsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ApplicationDefinitionsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetByID - Gets the managed application definition.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationDefinitionsClient) GetByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsGetByIDOptions) (ApplicationDefinitionsGetByIDResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return ApplicationDefinitionsGetByIDResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ApplicationDefinitionsGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotFound) {
		return ApplicationDefinitionsGetByIDResponse{}, client.getByIDHandleError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *ApplicationDefinitionsClient) getByIDCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsGetByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *ApplicationDefinitionsClient) getByIDHandleResponse(resp *http.Response) (ApplicationDefinitionsGetByIDResponse, error) {
	result := ApplicationDefinitionsGetByIDResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsGetByIDResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getByIDHandleError handles the GetByID error response.
func (client *ApplicationDefinitionsClient) getByIDHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Lists the managed application definitions in a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ApplicationDefinitionsClient) ListByResourceGroup(resourceGroupName string, options *ApplicationDefinitionsListByResourceGroupOptions) *ApplicationDefinitionsListByResourceGroupPager {
	return &ApplicationDefinitionsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ApplicationDefinitionsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ApplicationDefinitionListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ApplicationDefinitionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationDefinitionsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ApplicationDefinitionsClient) listByResourceGroupHandleResponse(resp *http.Response) (ApplicationDefinitionsListByResourceGroupResponse, error) {
	result := ApplicationDefinitionsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinitionListResult); err != nil {
		return ApplicationDefinitionsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ApplicationDefinitionsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}