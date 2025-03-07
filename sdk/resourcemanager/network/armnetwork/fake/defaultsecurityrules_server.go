//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v4"
	"net/http"
	"net/url"
	"regexp"
)

// DefaultSecurityRulesServer is a fake server for instances of the armnetwork.DefaultSecurityRulesClient type.
type DefaultSecurityRulesServer struct {
	// Get is the fake for method DefaultSecurityRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkSecurityGroupName string, defaultSecurityRuleName string, options *armnetwork.DefaultSecurityRulesClientGetOptions) (resp azfake.Responder[armnetwork.DefaultSecurityRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DefaultSecurityRulesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkSecurityGroupName string, options *armnetwork.DefaultSecurityRulesClientListOptions) (resp azfake.PagerResponder[armnetwork.DefaultSecurityRulesClientListResponse])
}

// NewDefaultSecurityRulesServerTransport creates a new instance of DefaultSecurityRulesServerTransport with the provided implementation.
// The returned DefaultSecurityRulesServerTransport instance is connected to an instance of armnetwork.DefaultSecurityRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDefaultSecurityRulesServerTransport(srv *DefaultSecurityRulesServer) *DefaultSecurityRulesServerTransport {
	return &DefaultSecurityRulesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.DefaultSecurityRulesClientListResponse]](),
	}
}

// DefaultSecurityRulesServerTransport connects instances of armnetwork.DefaultSecurityRulesClient to instances of DefaultSecurityRulesServer.
// Don't use this type directly, use NewDefaultSecurityRulesServerTransport instead.
type DefaultSecurityRulesServerTransport struct {
	srv          *DefaultSecurityRulesServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.DefaultSecurityRulesClientListResponse]]
}

// Do implements the policy.Transporter interface for DefaultSecurityRulesServerTransport.
func (d *DefaultSecurityRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DefaultSecurityRulesClient.Get":
		resp, err = d.dispatchGet(req)
	case "DefaultSecurityRulesClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DefaultSecurityRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkSecurityGroups/(?P<networkSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/defaultSecurityRules/(?P<defaultSecurityRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkSecurityGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityGroupName")])
	if err != nil {
		return nil, err
	}
	defaultSecurityRuleNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("defaultSecurityRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameUnescaped, networkSecurityGroupNameUnescaped, defaultSecurityRuleNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SecurityRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DefaultSecurityRulesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkSecurityGroups/(?P<networkSecurityGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/defaultSecurityRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkSecurityGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityGroupName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListPager(resourceGroupNameUnescaped, networkSecurityGroupNameUnescaped, nil)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.DefaultSecurityRulesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}
