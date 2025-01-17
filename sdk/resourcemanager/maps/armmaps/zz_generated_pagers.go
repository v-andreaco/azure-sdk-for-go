//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaps

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AccountsClientListByResourceGroupPager provides operations for iterating over paged responses.
type AccountsClientListByResourceGroupPager struct {
	client    *AccountsClient
	current   AccountsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccountsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccountsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.Accounts.NextLink == nil || len(*p.current.Accounts.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccountsClientListByResourceGroupResponse page.
func (p *AccountsClientListByResourceGroupPager) PageResponse() AccountsClientListByResourceGroupResponse {
	return p.current
}

// AccountsClientListBySubscriptionPager provides operations for iterating over paged responses.
type AccountsClientListBySubscriptionPager struct {
	client    *AccountsClient
	current   AccountsClientListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccountsClientListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccountsClientListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccountsClientListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.Accounts.NextLink == nil || len(*p.current.Accounts.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccountsClientListBySubscriptionResponse page.
func (p *AccountsClientListBySubscriptionPager) PageResponse() AccountsClientListBySubscriptionResponse {
	return p.current
}

// ClientListOperationsPager provides operations for iterating over paged responses.
type ClientListOperationsPager struct {
	client    *Client
	current   ClientListOperationsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientListOperationsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientListOperationsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientListOperationsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.Operations.NextLink == nil || len(*p.current.Operations.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listOperationsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientListOperationsResponse page.
func (p *ClientListOperationsPager) PageResponse() ClientListOperationsResponse {
	return p.current
}

// ClientListSubscriptionOperationsPager provides operations for iterating over paged responses.
type ClientListSubscriptionOperationsPager struct {
	client    *Client
	current   ClientListSubscriptionOperationsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientListSubscriptionOperationsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ClientListSubscriptionOperationsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ClientListSubscriptionOperationsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.Operations.NextLink == nil || len(*p.current.Operations.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listSubscriptionOperationsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ClientListSubscriptionOperationsResponse page.
func (p *ClientListSubscriptionOperationsPager) PageResponse() ClientListSubscriptionOperationsResponse {
	return p.current
}

// CreatorsClientListByAccountPager provides operations for iterating over paged responses.
type CreatorsClientListByAccountPager struct {
	client    *CreatorsClient
	current   CreatorsClientListByAccountResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, CreatorsClientListByAccountResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *CreatorsClientListByAccountPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *CreatorsClientListByAccountPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CreatorList.NextLink == nil || len(*p.current.CreatorList.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByAccountHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current CreatorsClientListByAccountResponse page.
func (p *CreatorsClientListByAccountPager) PageResponse() CreatorsClientListByAccountResponse {
	return p.current
}
