//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azidentity

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

const credNameBrowser = "InteractiveBrowserCredential"

// InteractiveBrowserCredentialOptions contains optional parameters for InteractiveBrowserCredential.
type InteractiveBrowserCredentialOptions struct {
	azcore.ClientOptions

	// AdditionallyAllowedTenants specifies additional tenants for which the credential may acquire
	// tokens. Add the wildcard value "*" to allow the credential to acquire tokens for any tenant.
	AdditionallyAllowedTenants []string

	// AuthenticationRecord returned by a call to a credential's Authenticate method. Set this option
	// to enable the credential to use data from a previous authentication.
	AuthenticationRecord AuthenticationRecord

	// ClientID is the ID of the application users will authenticate to.
	// Defaults to the ID of an Azure development application.
	ClientID string

	// DisableAutomaticAuthentication prevents the credential from automatically prompting the user to authenticate.
	// When this option is true, [InteractiveBrowserCredential.GetToken] will return [ErrAuthenticationRequired] when
	// user interaction is necessary to acquire a token.
	DisableAutomaticAuthentication bool

	// DisableInstanceDiscovery should be set true only by applications authenticating in disconnected clouds, or
	// private clouds such as Azure Stack. It determines whether the credential requests Azure AD instance metadata
	// from https://login.microsoft.com before authenticating. Setting this to true will skip this request, making
	// the application responsible for ensuring the configured authority is valid and trustworthy.
	DisableInstanceDiscovery bool

	// LoginHint pre-populates the account prompt with a username. Users may choose to authenticate a different account.
	LoginHint string

	// RedirectURL is the URL Azure Active Directory will redirect to with the access token. This is required
	// only when setting ClientID, and must match a redirect URI in the application's registration.
	// Applications which have registered "http://localhost" as a redirect URI need not set this option.
	RedirectURL string

	// TenantID is the Azure Active Directory tenant the credential authenticates in. Defaults to the
	// "organizations" tenant, which can authenticate work and school accounts.
	TenantID string

	// TokenCachePersistenceOptions enables persistent token caching when not nil.
	TokenCachePersistenceOptions *TokenCachePersistenceOptions
}

func (o *InteractiveBrowserCredentialOptions) init() {
	if o.TenantID == "" {
		o.TenantID = organizationsTenantID
	}
	if o.ClientID == "" {
		o.ClientID = developerSignOnClientID
	}
}

// InteractiveBrowserCredential opens a browser to interactively authenticate a user.
type InteractiveBrowserCredential struct {
	client *publicClient
}

// NewInteractiveBrowserCredential constructs a new InteractiveBrowserCredential. Pass nil to accept default options.
func NewInteractiveBrowserCredential(options *InteractiveBrowserCredentialOptions) (*InteractiveBrowserCredential, error) {
	cp := InteractiveBrowserCredentialOptions{}
	if options != nil {
		cp = *options
	}
	cp.init()
	msalOpts := publicClientOptions{
		ClientOptions:                  cp.ClientOptions,
		DisableAutomaticAuthentication: cp.DisableAutomaticAuthentication,
		DisableInstanceDiscovery:       cp.DisableInstanceDiscovery,
		LoginHint:                      cp.LoginHint,
		Record:                         cp.AuthenticationRecord,
		RedirectURL:                    cp.RedirectURL,
		TokenCachePersistenceOptions:   cp.TokenCachePersistenceOptions,
	}
	c, err := newPublicClient(cp.TenantID, cp.ClientID, credNameBrowser, msalOpts)
	if err != nil {
		return nil, err
	}
	return &InteractiveBrowserCredential{client: c}, nil
}

// Authenticate a user via the default browser. Subsequent calls to GetToken will automatically use the returned AuthenticationRecord.
func (c *InteractiveBrowserCredential) Authenticate(ctx context.Context, opts *policy.TokenRequestOptions) (AuthenticationRecord, error) {
	return c.client.Authenticate(ctx, opts)
}

// GetToken requests an access token from Azure Active Directory. This method is called automatically by Azure SDK clients.
func (c *InteractiveBrowserCredential) GetToken(ctx context.Context, opts policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return c.client.GetToken(ctx, opts)
}

var _ azcore.TokenCredential = (*InteractiveBrowserCredential)(nil)
