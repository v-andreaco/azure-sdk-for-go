// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// DiagnosticSettingsCategoryOperations contains the methods for the DiagnosticSettingsCategory group.
type DiagnosticSettingsCategoryOperations interface {
	// Get - Gets the diagnostic settings category for the specified resource.
	Get(ctx context.Context, resourceUri string, name string) (*DiagnosticSettingsCategoryResourceResponse, error)
	// List - Lists the diagnostic settings categories for the specified resource.
	List(ctx context.Context, resourceUri string) (*DiagnosticSettingsCategoryResourceCollectionResponse, error)
}

// diagnosticSettingsCategoryOperations implements the DiagnosticSettingsCategoryOperations interface.
type diagnosticSettingsCategoryOperations struct {
	*Client
}

// Get - Gets the diagnostic settings category for the specified resource.
func (client *diagnosticSettingsCategoryOperations) Get(ctx context.Context, resourceUri string, name string) (*DiagnosticSettingsCategoryResourceResponse, error) {
	req, err := client.getCreateRequest(resourceUri, name)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *diagnosticSettingsCategoryOperations) getCreateRequest(resourceUri string, name string) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/microsoft.insights/diagnosticSettingsCategories/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceUri)
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2017-05-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *diagnosticSettingsCategoryOperations) getHandleResponse(resp *azcore.Response) (*DiagnosticSettingsCategoryResourceResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := DiagnosticSettingsCategoryResourceResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DiagnosticSettingsCategoryResource)
}

// getHandleError handles the Get error response.
func (client *diagnosticSettingsCategoryOperations) getHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists the diagnostic settings categories for the specified resource.
func (client *diagnosticSettingsCategoryOperations) List(ctx context.Context, resourceUri string) (*DiagnosticSettingsCategoryResourceCollectionResponse, error) {
	req, err := client.listCreateRequest(resourceUri)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.listHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// listCreateRequest creates the List request.
func (client *diagnosticSettingsCategoryOperations) listCreateRequest(resourceUri string) (*azcore.Request, error) {
	urlPath := "/{resourceUri}/providers/microsoft.insights/diagnosticSettingsCategories"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceUri)
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2017-05-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *diagnosticSettingsCategoryOperations) listHandleResponse(resp *azcore.Response) (*DiagnosticSettingsCategoryResourceCollectionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := DiagnosticSettingsCategoryResourceCollectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DiagnosticSettingsCategoryResourceCollection)
}

// listHandleError handles the List error response.
func (client *diagnosticSettingsCategoryOperations) listHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}