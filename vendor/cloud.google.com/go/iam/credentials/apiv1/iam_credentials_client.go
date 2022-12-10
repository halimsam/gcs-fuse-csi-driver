// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package credentials

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	credentialspb "cloud.google.com/go/iam/credentials/apiv1/credentialspb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newIamCredentialsClientHook clientHook

// IamCredentialsCallOptions contains the retry settings for each method of IamCredentialsClient.
type IamCredentialsCallOptions struct {
	GenerateAccessToken []gax.CallOption
	GenerateIdToken     []gax.CallOption
	SignBlob            []gax.CallOption
	SignJwt             []gax.CallOption
}

func defaultIamCredentialsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("iamcredentials.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("iamcredentials.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://iamcredentials.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultIamCredentialsCallOptions() *IamCredentialsCallOptions {
	return &IamCredentialsCallOptions{
		GenerateAccessToken: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GenerateIdToken: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SignBlob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SignJwt: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalIamCredentialsClient is an interface that defines the methods available from IAM Service Account Credentials API.
type internalIamCredentialsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GenerateAccessToken(context.Context, *credentialspb.GenerateAccessTokenRequest, ...gax.CallOption) (*credentialspb.GenerateAccessTokenResponse, error)
	GenerateIdToken(context.Context, *credentialspb.GenerateIdTokenRequest, ...gax.CallOption) (*credentialspb.GenerateIdTokenResponse, error)
	SignBlob(context.Context, *credentialspb.SignBlobRequest, ...gax.CallOption) (*credentialspb.SignBlobResponse, error)
	SignJwt(context.Context, *credentialspb.SignJwtRequest, ...gax.CallOption) (*credentialspb.SignJwtResponse, error)
}

// IamCredentialsClient is a client for interacting with IAM Service Account Credentials API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service account is a special type of Google account that belongs to your
// application or a virtual machine (VM), instead of to an individual end user.
// Your application assumes the identity of the service account to call Google
// APIs, so that the users aren’t directly involved.
//
// Service account credentials are used to temporarily assume the identity
// of the service account. Supported credential types include OAuth 2.0 access
// tokens, OpenID Connect ID tokens, self-signed JSON Web Tokens (JWTs), and
// more.
type IamCredentialsClient struct {
	// The internal transport-dependent client.
	internalClient internalIamCredentialsClient

	// The call options for this service.
	CallOptions *IamCredentialsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *IamCredentialsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *IamCredentialsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *IamCredentialsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GenerateAccessToken generates an OAuth 2.0 access token for a service account.
func (c *IamCredentialsClient) GenerateAccessToken(ctx context.Context, req *credentialspb.GenerateAccessTokenRequest, opts ...gax.CallOption) (*credentialspb.GenerateAccessTokenResponse, error) {
	return c.internalClient.GenerateAccessToken(ctx, req, opts...)
}

// GenerateIdToken generates an OpenID Connect ID token for a service account.
func (c *IamCredentialsClient) GenerateIdToken(ctx context.Context, req *credentialspb.GenerateIdTokenRequest, opts ...gax.CallOption) (*credentialspb.GenerateIdTokenResponse, error) {
	return c.internalClient.GenerateIdToken(ctx, req, opts...)
}

// SignBlob signs a blob using a service account’s system-managed private key.
func (c *IamCredentialsClient) SignBlob(ctx context.Context, req *credentialspb.SignBlobRequest, opts ...gax.CallOption) (*credentialspb.SignBlobResponse, error) {
	return c.internalClient.SignBlob(ctx, req, opts...)
}

// SignJwt signs a JWT using a service account’s system-managed private key.
func (c *IamCredentialsClient) SignJwt(ctx context.Context, req *credentialspb.SignJwtRequest, opts ...gax.CallOption) (*credentialspb.SignJwtResponse, error) {
	return c.internalClient.SignJwt(ctx, req, opts...)
}

// iamCredentialsGRPCClient is a client for interacting with IAM Service Account Credentials API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type iamCredentialsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing IamCredentialsClient
	CallOptions **IamCredentialsCallOptions

	// The gRPC API client.
	iamCredentialsClient credentialspb.IAMCredentialsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewIamCredentialsClient creates a new iam credentials client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service account is a special type of Google account that belongs to your
// application or a virtual machine (VM), instead of to an individual end user.
// Your application assumes the identity of the service account to call Google
// APIs, so that the users aren’t directly involved.
//
// Service account credentials are used to temporarily assume the identity
// of the service account. Supported credential types include OAuth 2.0 access
// tokens, OpenID Connect ID tokens, self-signed JSON Web Tokens (JWTs), and
// more.
func NewIamCredentialsClient(ctx context.Context, opts ...option.ClientOption) (*IamCredentialsClient, error) {
	clientOpts := defaultIamCredentialsGRPCClientOptions()
	if newIamCredentialsClientHook != nil {
		hookOpts, err := newIamCredentialsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := IamCredentialsClient{CallOptions: defaultIamCredentialsCallOptions()}

	c := &iamCredentialsGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		iamCredentialsClient: credentialspb.NewIAMCredentialsClient(connPool),
		CallOptions:          &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *iamCredentialsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *iamCredentialsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *iamCredentialsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *iamCredentialsGRPCClient) GenerateAccessToken(ctx context.Context, req *credentialspb.GenerateAccessTokenRequest, opts ...gax.CallOption) (*credentialspb.GenerateAccessTokenResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GenerateAccessToken[0:len((*c.CallOptions).GenerateAccessToken):len((*c.CallOptions).GenerateAccessToken)], opts...)
	var resp *credentialspb.GenerateAccessTokenResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamCredentialsClient.GenerateAccessToken(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *iamCredentialsGRPCClient) GenerateIdToken(ctx context.Context, req *credentialspb.GenerateIdTokenRequest, opts ...gax.CallOption) (*credentialspb.GenerateIdTokenResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GenerateIdToken[0:len((*c.CallOptions).GenerateIdToken):len((*c.CallOptions).GenerateIdToken)], opts...)
	var resp *credentialspb.GenerateIdTokenResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamCredentialsClient.GenerateIdToken(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *iamCredentialsGRPCClient) SignBlob(ctx context.Context, req *credentialspb.SignBlobRequest, opts ...gax.CallOption) (*credentialspb.SignBlobResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SignBlob[0:len((*c.CallOptions).SignBlob):len((*c.CallOptions).SignBlob)], opts...)
	var resp *credentialspb.SignBlobResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamCredentialsClient.SignBlob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *iamCredentialsGRPCClient) SignJwt(ctx context.Context, req *credentialspb.SignJwtRequest, opts ...gax.CallOption) (*credentialspb.SignJwtResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SignJwt[0:len((*c.CallOptions).SignJwt):len((*c.CallOptions).SignJwt)], opts...)
	var resp *credentialspb.SignJwtResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.iamCredentialsClient.SignJwt(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
