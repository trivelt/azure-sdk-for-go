//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataprotection

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DppResourceGuardProxyClient contains the methods for the DppResourceGuardProxy group.
// Don't use this type directly, use NewDppResourceGuardProxyClient() instead.
type DppResourceGuardProxyClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDppResourceGuardProxyClient creates a new instance of DppResourceGuardProxyClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDppResourceGuardProxyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DppResourceGuardProxyClient, error) {
	cl, err := arm.NewClient(moduleName+".DppResourceGuardProxyClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DppResourceGuardProxyClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or Updates a ResourceGuardProxy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - resourceGuardProxyName - name of the resource guard proxy
//   - parameters - Request body for operation
//   - options - DppResourceGuardProxyClientCreateOrUpdateOptions contains the optional parameters for the DppResourceGuardProxyClient.CreateOrUpdate
//     method.
func (client *DppResourceGuardProxyClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, parameters ResourceGuardProxyBaseResource, options *DppResourceGuardProxyClientCreateOrUpdateOptions) (DppResourceGuardProxyClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vaultName, resourceGuardProxyName, parameters, options)
	if err != nil {
		return DppResourceGuardProxyClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DppResourceGuardProxyClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DppResourceGuardProxyClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DppResourceGuardProxyClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, parameters ResourceGuardProxyBaseResource, options *DppResourceGuardProxyClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DppResourceGuardProxyClient) createOrUpdateHandleResponse(resp *http.Response) (DppResourceGuardProxyClientCreateOrUpdateResponse, error) {
	result := DppResourceGuardProxyClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResource); err != nil {
		return DppResourceGuardProxyClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the ResourceGuardProxy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - resourceGuardProxyName - name of the resource guard proxy
//   - options - DppResourceGuardProxyClientDeleteOptions contains the optional parameters for the DppResourceGuardProxyClient.Delete
//     method.
func (client *DppResourceGuardProxyClient) Delete(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, options *DppResourceGuardProxyClientDeleteOptions) (DppResourceGuardProxyClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vaultName, resourceGuardProxyName, options)
	if err != nil {
		return DppResourceGuardProxyClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DppResourceGuardProxyClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DppResourceGuardProxyClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DppResourceGuardProxyClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DppResourceGuardProxyClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, options *DppResourceGuardProxyClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the ResourceGuardProxy object associated with the vault, and that matches the name in the request
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - resourceGuardProxyName - name of the resource guard proxy
//   - options - DppResourceGuardProxyClientGetOptions contains the optional parameters for the DppResourceGuardProxyClient.Get
//     method.
func (client *DppResourceGuardProxyClient) Get(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, options *DppResourceGuardProxyClientGetOptions) (DppResourceGuardProxyClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, resourceGuardProxyName, options)
	if err != nil {
		return DppResourceGuardProxyClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DppResourceGuardProxyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DppResourceGuardProxyClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DppResourceGuardProxyClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, options *DppResourceGuardProxyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DppResourceGuardProxyClient) getHandleResponse(resp *http.Response) (DppResourceGuardProxyClientGetResponse, error) {
	result := DppResourceGuardProxyClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResource); err != nil {
		return DppResourceGuardProxyClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns the list of ResourceGuardProxies associated with the vault
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - options - DppResourceGuardProxyClientListOptions contains the optional parameters for the DppResourceGuardProxyClient.NewListPager
//     method.
func (client *DppResourceGuardProxyClient) NewListPager(resourceGroupName string, vaultName string, options *DppResourceGuardProxyClientListOptions) *runtime.Pager[DppResourceGuardProxyClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DppResourceGuardProxyClientListResponse]{
		More: func(page DppResourceGuardProxyClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DppResourceGuardProxyClientListResponse) (DppResourceGuardProxyClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, vaultName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DppResourceGuardProxyClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DppResourceGuardProxyClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DppResourceGuardProxyClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DppResourceGuardProxyClient) listCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *DppResourceGuardProxyClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupResourceGuardProxies"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DppResourceGuardProxyClient) listHandleResponse(resp *http.Response) (DppResourceGuardProxyClientListResponse, error) {
	result := DppResourceGuardProxyClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResourceList); err != nil {
		return DppResourceGuardProxyClientListResponse{}, err
	}
	return result, nil
}

// UnlockDelete - UnlockDelete call for ResourceGuardProxy, executed before one can delete it
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - resourceGuardProxyName - name of the resource guard proxy
//   - parameters - Request body for operation
//   - options - DppResourceGuardProxyClientUnlockDeleteOptions contains the optional parameters for the DppResourceGuardProxyClient.UnlockDelete
//     method.
func (client *DppResourceGuardProxyClient) UnlockDelete(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, parameters UnlockDeleteRequest, options *DppResourceGuardProxyClientUnlockDeleteOptions) (DppResourceGuardProxyClientUnlockDeleteResponse, error) {
	req, err := client.unlockDeleteCreateRequest(ctx, resourceGroupName, vaultName, resourceGuardProxyName, parameters, options)
	if err != nil {
		return DppResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DppResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DppResourceGuardProxyClientUnlockDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.unlockDeleteHandleResponse(resp)
}

// unlockDeleteCreateRequest creates the UnlockDelete request.
func (client *DppResourceGuardProxyClient) unlockDeleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, parameters UnlockDeleteRequest, options *DppResourceGuardProxyClientUnlockDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}/unlockDelete"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// unlockDeleteHandleResponse handles the UnlockDelete response.
func (client *DppResourceGuardProxyClient) unlockDeleteHandleResponse(resp *http.Response) (DppResourceGuardProxyClientUnlockDeleteResponse, error) {
	result := DppResourceGuardProxyClientUnlockDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UnlockDeleteResponse); err != nil {
		return DppResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	return result, nil
}
