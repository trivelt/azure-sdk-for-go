//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

// GlobalReachConnectionsClient contains the methods for the GlobalReachConnections group.
// Don't use this type directly, use NewGlobalReachConnectionsClient() instead.
type GlobalReachConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGlobalReachConnectionsClient creates a new instance of GlobalReachConnectionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGlobalReachConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GlobalReachConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".GlobalReachConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GlobalReachConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a global reach connection in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - The name of the private cloud.
//   - globalReachConnectionName - Name of the global reach connection in the private cloud
//   - globalReachConnection - A global reach connection in the private cloud
//   - options - GlobalReachConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the GlobalReachConnectionsClient.BeginCreateOrUpdate
//     method.
func (client *GlobalReachConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, globalReachConnectionName string, globalReachConnection GlobalReachConnection, options *GlobalReachConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[GlobalReachConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, privateCloudName, globalReachConnectionName, globalReachConnection, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[GlobalReachConnectionsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[GlobalReachConnectionsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update a global reach connection in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *GlobalReachConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, privateCloudName string, globalReachConnectionName string, globalReachConnection GlobalReachConnection, options *GlobalReachConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, privateCloudName, globalReachConnectionName, globalReachConnection, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GlobalReachConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, globalReachConnectionName string, globalReachConnection GlobalReachConnection, options *GlobalReachConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/globalReachConnections/{globalReachConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if globalReachConnectionName == "" {
		return nil, errors.New("parameter globalReachConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalReachConnectionName}", url.PathEscape(globalReachConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, globalReachConnection); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a global reach connection in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - globalReachConnectionName - Name of the global reach connection in the private cloud
//   - options - GlobalReachConnectionsClientBeginDeleteOptions contains the optional parameters for the GlobalReachConnectionsClient.BeginDelete
//     method.
func (client *GlobalReachConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, privateCloudName string, globalReachConnectionName string, options *GlobalReachConnectionsClientBeginDeleteOptions) (*runtime.Poller[GlobalReachConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, privateCloudName, globalReachConnectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[GlobalReachConnectionsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[GlobalReachConnectionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a global reach connection in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *GlobalReachConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, privateCloudName string, globalReachConnectionName string, options *GlobalReachConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, privateCloudName, globalReachConnectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GlobalReachConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, globalReachConnectionName string, options *GlobalReachConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/globalReachConnections/{globalReachConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if globalReachConnectionName == "" {
		return nil, errors.New("parameter globalReachConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalReachConnectionName}", url.PathEscape(globalReachConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a global reach connection by name in a private cloud
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - globalReachConnectionName - Name of the global reach connection in the private cloud
//   - options - GlobalReachConnectionsClientGetOptions contains the optional parameters for the GlobalReachConnectionsClient.Get
//     method.
func (client *GlobalReachConnectionsClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, globalReachConnectionName string, options *GlobalReachConnectionsClientGetOptions) (GlobalReachConnectionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, globalReachConnectionName, options)
	if err != nil {
		return GlobalReachConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GlobalReachConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GlobalReachConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GlobalReachConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, globalReachConnectionName string, options *GlobalReachConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/globalReachConnections/{globalReachConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if globalReachConnectionName == "" {
		return nil, errors.New("parameter globalReachConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalReachConnectionName}", url.PathEscape(globalReachConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GlobalReachConnectionsClient) getHandleResponse(resp *http.Response) (GlobalReachConnectionsClientGetResponse, error) {
	result := GlobalReachConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GlobalReachConnection); err != nil {
		return GlobalReachConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List global reach connections in a private cloud
//
// Generated from API version 2023-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - options - GlobalReachConnectionsClientListOptions contains the optional parameters for the GlobalReachConnectionsClient.NewListPager
//     method.
func (client *GlobalReachConnectionsClient) NewListPager(resourceGroupName string, privateCloudName string, options *GlobalReachConnectionsClientListOptions) *runtime.Pager[GlobalReachConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GlobalReachConnectionsClientListResponse]{
		More: func(page GlobalReachConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GlobalReachConnectionsClientListResponse) (GlobalReachConnectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, privateCloudName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GlobalReachConnectionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return GlobalReachConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GlobalReachConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *GlobalReachConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *GlobalReachConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/globalReachConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GlobalReachConnectionsClient) listHandleResponse(resp *http.Response) (GlobalReachConnectionsClientListResponse, error) {
	result := GlobalReachConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GlobalReachConnectionList); err != nil {
		return GlobalReachConnectionsClientListResponse{}, err
	}
	return result, nil
}
