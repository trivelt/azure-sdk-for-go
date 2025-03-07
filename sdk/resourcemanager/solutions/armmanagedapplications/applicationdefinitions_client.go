//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedapplications

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

// ApplicationDefinitionsClient contains the methods for the ApplicationDefinitions group.
// Don't use this type directly, use NewApplicationDefinitionsClient() instead.
type ApplicationDefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewApplicationDefinitionsClient creates a new instance of ApplicationDefinitionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName+".ApplicationDefinitionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ApplicationDefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - parameters - Parameters supplied to the create or update an managed application definition.
//   - options - ApplicationDefinitionsClientCreateOrUpdateOptions contains the optional parameters for the ApplicationDefinitionsClient.CreateOrUpdate
//     method.
func (client *ApplicationDefinitionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientCreateOrUpdateOptions) (ApplicationDefinitionsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
	if err != nil {
		return ApplicationDefinitionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationDefinitionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationDefinitionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ApplicationDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ApplicationDefinitionsClient) createOrUpdateHandleResponse(resp *http.Response) (ApplicationDefinitionsClientCreateOrUpdateResponse, error) {
	result := ApplicationDefinitionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateByID - Creates or updates a managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - parameters - Parameters supplied to the create or update a managed application definition.
//   - options - ApplicationDefinitionsClientCreateOrUpdateByIDOptions contains the optional parameters for the ApplicationDefinitionsClient.CreateOrUpdateByID
//     method.
func (client *ApplicationDefinitionsClient) CreateOrUpdateByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientCreateOrUpdateByIDOptions) (ApplicationDefinitionsClientCreateOrUpdateByIDResponse, error) {
	var err error
	req, err := client.createOrUpdateByIDCreateRequest(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
	if err != nil {
		return ApplicationDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	resp, err := client.createOrUpdateByIDHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateByIDCreateRequest creates the CreateOrUpdateByID request.
func (client *ApplicationDefinitionsClient) createOrUpdateByIDCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinition, options *ApplicationDefinitionsClientCreateOrUpdateByIDOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateByIDHandleResponse handles the CreateOrUpdateByID response.
func (client *ApplicationDefinitionsClient) createOrUpdateByIDHandleResponse(resp *http.Response) (ApplicationDefinitionsClientCreateOrUpdateByIDResponse, error) {
	result := ApplicationDefinitionsClientCreateOrUpdateByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsClientCreateOrUpdateByIDResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - options - ApplicationDefinitionsClientDeleteOptions contains the optional parameters for the ApplicationDefinitionsClient.Delete
//     method.
func (client *ApplicationDefinitionsClient) Delete(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientDeleteOptions) (ApplicationDefinitionsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return ApplicationDefinitionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationDefinitionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationDefinitionsClientDeleteResponse{}, err
	}
	return ApplicationDefinitionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ApplicationDefinitionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteByID - Deletes the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - options - ApplicationDefinitionsClientDeleteByIDOptions contains the optional parameters for the ApplicationDefinitionsClient.DeleteByID
//     method.
func (client *ApplicationDefinitionsClient) DeleteByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientDeleteByIDOptions) (ApplicationDefinitionsClientDeleteByIDResponse, error) {
	var err error
	req, err := client.deleteByIDCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return ApplicationDefinitionsClientDeleteByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationDefinitionsClientDeleteByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationDefinitionsClientDeleteByIDResponse{}, err
	}
	return ApplicationDefinitionsClientDeleteByIDResponse{}, nil
}

// deleteByIDCreateRequest creates the DeleteByID request.
func (client *ApplicationDefinitionsClient) deleteByIDCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientDeleteByIDOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - options - ApplicationDefinitionsClientGetOptions contains the optional parameters for the ApplicationDefinitionsClient.Get
//     method.
func (client *ApplicationDefinitionsClient) Get(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientGetOptions) (ApplicationDefinitionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return ApplicationDefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationDefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ApplicationDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ApplicationDefinitionsClient) getHandleResponse(resp *http.Response) (ApplicationDefinitionsClientGetResponse, error) {
	result := ApplicationDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// GetByID - Gets the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - options - ApplicationDefinitionsClientGetByIDOptions contains the optional parameters for the ApplicationDefinitionsClient.GetByID
//     method.
func (client *ApplicationDefinitionsClient) GetByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientGetByIDOptions) (ApplicationDefinitionsClientGetByIDResponse, error) {
	var err error
	req, err := client.getByIDCreateRequest(ctx, resourceGroupName, applicationDefinitionName, options)
	if err != nil {
		return ApplicationDefinitionsClientGetByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationDefinitionsClientGetByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNotFound) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationDefinitionsClientGetByIDResponse{}, err
	}
	resp, err := client.getByIDHandleResponse(httpResp)
	return resp, err
}

// getByIDCreateRequest creates the GetByID request.
func (client *ApplicationDefinitionsClient) getByIDCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, options *ApplicationDefinitionsClientGetByIDOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *ApplicationDefinitionsClient) getByIDHandleResponse(resp *http.Response) (ApplicationDefinitionsClientGetByIDResponse, error) {
	result := ApplicationDefinitionsClientGetByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsClientGetByIDResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists the managed application definitions in a resource group.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ApplicationDefinitionsClientListByResourceGroupOptions contains the optional parameters for the ApplicationDefinitionsClient.NewListByResourceGroupPager
//     method.
func (client *ApplicationDefinitionsClient) NewListByResourceGroupPager(resourceGroupName string, options *ApplicationDefinitionsClientListByResourceGroupOptions) *runtime.Pager[ApplicationDefinitionsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationDefinitionsClientListByResourceGroupResponse]{
		More: func(page ApplicationDefinitionsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationDefinitionsClientListByResourceGroupResponse) (ApplicationDefinitionsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationDefinitionsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ApplicationDefinitionsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationDefinitionsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ApplicationDefinitionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ApplicationDefinitionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ApplicationDefinitionsClient) listByResourceGroupHandleResponse(resp *http.Response) (ApplicationDefinitionsClientListByResourceGroupResponse, error) {
	result := ApplicationDefinitionsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinitionListResult); err != nil {
		return ApplicationDefinitionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all the application definitions within a subscription.
//
// Generated from API version 2021-07-01
//   - options - ApplicationDefinitionsClientListBySubscriptionOptions contains the optional parameters for the ApplicationDefinitionsClient.NewListBySubscriptionPager
//     method.
func (client *ApplicationDefinitionsClient) NewListBySubscriptionPager(options *ApplicationDefinitionsClientListBySubscriptionOptions) *runtime.Pager[ApplicationDefinitionsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationDefinitionsClientListBySubscriptionResponse]{
		More: func(page ApplicationDefinitionsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationDefinitionsClientListBySubscriptionResponse) (ApplicationDefinitionsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationDefinitionsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ApplicationDefinitionsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationDefinitionsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ApplicationDefinitionsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ApplicationDefinitionsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Solutions/applicationDefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ApplicationDefinitionsClient) listBySubscriptionHandleResponse(resp *http.Response) (ApplicationDefinitionsClientListBySubscriptionResponse, error) {
	result := ApplicationDefinitionsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinitionListResult); err != nil {
		return ApplicationDefinitionsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - parameters - Parameters supplied to the update a managed application definition.
//   - options - ApplicationDefinitionsClientUpdateOptions contains the optional parameters for the ApplicationDefinitionsClient.Update
//     method.
func (client *ApplicationDefinitionsClient) Update(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinitionPatchable, options *ApplicationDefinitionsClientUpdateOptions) (ApplicationDefinitionsClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
	if err != nil {
		return ApplicationDefinitionsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationDefinitionsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationDefinitionsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ApplicationDefinitionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinitionPatchable, options *ApplicationDefinitionsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ApplicationDefinitionsClient) updateHandleResponse(resp *http.Response) (ApplicationDefinitionsClientUpdateResponse, error) {
	result := ApplicationDefinitionsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsClientUpdateResponse{}, err
	}
	return result, nil
}

// UpdateByID - Updates the managed application definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationDefinitionName - The name of the managed application definition.
//   - parameters - Parameters supplied to the update a managed application definition.
//   - options - ApplicationDefinitionsClientUpdateByIDOptions contains the optional parameters for the ApplicationDefinitionsClient.UpdateByID
//     method.
func (client *ApplicationDefinitionsClient) UpdateByID(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinitionPatchable, options *ApplicationDefinitionsClientUpdateByIDOptions) (ApplicationDefinitionsClientUpdateByIDResponse, error) {
	var err error
	req, err := client.updateByIDCreateRequest(ctx, resourceGroupName, applicationDefinitionName, parameters, options)
	if err != nil {
		return ApplicationDefinitionsClientUpdateByIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ApplicationDefinitionsClientUpdateByIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ApplicationDefinitionsClientUpdateByIDResponse{}, err
	}
	resp, err := client.updateByIDHandleResponse(httpResp)
	return resp, err
}

// updateByIDCreateRequest creates the UpdateByID request.
func (client *ApplicationDefinitionsClient) updateByIDCreateRequest(ctx context.Context, resourceGroupName string, applicationDefinitionName string, parameters ApplicationDefinitionPatchable, options *ApplicationDefinitionsClientUpdateByIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationDefinitionName == "" {
		return nil, errors.New("parameter applicationDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationDefinitionName}", url.PathEscape(applicationDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateByIDHandleResponse handles the UpdateByID response.
func (client *ApplicationDefinitionsClient) updateByIDHandleResponse(resp *http.Response) (ApplicationDefinitionsClientUpdateByIDResponse, error) {
	result := ApplicationDefinitionsClientUpdateByIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationDefinition); err != nil {
		return ApplicationDefinitionsClientUpdateByIDResponse{}, err
	}
	return result, nil
}
