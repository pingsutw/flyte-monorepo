// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	fosite "github.com/ory/fosite"

	http "net/http"

	interfaces "github.com/flyteorg/flyteadmin/auth/interfaces"

	jwk "github.com/lestrrat-go/jwx/jwk"

	mock "github.com/stretchr/testify/mock"

	oauth2 "github.com/ory/fosite/handler/oauth2"

	service "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/service"
)

// OAuth2Provider is an autogenerated mock type for the OAuth2Provider type
type OAuth2Provider struct {
	mock.Mock
}

type OAuth2Provider_IntrospectToken struct {
	*mock.Call
}

func (_m OAuth2Provider_IntrospectToken) Return(_a0 fosite.TokenType, _a1 fosite.AccessRequester, _a2 error) *OAuth2Provider_IntrospectToken {
	return &OAuth2Provider_IntrospectToken{Call: _m.Call.Return(_a0, _a1, _a2)}
}

func (_m *OAuth2Provider) OnIntrospectToken(ctx context.Context, token string, tokenUse fosite.TokenType, session fosite.Session, scope ...string) *OAuth2Provider_IntrospectToken {
	c_call := _m.On("IntrospectToken", ctx, token, tokenUse, session, scope)
	return &OAuth2Provider_IntrospectToken{Call: c_call}
}

func (_m *OAuth2Provider) OnIntrospectTokenMatch(matchers ...interface{}) *OAuth2Provider_IntrospectToken {
	c_call := _m.On("IntrospectToken", matchers...)
	return &OAuth2Provider_IntrospectToken{Call: c_call}
}

// IntrospectToken provides a mock function with given fields: ctx, token, tokenUse, session, scope
func (_m *OAuth2Provider) IntrospectToken(ctx context.Context, token string, tokenUse fosite.TokenType, session fosite.Session, scope ...string) (fosite.TokenType, fosite.AccessRequester, error) {
	_va := make([]interface{}, len(scope))
	for _i := range scope {
		_va[_i] = scope[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, token, tokenUse, session)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 fosite.TokenType
	if rf, ok := ret.Get(0).(func(context.Context, string, fosite.TokenType, fosite.Session, ...string) fosite.TokenType); ok {
		r0 = rf(ctx, token, tokenUse, session, scope...)
	} else {
		r0 = ret.Get(0).(fosite.TokenType)
	}

	var r1 fosite.AccessRequester
	if rf, ok := ret.Get(1).(func(context.Context, string, fosite.TokenType, fosite.Session, ...string) fosite.AccessRequester); ok {
		r1 = rf(ctx, token, tokenUse, session, scope...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(fosite.AccessRequester)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, fosite.TokenType, fosite.Session, ...string) error); ok {
		r2 = rf(ctx, token, tokenUse, session, scope...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type OAuth2Provider_KeySet struct {
	*mock.Call
}

func (_m OAuth2Provider_KeySet) Return(_a0 jwk.Set) *OAuth2Provider_KeySet {
	return &OAuth2Provider_KeySet{Call: _m.Call.Return(_a0)}
}

func (_m *OAuth2Provider) OnKeySet() *OAuth2Provider_KeySet {
	c_call := _m.On("KeySet")
	return &OAuth2Provider_KeySet{Call: c_call}
}

func (_m *OAuth2Provider) OnKeySetMatch(matchers ...interface{}) *OAuth2Provider_KeySet {
	c_call := _m.On("KeySet", matchers...)
	return &OAuth2Provider_KeySet{Call: c_call}
}

// KeySet provides a mock function with given fields:
func (_m *OAuth2Provider) KeySet() jwk.Set {
	ret := _m.Called()

	var r0 jwk.Set
	if rf, ok := ret.Get(0).(func() jwk.Set); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jwk.Set)
		}
	}

	return r0
}

type OAuth2Provider_NewAccessRequest struct {
	*mock.Call
}

func (_m OAuth2Provider_NewAccessRequest) Return(_a0 fosite.AccessRequester, _a1 error) *OAuth2Provider_NewAccessRequest {
	return &OAuth2Provider_NewAccessRequest{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OAuth2Provider) OnNewAccessRequest(ctx context.Context, req *http.Request, session fosite.Session) *OAuth2Provider_NewAccessRequest {
	c_call := _m.On("NewAccessRequest", ctx, req, session)
	return &OAuth2Provider_NewAccessRequest{Call: c_call}
}

func (_m *OAuth2Provider) OnNewAccessRequestMatch(matchers ...interface{}) *OAuth2Provider_NewAccessRequest {
	c_call := _m.On("NewAccessRequest", matchers...)
	return &OAuth2Provider_NewAccessRequest{Call: c_call}
}

// NewAccessRequest provides a mock function with given fields: ctx, req, session
func (_m *OAuth2Provider) NewAccessRequest(ctx context.Context, req *http.Request, session fosite.Session) (fosite.AccessRequester, error) {
	ret := _m.Called(ctx, req, session)

	var r0 fosite.AccessRequester
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request, fosite.Session) fosite.AccessRequester); ok {
		r0 = rf(ctx, req, session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fosite.AccessRequester)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request, fosite.Session) error); ok {
		r1 = rf(ctx, req, session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OAuth2Provider_NewAccessResponse struct {
	*mock.Call
}

func (_m OAuth2Provider_NewAccessResponse) Return(_a0 fosite.AccessResponder, _a1 error) *OAuth2Provider_NewAccessResponse {
	return &OAuth2Provider_NewAccessResponse{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OAuth2Provider) OnNewAccessResponse(ctx context.Context, requester fosite.AccessRequester) *OAuth2Provider_NewAccessResponse {
	c_call := _m.On("NewAccessResponse", ctx, requester)
	return &OAuth2Provider_NewAccessResponse{Call: c_call}
}

func (_m *OAuth2Provider) OnNewAccessResponseMatch(matchers ...interface{}) *OAuth2Provider_NewAccessResponse {
	c_call := _m.On("NewAccessResponse", matchers...)
	return &OAuth2Provider_NewAccessResponse{Call: c_call}
}

// NewAccessResponse provides a mock function with given fields: ctx, requester
func (_m *OAuth2Provider) NewAccessResponse(ctx context.Context, requester fosite.AccessRequester) (fosite.AccessResponder, error) {
	ret := _m.Called(ctx, requester)

	var r0 fosite.AccessResponder
	if rf, ok := ret.Get(0).(func(context.Context, fosite.AccessRequester) fosite.AccessResponder); ok {
		r0 = rf(ctx, requester)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fosite.AccessResponder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, fosite.AccessRequester) error); ok {
		r1 = rf(ctx, requester)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OAuth2Provider_NewAuthorizeRequest struct {
	*mock.Call
}

func (_m OAuth2Provider_NewAuthorizeRequest) Return(_a0 fosite.AuthorizeRequester, _a1 error) *OAuth2Provider_NewAuthorizeRequest {
	return &OAuth2Provider_NewAuthorizeRequest{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OAuth2Provider) OnNewAuthorizeRequest(ctx context.Context, req *http.Request) *OAuth2Provider_NewAuthorizeRequest {
	c_call := _m.On("NewAuthorizeRequest", ctx, req)
	return &OAuth2Provider_NewAuthorizeRequest{Call: c_call}
}

func (_m *OAuth2Provider) OnNewAuthorizeRequestMatch(matchers ...interface{}) *OAuth2Provider_NewAuthorizeRequest {
	c_call := _m.On("NewAuthorizeRequest", matchers...)
	return &OAuth2Provider_NewAuthorizeRequest{Call: c_call}
}

// NewAuthorizeRequest provides a mock function with given fields: ctx, req
func (_m *OAuth2Provider) NewAuthorizeRequest(ctx context.Context, req *http.Request) (fosite.AuthorizeRequester, error) {
	ret := _m.Called(ctx, req)

	var r0 fosite.AuthorizeRequester
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) fosite.AuthorizeRequester); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fosite.AuthorizeRequester)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OAuth2Provider_NewAuthorizeResponse struct {
	*mock.Call
}

func (_m OAuth2Provider_NewAuthorizeResponse) Return(_a0 fosite.AuthorizeResponder, _a1 error) *OAuth2Provider_NewAuthorizeResponse {
	return &OAuth2Provider_NewAuthorizeResponse{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OAuth2Provider) OnNewAuthorizeResponse(ctx context.Context, requester fosite.AuthorizeRequester, session fosite.Session) *OAuth2Provider_NewAuthorizeResponse {
	c_call := _m.On("NewAuthorizeResponse", ctx, requester, session)
	return &OAuth2Provider_NewAuthorizeResponse{Call: c_call}
}

func (_m *OAuth2Provider) OnNewAuthorizeResponseMatch(matchers ...interface{}) *OAuth2Provider_NewAuthorizeResponse {
	c_call := _m.On("NewAuthorizeResponse", matchers...)
	return &OAuth2Provider_NewAuthorizeResponse{Call: c_call}
}

// NewAuthorizeResponse provides a mock function with given fields: ctx, requester, session
func (_m *OAuth2Provider) NewAuthorizeResponse(ctx context.Context, requester fosite.AuthorizeRequester, session fosite.Session) (fosite.AuthorizeResponder, error) {
	ret := _m.Called(ctx, requester, session)

	var r0 fosite.AuthorizeResponder
	if rf, ok := ret.Get(0).(func(context.Context, fosite.AuthorizeRequester, fosite.Session) fosite.AuthorizeResponder); ok {
		r0 = rf(ctx, requester, session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fosite.AuthorizeResponder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, fosite.AuthorizeRequester, fosite.Session) error); ok {
		r1 = rf(ctx, requester, session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OAuth2Provider_NewIntrospectionRequest struct {
	*mock.Call
}

func (_m OAuth2Provider_NewIntrospectionRequest) Return(_a0 fosite.IntrospectionResponder, _a1 error) *OAuth2Provider_NewIntrospectionRequest {
	return &OAuth2Provider_NewIntrospectionRequest{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OAuth2Provider) OnNewIntrospectionRequest(ctx context.Context, r *http.Request, session fosite.Session) *OAuth2Provider_NewIntrospectionRequest {
	c_call := _m.On("NewIntrospectionRequest", ctx, r, session)
	return &OAuth2Provider_NewIntrospectionRequest{Call: c_call}
}

func (_m *OAuth2Provider) OnNewIntrospectionRequestMatch(matchers ...interface{}) *OAuth2Provider_NewIntrospectionRequest {
	c_call := _m.On("NewIntrospectionRequest", matchers...)
	return &OAuth2Provider_NewIntrospectionRequest{Call: c_call}
}

// NewIntrospectionRequest provides a mock function with given fields: ctx, r, session
func (_m *OAuth2Provider) NewIntrospectionRequest(ctx context.Context, r *http.Request, session fosite.Session) (fosite.IntrospectionResponder, error) {
	ret := _m.Called(ctx, r, session)

	var r0 fosite.IntrospectionResponder
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request, fosite.Session) fosite.IntrospectionResponder); ok {
		r0 = rf(ctx, r, session)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fosite.IntrospectionResponder)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *http.Request, fosite.Session) error); ok {
		r1 = rf(ctx, r, session)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type OAuth2Provider_NewJWTSessionToken struct {
	*mock.Call
}

func (_m OAuth2Provider_NewJWTSessionToken) Return(_a0 *oauth2.JWTSession) *OAuth2Provider_NewJWTSessionToken {
	return &OAuth2Provider_NewJWTSessionToken{Call: _m.Call.Return(_a0)}
}

func (_m *OAuth2Provider) OnNewJWTSessionToken(subject string, appID string, issuer string, audience string, userInfoClaims *service.UserInfoResponse) *OAuth2Provider_NewJWTSessionToken {
	c_call := _m.On("NewJWTSessionToken", subject, appID, issuer, audience, userInfoClaims)
	return &OAuth2Provider_NewJWTSessionToken{Call: c_call}
}

func (_m *OAuth2Provider) OnNewJWTSessionTokenMatch(matchers ...interface{}) *OAuth2Provider_NewJWTSessionToken {
	c_call := _m.On("NewJWTSessionToken", matchers...)
	return &OAuth2Provider_NewJWTSessionToken{Call: c_call}
}

// NewJWTSessionToken provides a mock function with given fields: subject, appID, issuer, audience, userInfoClaims
func (_m *OAuth2Provider) NewJWTSessionToken(subject string, appID string, issuer string, audience string, userInfoClaims *service.UserInfoResponse) *oauth2.JWTSession {
	ret := _m.Called(subject, appID, issuer, audience, userInfoClaims)

	var r0 *oauth2.JWTSession
	if rf, ok := ret.Get(0).(func(string, string, string, string, *service.UserInfoResponse) *oauth2.JWTSession); ok {
		r0 = rf(subject, appID, issuer, audience, userInfoClaims)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.JWTSession)
		}
	}

	return r0
}

type OAuth2Provider_NewRevocationRequest struct {
	*mock.Call
}

func (_m OAuth2Provider_NewRevocationRequest) Return(_a0 error) *OAuth2Provider_NewRevocationRequest {
	return &OAuth2Provider_NewRevocationRequest{Call: _m.Call.Return(_a0)}
}

func (_m *OAuth2Provider) OnNewRevocationRequest(ctx context.Context, r *http.Request) *OAuth2Provider_NewRevocationRequest {
	c_call := _m.On("NewRevocationRequest", ctx, r)
	return &OAuth2Provider_NewRevocationRequest{Call: c_call}
}

func (_m *OAuth2Provider) OnNewRevocationRequestMatch(matchers ...interface{}) *OAuth2Provider_NewRevocationRequest {
	c_call := _m.On("NewRevocationRequest", matchers...)
	return &OAuth2Provider_NewRevocationRequest{Call: c_call}
}

// NewRevocationRequest provides a mock function with given fields: ctx, r
func (_m *OAuth2Provider) NewRevocationRequest(ctx context.Context, r *http.Request) error {
	ret := _m.Called(ctx, r)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) error); ok {
		r0 = rf(ctx, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type OAuth2Provider_ValidateAccessToken struct {
	*mock.Call
}

func (_m OAuth2Provider_ValidateAccessToken) Return(_a0 interfaces.IdentityContext, _a1 error) *OAuth2Provider_ValidateAccessToken {
	return &OAuth2Provider_ValidateAccessToken{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *OAuth2Provider) OnValidateAccessToken(ctx context.Context, expectedAudience string, tokenStr string) *OAuth2Provider_ValidateAccessToken {
	c_call := _m.On("ValidateAccessToken", ctx, expectedAudience, tokenStr)
	return &OAuth2Provider_ValidateAccessToken{Call: c_call}
}

func (_m *OAuth2Provider) OnValidateAccessTokenMatch(matchers ...interface{}) *OAuth2Provider_ValidateAccessToken {
	c_call := _m.On("ValidateAccessToken", matchers...)
	return &OAuth2Provider_ValidateAccessToken{Call: c_call}
}

// ValidateAccessToken provides a mock function with given fields: ctx, expectedAudience, tokenStr
func (_m *OAuth2Provider) ValidateAccessToken(ctx context.Context, expectedAudience string, tokenStr string) (interfaces.IdentityContext, error) {
	ret := _m.Called(ctx, expectedAudience, tokenStr)

	var r0 interfaces.IdentityContext
	if rf, ok := ret.Get(0).(func(context.Context, string, string) interfaces.IdentityContext); ok {
		r0 = rf(ctx, expectedAudience, tokenStr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.IdentityContext)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, expectedAudience, tokenStr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteAccessError provides a mock function with given fields: rw, requester, err
func (_m *OAuth2Provider) WriteAccessError(rw http.ResponseWriter, requester fosite.AccessRequester, err error) {
	_m.Called(rw, requester, err)
}

// WriteAccessResponse provides a mock function with given fields: rw, requester, responder
func (_m *OAuth2Provider) WriteAccessResponse(rw http.ResponseWriter, requester fosite.AccessRequester, responder fosite.AccessResponder) {
	_m.Called(rw, requester, responder)
}

// WriteAuthorizeError provides a mock function with given fields: rw, requester, err
func (_m *OAuth2Provider) WriteAuthorizeError(rw http.ResponseWriter, requester fosite.AuthorizeRequester, err error) {
	_m.Called(rw, requester, err)
}

// WriteAuthorizeResponse provides a mock function with given fields: rw, requester, responder
func (_m *OAuth2Provider) WriteAuthorizeResponse(rw http.ResponseWriter, requester fosite.AuthorizeRequester, responder fosite.AuthorizeResponder) {
	_m.Called(rw, requester, responder)
}

// WriteIntrospectionError provides a mock function with given fields: rw, err
func (_m *OAuth2Provider) WriteIntrospectionError(rw http.ResponseWriter, err error) {
	_m.Called(rw, err)
}

// WriteIntrospectionResponse provides a mock function with given fields: rw, r
func (_m *OAuth2Provider) WriteIntrospectionResponse(rw http.ResponseWriter, r fosite.IntrospectionResponder) {
	_m.Called(rw, r)
}

// WriteRevocationResponse provides a mock function with given fields: rw, err
func (_m *OAuth2Provider) WriteRevocationResponse(rw http.ResponseWriter, err error) {
	_m.Called(rw, err)
}
