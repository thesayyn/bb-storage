// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: pkg/proto/configuration/http/http.proto

package http

import (
	auth "github.com/buildbarn/bb-storage/pkg/proto/auth"
	jwt "github.com/buildbarn/bb-storage/pkg/proto/configuration/jwt"
	tls "github.com/buildbarn/bb-storage/pkg/proto/configuration/tls"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClientConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tls          *tls.ClientConfiguration            `protobuf:"bytes,1,opt,name=tls,proto3" json:"tls,omitempty"`
	ProxyUrl     string                              `protobuf:"bytes,2,opt,name=proxy_url,json=proxyUrl,proto3" json:"proxy_url,omitempty"`
	AddHeaders   []*ClientConfiguration_HeaderValues `protobuf:"bytes,5,rep,name=add_headers,json=addHeaders,proto3" json:"add_headers,omitempty"`
	DisableHttp2 bool                                `protobuf:"varint,6,opt,name=disable_http2,json=disableHttp2,proto3" json:"disable_http2,omitempty"`
}

func (x *ClientConfiguration) Reset() {
	*x = ClientConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfiguration) ProtoMessage() {}

func (x *ClientConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfiguration.ProtoReflect.Descriptor instead.
func (*ClientConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_http_http_proto_rawDescGZIP(), []int{0}
}

func (x *ClientConfiguration) GetTls() *tls.ClientConfiguration {
	if x != nil {
		return x.Tls
	}
	return nil
}

func (x *ClientConfiguration) GetProxyUrl() string {
	if x != nil {
		return x.ProxyUrl
	}
	return ""
}

func (x *ClientConfiguration) GetAddHeaders() []*ClientConfiguration_HeaderValues {
	if x != nil {
		return x.AddHeaders
	}
	return nil
}

func (x *ClientConfiguration) GetDisableHttp2() bool {
	if x != nil {
		return x.DisableHttp2
	}
	return false
}

type ServerConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenAddresses      []string              `protobuf:"bytes,1,rep,name=listen_addresses,json=listenAddresses,proto3" json:"listen_addresses,omitempty"`
	AuthenticationPolicy *AuthenticationPolicy `protobuf:"bytes,2,opt,name=authentication_policy,json=authenticationPolicy,proto3" json:"authentication_policy,omitempty"`
}

func (x *ServerConfiguration) Reset() {
	*x = ServerConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfiguration) ProtoMessage() {}

func (x *ServerConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfiguration.ProtoReflect.Descriptor instead.
func (*ServerConfiguration) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_http_http_proto_rawDescGZIP(), []int{1}
}

func (x *ServerConfiguration) GetListenAddresses() []string {
	if x != nil {
		return x.ListenAddresses
	}
	return nil
}

func (x *ServerConfiguration) GetAuthenticationPolicy() *AuthenticationPolicy {
	if x != nil {
		return x.AuthenticationPolicy
	}
	return nil
}

type AuthenticationPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Policy:
	//
	//	*AuthenticationPolicy_Allow
	//	*AuthenticationPolicy_Any
	//	*AuthenticationPolicy_Deny
	//	*AuthenticationPolicy_Jwt
	//	*AuthenticationPolicy_Oidc
	Policy isAuthenticationPolicy_Policy `protobuf_oneof:"policy"`
}

func (x *AuthenticationPolicy) Reset() {
	*x = AuthenticationPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationPolicy) ProtoMessage() {}

func (x *AuthenticationPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationPolicy.ProtoReflect.Descriptor instead.
func (*AuthenticationPolicy) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_http_http_proto_rawDescGZIP(), []int{2}
}

func (m *AuthenticationPolicy) GetPolicy() isAuthenticationPolicy_Policy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (x *AuthenticationPolicy) GetAllow() *auth.AuthenticationMetadata {
	if x, ok := x.GetPolicy().(*AuthenticationPolicy_Allow); ok {
		return x.Allow
	}
	return nil
}

func (x *AuthenticationPolicy) GetAny() *AnyAuthenticationPolicy {
	if x, ok := x.GetPolicy().(*AuthenticationPolicy_Any); ok {
		return x.Any
	}
	return nil
}

func (x *AuthenticationPolicy) GetDeny() string {
	if x, ok := x.GetPolicy().(*AuthenticationPolicy_Deny); ok {
		return x.Deny
	}
	return ""
}

func (x *AuthenticationPolicy) GetJwt() *jwt.AuthorizationHeaderParserConfiguration {
	if x, ok := x.GetPolicy().(*AuthenticationPolicy_Jwt); ok {
		return x.Jwt
	}
	return nil
}

func (x *AuthenticationPolicy) GetOidc() *OIDCAuthenticationPolicy {
	if x, ok := x.GetPolicy().(*AuthenticationPolicy_Oidc); ok {
		return x.Oidc
	}
	return nil
}

type isAuthenticationPolicy_Policy interface {
	isAuthenticationPolicy_Policy()
}

type AuthenticationPolicy_Allow struct {
	Allow *auth.AuthenticationMetadata `protobuf:"bytes,1,opt,name=allow,proto3,oneof"`
}

type AuthenticationPolicy_Any struct {
	Any *AnyAuthenticationPolicy `protobuf:"bytes,2,opt,name=any,proto3,oneof"`
}

type AuthenticationPolicy_Deny struct {
	Deny string `protobuf:"bytes,3,opt,name=deny,proto3,oneof"`
}

type AuthenticationPolicy_Jwt struct {
	Jwt *jwt.AuthorizationHeaderParserConfiguration `protobuf:"bytes,4,opt,name=jwt,proto3,oneof"`
}

type AuthenticationPolicy_Oidc struct {
	Oidc *OIDCAuthenticationPolicy `protobuf:"bytes,5,opt,name=oidc,proto3,oneof"`
}

func (*AuthenticationPolicy_Allow) isAuthenticationPolicy_Policy() {}

func (*AuthenticationPolicy_Any) isAuthenticationPolicy_Policy() {}

func (*AuthenticationPolicy_Deny) isAuthenticationPolicy_Policy() {}

func (*AuthenticationPolicy_Jwt) isAuthenticationPolicy_Policy() {}

func (*AuthenticationPolicy_Oidc) isAuthenticationPolicy_Policy() {}

type AnyAuthenticationPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Policies []*AuthenticationPolicy `protobuf:"bytes,1,rep,name=policies,proto3" json:"policies,omitempty"`
}

func (x *AnyAuthenticationPolicy) Reset() {
	*x = AnyAuthenticationPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnyAuthenticationPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnyAuthenticationPolicy) ProtoMessage() {}

func (x *AnyAuthenticationPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnyAuthenticationPolicy.ProtoReflect.Descriptor instead.
func (*AnyAuthenticationPolicy) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_http_http_proto_rawDescGZIP(), []int{3}
}

func (x *AnyAuthenticationPolicy) GetPolicies() []*AuthenticationPolicy {
	if x != nil {
		return x.Policies
	}
	return nil
}

type OIDCAuthenticationPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId                             string               `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret                         string               `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	AuthorizationEndpointUrl             string               `protobuf:"bytes,3,opt,name=authorization_endpoint_url,json=authorizationEndpointUrl,proto3" json:"authorization_endpoint_url,omitempty"`
	TokenEndpointUrl                     string               `protobuf:"bytes,4,opt,name=token_endpoint_url,json=tokenEndpointUrl,proto3" json:"token_endpoint_url,omitempty"`
	UserInfoEndpointUrl                  string               `protobuf:"bytes,5,opt,name=user_info_endpoint_url,json=userInfoEndpointUrl,proto3" json:"user_info_endpoint_url,omitempty"`
	MetadataExtractionJmespathExpression string               `protobuf:"bytes,6,opt,name=metadata_extraction_jmespath_expression,json=metadataExtractionJmespathExpression,proto3" json:"metadata_extraction_jmespath_expression,omitempty"`
	RedirectUrl                          string               `protobuf:"bytes,7,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	Scopes                               []string             `protobuf:"bytes,8,rep,name=scopes,proto3" json:"scopes,omitempty"`
	CookieSeed                           []byte               `protobuf:"bytes,9,opt,name=cookie_seed,json=cookieSeed,proto3" json:"cookie_seed,omitempty"`
	HttpClient                           *ClientConfiguration `protobuf:"bytes,10,opt,name=http_client,json=httpClient,proto3" json:"http_client,omitempty"`
}

func (x *OIDCAuthenticationPolicy) Reset() {
	*x = OIDCAuthenticationPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OIDCAuthenticationPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OIDCAuthenticationPolicy) ProtoMessage() {}

func (x *OIDCAuthenticationPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OIDCAuthenticationPolicy.ProtoReflect.Descriptor instead.
func (*OIDCAuthenticationPolicy) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_http_http_proto_rawDescGZIP(), []int{4}
}

func (x *OIDCAuthenticationPolicy) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *OIDCAuthenticationPolicy) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *OIDCAuthenticationPolicy) GetAuthorizationEndpointUrl() string {
	if x != nil {
		return x.AuthorizationEndpointUrl
	}
	return ""
}

func (x *OIDCAuthenticationPolicy) GetTokenEndpointUrl() string {
	if x != nil {
		return x.TokenEndpointUrl
	}
	return ""
}

func (x *OIDCAuthenticationPolicy) GetUserInfoEndpointUrl() string {
	if x != nil {
		return x.UserInfoEndpointUrl
	}
	return ""
}

func (x *OIDCAuthenticationPolicy) GetMetadataExtractionJmespathExpression() string {
	if x != nil {
		return x.MetadataExtractionJmespathExpression
	}
	return ""
}

func (x *OIDCAuthenticationPolicy) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *OIDCAuthenticationPolicy) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *OIDCAuthenticationPolicy) GetCookieSeed() []byte {
	if x != nil {
		return x.CookieSeed
	}
	return nil
}

func (x *OIDCAuthenticationPolicy) GetHttpClient() *ClientConfiguration {
	if x != nil {
		return x.HttpClient
	}
	return nil
}

type ClientConfiguration_HeaderValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header string   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ClientConfiguration_HeaderValues) Reset() {
	*x = ClientConfiguration_HeaderValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConfiguration_HeaderValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConfiguration_HeaderValues) ProtoMessage() {}

func (x *ClientConfiguration_HeaderValues) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_configuration_http_http_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConfiguration_HeaderValues.ProtoReflect.Descriptor instead.
func (*ClientConfiguration_HeaderValues) Descriptor() ([]byte, []int) {
	return file_pkg_proto_configuration_http_http_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ClientConfiguration_HeaderValues) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *ClientConfiguration_HeaderValues) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_pkg_proto_configuration_http_http_proto protoreflect.FileDescriptor

var file_pkg_proto_configuration_http_http_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x1a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x77, 0x74, 0x2f,
	0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xbc, 0x02, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x03, 0x74, 0x6c, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72,
	0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x74, 0x6c, 0x73, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x03, 0x74, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x5f, 0x0a, 0x0b, 0x61, 0x64, 0x64,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3e,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x0a,
	0x61, 0x64, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x74, 0x74, 0x70, 0x32, 0x1a,
	0x3e, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22,
	0xa9, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x12, 0x67, 0x0a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x14, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0xe8, 0x02, 0x0a, 0x14,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x12, 0x3e, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x05, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x49, 0x0a, 0x03, 0x61, 0x6e, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x35, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x41, 0x6e, 0x79, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x48, 0x00, 0x52, 0x03, 0x61, 0x6e, 0x79, 0x12,
	0x14, 0x0a, 0x04, 0x64, 0x65, 0x6e, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x04, 0x64, 0x65, 0x6e, 0x79, 0x12, 0x57, 0x0a, 0x03, 0x6a, 0x77, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x43, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6a, 0x77, 0x74,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x4c,
	0x0a, 0x04, 0x6f, 0x69, 0x64, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x4f, 0x49, 0x44, 0x43,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x48, 0x00, 0x52, 0x04, 0x6f, 0x69, 0x64, 0x63, 0x42, 0x08, 0x0a, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x69, 0x0a, 0x17, 0x41, 0x6e, 0x79, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x4e, 0x0a, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x08, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65,
	0x73, 0x22, 0x84, 0x04, 0x0a, 0x18, 0x4f, 0x49, 0x44, 0x43, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x12, 0x3c, 0x0a, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x18, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x2c,
	0x0a, 0x12, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x33, 0x0a, 0x16,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x55, 0x72,
	0x6c, 0x12, 0x55, 0x0a, 0x27, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x78,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6a, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x74,
	0x68, 0x5f, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x24, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4a, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x74, 0x68, 0x45, 0x78,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x5f, 0x73, 0x65,
	0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65,
	0x53, 0x65, 0x65, 0x64, 0x12, 0x52, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62, 0x75, 0x69, 0x6c,
	0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x68, 0x74,
	0x74, 0x70, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e,
	0x2f, 0x62, 0x62, 0x2d, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_configuration_http_http_proto_rawDescOnce sync.Once
	file_pkg_proto_configuration_http_http_proto_rawDescData = file_pkg_proto_configuration_http_http_proto_rawDesc
)

func file_pkg_proto_configuration_http_http_proto_rawDescGZIP() []byte {
	file_pkg_proto_configuration_http_http_proto_rawDescOnce.Do(func() {
		file_pkg_proto_configuration_http_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_configuration_http_http_proto_rawDescData)
	})
	return file_pkg_proto_configuration_http_http_proto_rawDescData
}

var file_pkg_proto_configuration_http_http_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_proto_configuration_http_http_proto_goTypes = []interface{}{
	(*ClientConfiguration)(nil),                        // 0: buildbarn.configuration.http.ClientConfiguration
	(*ServerConfiguration)(nil),                        // 1: buildbarn.configuration.http.ServerConfiguration
	(*AuthenticationPolicy)(nil),                       // 2: buildbarn.configuration.http.AuthenticationPolicy
	(*AnyAuthenticationPolicy)(nil),                    // 3: buildbarn.configuration.http.AnyAuthenticationPolicy
	(*OIDCAuthenticationPolicy)(nil),                   // 4: buildbarn.configuration.http.OIDCAuthenticationPolicy
	(*ClientConfiguration_HeaderValues)(nil),           // 5: buildbarn.configuration.http.ClientConfiguration.HeaderValues
	(*tls.ClientConfiguration)(nil),                    // 6: buildbarn.configuration.tls.ClientConfiguration
	(*auth.AuthenticationMetadata)(nil),                // 7: buildbarn.auth.AuthenticationMetadata
	(*jwt.AuthorizationHeaderParserConfiguration)(nil), // 8: buildbarn.configuration.jwt.AuthorizationHeaderParserConfiguration
}
var file_pkg_proto_configuration_http_http_proto_depIdxs = []int32{
	6, // 0: buildbarn.configuration.http.ClientConfiguration.tls:type_name -> buildbarn.configuration.tls.ClientConfiguration
	5, // 1: buildbarn.configuration.http.ClientConfiguration.add_headers:type_name -> buildbarn.configuration.http.ClientConfiguration.HeaderValues
	2, // 2: buildbarn.configuration.http.ServerConfiguration.authentication_policy:type_name -> buildbarn.configuration.http.AuthenticationPolicy
	7, // 3: buildbarn.configuration.http.AuthenticationPolicy.allow:type_name -> buildbarn.auth.AuthenticationMetadata
	3, // 4: buildbarn.configuration.http.AuthenticationPolicy.any:type_name -> buildbarn.configuration.http.AnyAuthenticationPolicy
	8, // 5: buildbarn.configuration.http.AuthenticationPolicy.jwt:type_name -> buildbarn.configuration.jwt.AuthorizationHeaderParserConfiguration
	4, // 6: buildbarn.configuration.http.AuthenticationPolicy.oidc:type_name -> buildbarn.configuration.http.OIDCAuthenticationPolicy
	2, // 7: buildbarn.configuration.http.AnyAuthenticationPolicy.policies:type_name -> buildbarn.configuration.http.AuthenticationPolicy
	0, // 8: buildbarn.configuration.http.OIDCAuthenticationPolicy.http_client:type_name -> buildbarn.configuration.http.ClientConfiguration
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_pkg_proto_configuration_http_http_proto_init() }
func file_pkg_proto_configuration_http_http_proto_init() {
	if File_pkg_proto_configuration_http_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_configuration_http_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_configuration_http_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_configuration_http_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_configuration_http_http_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnyAuthenticationPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_configuration_http_http_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OIDCAuthenticationPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_proto_configuration_http_http_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConfiguration_HeaderValues); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pkg_proto_configuration_http_http_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*AuthenticationPolicy_Allow)(nil),
		(*AuthenticationPolicy_Any)(nil),
		(*AuthenticationPolicy_Deny)(nil),
		(*AuthenticationPolicy_Jwt)(nil),
		(*AuthenticationPolicy_Oidc)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_configuration_http_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_proto_configuration_http_http_proto_goTypes,
		DependencyIndexes: file_pkg_proto_configuration_http_http_proto_depIdxs,
		MessageInfos:      file_pkg_proto_configuration_http_http_proto_msgTypes,
	}.Build()
	File_pkg_proto_configuration_http_http_proto = out.File
	file_pkg_proto_configuration_http_http_proto_rawDesc = nil
	file_pkg_proto_configuration_http_http_proto_goTypes = nil
	file_pkg_proto_configuration_http_http_proto_depIdxs = nil
}
