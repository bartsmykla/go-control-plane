// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.28.3
// source: envoy/service/auth/v3/external_auth.proto

package authv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The request attributes.
	Attributes *AttributeContext `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_auth_v3_external_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_auth_v3_external_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_envoy_service_auth_v3_external_auth_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetAttributes() *AttributeContext {
	if x != nil {
		return x.Attributes
	}
	return nil
}

// HTTP attributes for a denied response.
type DeniedHttpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field allows the authorization service to send an HTTP response status code to the
	// downstream client. If not set, Envoy sends “403 Forbidden“ HTTP status code by default.
	Status *v3.HttpStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// This field allows the authorization service to send HTTP response headers
	// to the downstream client. Note that the :ref:`append field in HeaderValueOption <envoy_v3_api_field_config.core.v3.HeaderValueOption.append>` defaults to
	// false when used in this message.
	Headers []*v31.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	// This field allows the authorization service to send a response body data
	// to the downstream client.
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *DeniedHttpResponse) Reset() {
	*x = DeniedHttpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_auth_v3_external_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeniedHttpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeniedHttpResponse) ProtoMessage() {}

func (x *DeniedHttpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_auth_v3_external_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeniedHttpResponse.ProtoReflect.Descriptor instead.
func (*DeniedHttpResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_auth_v3_external_auth_proto_rawDescGZIP(), []int{1}
}

func (x *DeniedHttpResponse) GetStatus() *v3.HttpStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *DeniedHttpResponse) GetHeaders() []*v31.HeaderValueOption {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *DeniedHttpResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

// HTTP attributes for an OK response.
// [#next-free-field: 9]
type OkHttpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HTTP entity headers in addition to the original request headers. This allows the authorization
	// service to append, to add or to override headers from the original request before
	// dispatching it to the upstream. Note that the :ref:`append field in HeaderValueOption <envoy_v3_api_field_config.core.v3.HeaderValueOption.append>` defaults to
	// false when used in this message. By setting the “append“ field to “true“,
	// the filter will append the correspondent header value to the matched request header.
	// By leaving “append“ as false, the filter will either add a new header, or override an existing
	// one if there is a match.
	Headers []*v31.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	// HTTP entity headers to remove from the original request before dispatching
	// it to the upstream. This allows the authorization service to act on auth
	// related headers (like “Authorization“), process them, and consume them.
	// Under this model, the upstream will either receive the request (if it's
	// authorized) or not receive it (if it's not), but will not see headers
	// containing authorization credentials.
	//
	// Pseudo headers (such as “:authority“, “:method“, “:path“ etc), as well as
	// the header “Host“, may not be removed as that would make the request
	// malformed. If mentioned in “headers_to_remove“ these special headers will
	// be ignored.
	//
	// When using the HTTP service this must instead be set by the HTTP
	// authorization service as a comma separated list like so:
	// “x-envoy-auth-headers-to-remove: one-auth-header, another-auth-header“.
	HeadersToRemove []string `protobuf:"bytes,5,rep,name=headers_to_remove,json=headersToRemove,proto3" json:"headers_to_remove,omitempty"`
	// This field has been deprecated in favor of :ref:`CheckResponse.dynamic_metadata
	// <envoy_v3_api_field_service.auth.v3.CheckResponse.dynamic_metadata>`. Until it is removed,
	// setting this field overrides :ref:`CheckResponse.dynamic_metadata
	// <envoy_v3_api_field_service.auth.v3.CheckResponse.dynamic_metadata>`.
	//
	// Deprecated: Marked as deprecated in envoy/service/auth/v3/external_auth.proto.
	DynamicMetadata *structpb.Struct `protobuf:"bytes,3,opt,name=dynamic_metadata,json=dynamicMetadata,proto3" json:"dynamic_metadata,omitempty"`
	// This field allows the authorization service to send HTTP response headers
	// to the downstream client on success. Note that the :ref:`append field in HeaderValueOption <envoy_v3_api_field_config.core.v3.HeaderValueOption.append>`
	// defaults to false when used in this message.
	ResponseHeadersToAdd []*v31.HeaderValueOption `protobuf:"bytes,6,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	// This field allows the authorization service to set (and overwrite) query
	// string parameters on the original request before it is sent upstream.
	QueryParametersToSet []*v31.QueryParameter `protobuf:"bytes,7,rep,name=query_parameters_to_set,json=queryParametersToSet,proto3" json:"query_parameters_to_set,omitempty"`
	// This field allows the authorization service to specify which query parameters
	// should be removed from the original request before it is sent upstream. Each
	// element in this list is a case-sensitive query parameter name to be removed.
	QueryParametersToRemove []string `protobuf:"bytes,8,rep,name=query_parameters_to_remove,json=queryParametersToRemove,proto3" json:"query_parameters_to_remove,omitempty"`
}

func (x *OkHttpResponse) Reset() {
	*x = OkHttpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_auth_v3_external_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OkHttpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OkHttpResponse) ProtoMessage() {}

func (x *OkHttpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_auth_v3_external_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OkHttpResponse.ProtoReflect.Descriptor instead.
func (*OkHttpResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_auth_v3_external_auth_proto_rawDescGZIP(), []int{2}
}

func (x *OkHttpResponse) GetHeaders() []*v31.HeaderValueOption {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *OkHttpResponse) GetHeadersToRemove() []string {
	if x != nil {
		return x.HeadersToRemove
	}
	return nil
}

// Deprecated: Marked as deprecated in envoy/service/auth/v3/external_auth.proto.
func (x *OkHttpResponse) GetDynamicMetadata() *structpb.Struct {
	if x != nil {
		return x.DynamicMetadata
	}
	return nil
}

func (x *OkHttpResponse) GetResponseHeadersToAdd() []*v31.HeaderValueOption {
	if x != nil {
		return x.ResponseHeadersToAdd
	}
	return nil
}

func (x *OkHttpResponse) GetQueryParametersToSet() []*v31.QueryParameter {
	if x != nil {
		return x.QueryParametersToSet
	}
	return nil
}

func (x *OkHttpResponse) GetQueryParametersToRemove() []string {
	if x != nil {
		return x.QueryParametersToRemove
	}
	return nil
}

// Intended for gRPC and Network Authorization servers “only“.
type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status “OK“ allows the request. Any other status indicates the request should be denied, and
	// for HTTP filter, if not overridden by :ref:`denied HTTP response status <envoy_v3_api_field_service.auth.v3.DeniedHttpResponse.status>`
	// Envoy sends “403 Forbidden“ HTTP status code by default.
	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// An message that contains HTTP response attributes. This message is
	// used when the authorization service needs to send custom responses to the
	// downstream client or, to modify/add request headers being dispatched to the upstream.
	//
	// Types that are assignable to HttpResponse:
	//
	//	*CheckResponse_DeniedResponse
	//	*CheckResponse_OkResponse
	HttpResponse isCheckResponse_HttpResponse `protobuf_oneof:"http_response"`
	// Optional response metadata that will be emitted as dynamic metadata to be consumed by the next
	// filter. This metadata lives in a namespace specified by the canonical name of extension filter
	// that requires it:
	//
	// - :ref:`envoy.filters.http.ext_authz <config_http_filters_ext_authz_dynamic_metadata>` for HTTP filter.
	// - :ref:`envoy.filters.network.ext_authz <config_network_filters_ext_authz_dynamic_metadata>` for network filter.
	DynamicMetadata *structpb.Struct `protobuf:"bytes,4,opt,name=dynamic_metadata,json=dynamicMetadata,proto3" json:"dynamic_metadata,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_auth_v3_external_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_auth_v3_external_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_auth_v3_external_auth_proto_rawDescGZIP(), []int{3}
}

func (x *CheckResponse) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (m *CheckResponse) GetHttpResponse() isCheckResponse_HttpResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (x *CheckResponse) GetDeniedResponse() *DeniedHttpResponse {
	if x, ok := x.GetHttpResponse().(*CheckResponse_DeniedResponse); ok {
		return x.DeniedResponse
	}
	return nil
}

func (x *CheckResponse) GetOkResponse() *OkHttpResponse {
	if x, ok := x.GetHttpResponse().(*CheckResponse_OkResponse); ok {
		return x.OkResponse
	}
	return nil
}

func (x *CheckResponse) GetDynamicMetadata() *structpb.Struct {
	if x != nil {
		return x.DynamicMetadata
	}
	return nil
}

type isCheckResponse_HttpResponse interface {
	isCheckResponse_HttpResponse()
}

type CheckResponse_DeniedResponse struct {
	// Supplies http attributes for a denied response.
	DeniedResponse *DeniedHttpResponse `protobuf:"bytes,2,opt,name=denied_response,json=deniedResponse,proto3,oneof"`
}

type CheckResponse_OkResponse struct {
	// Supplies http attributes for an ok response.
	OkResponse *OkHttpResponse `protobuf:"bytes,3,opt,name=ok_response,json=okResponse,proto3,oneof"`
}

func (*CheckResponse_DeniedResponse) isCheckResponse_HttpResponse() {}

func (*CheckResponse_OkResponse) isCheckResponse_HttpResponse() {}

var File_envoy_service_auth_v3_external_auth_proto protoreflect.FileDescriptor

var file_envoy_service_auth_v3_external_auth_proto_rawDesc = []byte{
	0x0a, 0x29, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65,
	0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x47, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x3a, 0x29, 0x9a, 0xc5, 0x88,
	0x1e, 0x24, 0x0a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xcf, 0x01, 0x0a, 0x12, 0x44, 0x65, 0x6e, 0x69, 0x65,
	0x64, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x74,
	0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x41, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x2f, 0x9a, 0xc5, 0x88, 0x1e, 0x2a, 0x0a, 0x28,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x48, 0x74, 0x74, 0x70,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xf7, 0x03, 0x0a, 0x0e, 0x4f, 0x6b, 0x48,
	0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x2a,
	0x0a, 0x11, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x4f, 0x0a, 0x10, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x0b, 0x92,
	0xc7, 0x86, 0xd8, 0x04, 0x03, 0x33, 0x2e, 0x30, 0x18, 0x01, 0x52, 0x0f, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5e, 0x0a, 0x17, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f,
	0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x5b, 0x0a, 0x17, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f,
	0x74, 0x6f, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x52, 0x14, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x73, 0x54, 0x6f, 0x53, 0x65, 0x74, 0x12, 0x3b, 0x0a, 0x1a, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x17, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x3a, 0x2b, 0x9a, 0xc5, 0x88, 0x1e, 0x26, 0x0a, 0x24, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x6b, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0xdc, 0x02, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x54, 0x0a, 0x0f, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76,
	0x33, 0x2e, 0x44, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0e, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x6f, 0x6b, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x33, 0x2e, 0x4f, 0x6b, 0x48, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x42, 0x0a, 0x10, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x3a, 0x2a, 0x9a, 0xc5, 0x88, 0x1e, 0x25, 0x0a, 0x23, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x32, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0f, 0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x65, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x54, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x23, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x33, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x87, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x02, 0x0a, 0x23, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x33, 0x42, 0x11, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x33, 0x3b, 0x61, 0x75, 0x74, 0x68,
	0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_auth_v3_external_auth_proto_rawDescOnce sync.Once
	file_envoy_service_auth_v3_external_auth_proto_rawDescData = file_envoy_service_auth_v3_external_auth_proto_rawDesc
)

func file_envoy_service_auth_v3_external_auth_proto_rawDescGZIP() []byte {
	file_envoy_service_auth_v3_external_auth_proto_rawDescOnce.Do(func() {
		file_envoy_service_auth_v3_external_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_auth_v3_external_auth_proto_rawDescData)
	})
	return file_envoy_service_auth_v3_external_auth_proto_rawDescData
}

var file_envoy_service_auth_v3_external_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_service_auth_v3_external_auth_proto_goTypes = []interface{}{
	(*CheckRequest)(nil),          // 0: envoy.service.auth.v3.CheckRequest
	(*DeniedHttpResponse)(nil),    // 1: envoy.service.auth.v3.DeniedHttpResponse
	(*OkHttpResponse)(nil),        // 2: envoy.service.auth.v3.OkHttpResponse
	(*CheckResponse)(nil),         // 3: envoy.service.auth.v3.CheckResponse
	(*AttributeContext)(nil),      // 4: envoy.service.auth.v3.AttributeContext
	(*v3.HttpStatus)(nil),         // 5: envoy.type.v3.HttpStatus
	(*v31.HeaderValueOption)(nil), // 6: envoy.config.core.v3.HeaderValueOption
	(*structpb.Struct)(nil),       // 7: google.protobuf.Struct
	(*v31.QueryParameter)(nil),    // 8: envoy.config.core.v3.QueryParameter
	(*status.Status)(nil),         // 9: google.rpc.Status
}
var file_envoy_service_auth_v3_external_auth_proto_depIdxs = []int32{
	4,  // 0: envoy.service.auth.v3.CheckRequest.attributes:type_name -> envoy.service.auth.v3.AttributeContext
	5,  // 1: envoy.service.auth.v3.DeniedHttpResponse.status:type_name -> envoy.type.v3.HttpStatus
	6,  // 2: envoy.service.auth.v3.DeniedHttpResponse.headers:type_name -> envoy.config.core.v3.HeaderValueOption
	6,  // 3: envoy.service.auth.v3.OkHttpResponse.headers:type_name -> envoy.config.core.v3.HeaderValueOption
	7,  // 4: envoy.service.auth.v3.OkHttpResponse.dynamic_metadata:type_name -> google.protobuf.Struct
	6,  // 5: envoy.service.auth.v3.OkHttpResponse.response_headers_to_add:type_name -> envoy.config.core.v3.HeaderValueOption
	8,  // 6: envoy.service.auth.v3.OkHttpResponse.query_parameters_to_set:type_name -> envoy.config.core.v3.QueryParameter
	9,  // 7: envoy.service.auth.v3.CheckResponse.status:type_name -> google.rpc.Status
	1,  // 8: envoy.service.auth.v3.CheckResponse.denied_response:type_name -> envoy.service.auth.v3.DeniedHttpResponse
	2,  // 9: envoy.service.auth.v3.CheckResponse.ok_response:type_name -> envoy.service.auth.v3.OkHttpResponse
	7,  // 10: envoy.service.auth.v3.CheckResponse.dynamic_metadata:type_name -> google.protobuf.Struct
	0,  // 11: envoy.service.auth.v3.Authorization.Check:input_type -> envoy.service.auth.v3.CheckRequest
	3,  // 12: envoy.service.auth.v3.Authorization.Check:output_type -> envoy.service.auth.v3.CheckResponse
	12, // [12:13] is the sub-list for method output_type
	11, // [11:12] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_envoy_service_auth_v3_external_auth_proto_init() }
func file_envoy_service_auth_v3_external_auth_proto_init() {
	if File_envoy_service_auth_v3_external_auth_proto != nil {
		return
	}
	file_envoy_service_auth_v3_attribute_context_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_auth_v3_external_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_envoy_service_auth_v3_external_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeniedHttpResponse); i {
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
		file_envoy_service_auth_v3_external_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OkHttpResponse); i {
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
		file_envoy_service_auth_v3_external_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
	file_envoy_service_auth_v3_external_auth_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*CheckResponse_DeniedResponse)(nil),
		(*CheckResponse_OkResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_service_auth_v3_external_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_auth_v3_external_auth_proto_goTypes,
		DependencyIndexes: file_envoy_service_auth_v3_external_auth_proto_depIdxs,
		MessageInfos:      file_envoy_service_auth_v3_external_auth_proto_msgTypes,
	}.Build()
	File_envoy_service_auth_v3_external_auth_proto = out.File
	file_envoy_service_auth_v3_external_auth_proto_rawDesc = nil
	file_envoy_service_auth_v3_external_auth_proto_goTypes = nil
	file_envoy_service_auth_v3_external_auth_proto_depIdxs = nil
}
