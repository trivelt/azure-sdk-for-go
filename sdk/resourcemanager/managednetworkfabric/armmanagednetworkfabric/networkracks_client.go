//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagednetworkfabric

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

// NetworkRacksClient contains the methods for the NetworkRacks group.
// Don't use this type directly, use NewNetworkRacksClient() instead.
type NetworkRacksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkRacksClient creates a new instance of NetworkRacksClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkRacksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkRacksClient, error) {
	cl, err := arm.NewClient(moduleName+".NetworkRacksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkRacksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create Network Rack resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkRackName - Name of the Network Rack.
//   - body - Request payload.
//   - options - NetworkRacksClientBeginCreateOptions contains the optional parameters for the NetworkRacksClient.BeginCreate
//     method.
func (client *NetworkRacksClient) BeginCreate(ctx context.Context, resourceGroupName string, networkRackName string, body NetworkRack, options *NetworkRacksClientBeginCreateOptions) (*runtime.Poller[NetworkRacksClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, networkRackName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkRacksClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkRacksClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Create Network Rack resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkRacksClient) create(ctx context.Context, resourceGroupName string, networkRackName string, body NetworkRack, options *NetworkRacksClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkRackName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *NetworkRacksClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkRackName string, body NetworkRack, options *NetworkRacksClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkRacks/{networkRackName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkRackName == "" {
		return nil, errors.New("parameter networkRackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkRackName}", url.PathEscape(networkRackName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginDelete - Delete Network Rack resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkRackName - Name of the Network Rack.
//   - options - NetworkRacksClientBeginDeleteOptions contains the optional parameters for the NetworkRacksClient.BeginDelete
//     method.
func (client *NetworkRacksClient) BeginDelete(ctx context.Context, resourceGroupName string, networkRackName string, options *NetworkRacksClientBeginDeleteOptions) (*runtime.Poller[NetworkRacksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkRackName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkRacksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkRacksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete Network Rack resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkRacksClient) deleteOperation(ctx context.Context, resourceGroupName string, networkRackName string, options *NetworkRacksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkRackName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NetworkRacksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkRackName string, options *NetworkRacksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkRacks/{networkRackName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkRackName == "" {
		return nil, errors.New("parameter networkRackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkRackName}", url.PathEscape(networkRackName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get Network Rack resource details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkRackName - Name of the Network Rack.
//   - options - NetworkRacksClientGetOptions contains the optional parameters for the NetworkRacksClient.Get method.
func (client *NetworkRacksClient) Get(ctx context.Context, resourceGroupName string, networkRackName string, options *NetworkRacksClientGetOptions) (NetworkRacksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkRackName, options)
	if err != nil {
		return NetworkRacksClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkRacksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkRacksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NetworkRacksClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkRackName string, options *NetworkRacksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkRacks/{networkRackName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkRackName == "" {
		return nil, errors.New("parameter networkRackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkRackName}", url.PathEscape(networkRackName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkRacksClient) getHandleResponse(resp *http.Response) (NetworkRacksClientGetResponse, error) {
	result := NetworkRacksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkRack); err != nil {
		return NetworkRacksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all Network Rack resources in the given resource group.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - NetworkRacksClientListByResourceGroupOptions contains the optional parameters for the NetworkRacksClient.NewListByResourceGroupPager
//     method.
func (client *NetworkRacksClient) NewListByResourceGroupPager(resourceGroupName string, options *NetworkRacksClientListByResourceGroupOptions) *runtime.Pager[NetworkRacksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkRacksClientListByResourceGroupResponse]{
		More: func(page NetworkRacksClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkRacksClientListByResourceGroupResponse) (NetworkRacksClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkRacksClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NetworkRacksClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkRacksClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NetworkRacksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkRacksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkRacks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *NetworkRacksClient) listByResourceGroupHandleResponse(resp *http.Response) (NetworkRacksClientListByResourceGroupResponse, error) {
	result := NetworkRacksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkRacksListResult); err != nil {
		return NetworkRacksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List all Network Rack resources in the given subscription
//
// Generated from API version 2023-06-15
//   - options - NetworkRacksClientListBySubscriptionOptions contains the optional parameters for the NetworkRacksClient.NewListBySubscriptionPager
//     method.
func (client *NetworkRacksClient) NewListBySubscriptionPager(options *NetworkRacksClientListBySubscriptionOptions) *runtime.Pager[NetworkRacksClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkRacksClientListBySubscriptionResponse]{
		More: func(page NetworkRacksClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkRacksClientListBySubscriptionResponse) (NetworkRacksClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkRacksClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NetworkRacksClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkRacksClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *NetworkRacksClient) listBySubscriptionCreateRequest(ctx context.Context, options *NetworkRacksClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetworkFabric/networkRacks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *NetworkRacksClient) listBySubscriptionHandleResponse(resp *http.Response) (NetworkRacksClientListBySubscriptionResponse, error) {
	result := NetworkRacksClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkRacksListResult); err != nil {
		return NetworkRacksClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update certain properties of the Network Rack resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkRackName - Name of the Network Rack.
//   - body - Network Rack properties to update.
//   - options - NetworkRacksClientBeginUpdateOptions contains the optional parameters for the NetworkRacksClient.BeginUpdate
//     method.
func (client *NetworkRacksClient) BeginUpdate(ctx context.Context, resourceGroupName string, networkRackName string, body TagsUpdate, options *NetworkRacksClientBeginUpdateOptions) (*runtime.Poller[NetworkRacksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, networkRackName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkRacksClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[NetworkRacksClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update certain properties of the Network Rack resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkRacksClient) update(ctx context.Context, resourceGroupName string, networkRackName string, body TagsUpdate, options *NetworkRacksClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, networkRackName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *NetworkRacksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, networkRackName string, body TagsUpdate, options *NetworkRacksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkRacks/{networkRackName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkRackName == "" {
		return nil, errors.New("parameter networkRackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkRackName}", url.PathEscape(networkRackName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}
