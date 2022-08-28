// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// https://developers.google.com/protocol-buffers/
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conformance/conformance.proto

package conformance

import (
	protoreflect "github.com/seeflood/protobuf-go/reflect/protoreflect"
	protoimpl "github.com/seeflood/protobuf-go/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type WireFormat int32

const (
	WireFormat_UNSPECIFIED WireFormat = 0
	WireFormat_PROTOBUF    WireFormat = 1
	WireFormat_JSON        WireFormat = 2
	WireFormat_JSPB        WireFormat = 3 // Google internal only. Opensource testees just skip it.
	WireFormat_TEXT_FORMAT WireFormat = 4
)

// Enum value maps for WireFormat.
var (
	WireFormat_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "PROTOBUF",
		2: "JSON",
		3: "JSPB",
		4: "TEXT_FORMAT",
	}
	WireFormat_value = map[string]int32{
		"UNSPECIFIED": 0,
		"PROTOBUF":    1,
		"JSON":        2,
		"JSPB":        3,
		"TEXT_FORMAT": 4,
	}
)

func (x WireFormat) Enum() *WireFormat {
	p := new(WireFormat)
	*p = x
	return p
}

func (x WireFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WireFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_conformance_conformance_proto_enumTypes[0].Descriptor()
}

func (WireFormat) Type() protoreflect.EnumType {
	return &file_conformance_conformance_proto_enumTypes[0]
}

func (x WireFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WireFormat.Descriptor instead.
func (WireFormat) EnumDescriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{0}
}

type TestCategory int32

const (
	TestCategory_UNSPECIFIED_TEST TestCategory = 0
	TestCategory_BINARY_TEST      TestCategory = 1 // Test binary wire format.
	TestCategory_JSON_TEST        TestCategory = 2 // Test json wire format.
	// Similar to JSON_TEST. However, during parsing json, testee should ignore
	// unknown fields. This feature is optional. Each implementation can decide
	// whether to support it.  See
	// https://developers.google.com/protocol-buffers/docs/proto3#json_options
	// for more detail.
	TestCategory_JSON_IGNORE_UNKNOWN_PARSING_TEST TestCategory = 3
	// Test jspb wire format. Google internal only. Opensource testees just skip it.
	TestCategory_JSPB_TEST TestCategory = 4
	// Test text format. For cpp, java and python, testees can already deal with
	// this type. Testees of other languages can simply skip it.
	TestCategory_TEXT_FORMAT_TEST TestCategory = 5
)

// Enum value maps for TestCategory.
var (
	TestCategory_name = map[int32]string{
		0: "UNSPECIFIED_TEST",
		1: "BINARY_TEST",
		2: "JSON_TEST",
		3: "JSON_IGNORE_UNKNOWN_PARSING_TEST",
		4: "JSPB_TEST",
		5: "TEXT_FORMAT_TEST",
	}
	TestCategory_value = map[string]int32{
		"UNSPECIFIED_TEST":                 0,
		"BINARY_TEST":                      1,
		"JSON_TEST":                        2,
		"JSON_IGNORE_UNKNOWN_PARSING_TEST": 3,
		"JSPB_TEST":                        4,
		"TEXT_FORMAT_TEST":                 5,
	}
)

func (x TestCategory) Enum() *TestCategory {
	p := new(TestCategory)
	*p = x
	return p
}

func (x TestCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TestCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_conformance_conformance_proto_enumTypes[1].Descriptor()
}

func (TestCategory) Type() protoreflect.EnumType {
	return &file_conformance_conformance_proto_enumTypes[1]
}

func (x TestCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TestCategory.Descriptor instead.
func (TestCategory) EnumDescriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{1}
}

// The conformance runner will request a list of failures as the first request.
// This will be known by message_type == "conformance.FailureSet", a conformance
// test should return a serialized FailureSet in protobuf_payload.
type FailureSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Failure []string `protobuf:"bytes,1,rep,name=failure,proto3" json:"failure,omitempty"`
}

func (x *FailureSet) Reset() {
	*x = FailureSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conformance_conformance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailureSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailureSet) ProtoMessage() {}

func (x *FailureSet) ProtoReflect() protoreflect.Message {
	mi := &file_conformance_conformance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailureSet.ProtoReflect.Descriptor instead.
func (*FailureSet) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{0}
}

func (x *FailureSet) GetFailure() []string {
	if x != nil {
		return x.Failure
	}
	return nil
}

// Represents a single test case's input.  The testee should:
//
//   1. parse this proto (which should always succeed)
//   2. parse the protobuf or JSON payload in "payload" (which may fail)
//   3. if the parse succeeded, serialize the message in the requested format.
type ConformanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The payload (whether protobuf of JSON) is always for a
	// protobuf_test_messages.proto3.TestAllTypes proto (as defined in
	// src/google/protobuf/proto3_test_messages.proto).
	//
	// TODO(haberman): if/when we expand the conformance tests to support proto2,
	// we will want to include a field that lets the payload/response be a
	// protobuf_test_messages.google.protobuf.TestAllTypes message instead.
	//
	// Types that are assignable to Payload:
	//	*ConformanceRequest_ProtobufPayload
	//	*ConformanceRequest_JsonPayload
	//	*ConformanceRequest_JspbPayload
	//	*ConformanceRequest_TextPayload
	Payload isConformanceRequest_Payload `protobuf_oneof:"payload"`
	// Which format should the testee serialize its message to?
	RequestedOutputFormat WireFormat `protobuf:"varint,3,opt,name=requested_output_format,json=requestedOutputFormat,proto3,enum=conformance.WireFormat" json:"requested_output_format,omitempty"`
	// The full name for the test message to use; for the moment, either:
	// protobuf_test_messages.proto3.TestAllTypesProto3 or
	// protobuf_test_messages.google.protobuf.TestAllTypesProto2.
	MessageType string `protobuf:"bytes,4,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	// Each test is given a specific test category. Some category may need
	// specific support in testee programs. Refer to the definition of TestCategory
	// for more information.
	TestCategory TestCategory `protobuf:"varint,5,opt,name=test_category,json=testCategory,proto3,enum=conformance.TestCategory" json:"test_category,omitempty"`
	// Specify details for how to encode jspb.
	JspbEncodingOptions *JspbEncodingConfig `protobuf:"bytes,6,opt,name=jspb_encoding_options,json=jspbEncodingOptions,proto3" json:"jspb_encoding_options,omitempty"`
	// This can be used in json and text format. If true, testee should print
	// unknown fields instead of ignore. This feature is optional.
	PrintUnknownFields bool `protobuf:"varint,9,opt,name=print_unknown_fields,json=printUnknownFields,proto3" json:"print_unknown_fields,omitempty"`
}

func (x *ConformanceRequest) Reset() {
	*x = ConformanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conformance_conformance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConformanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConformanceRequest) ProtoMessage() {}

func (x *ConformanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_conformance_conformance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConformanceRequest.ProtoReflect.Descriptor instead.
func (*ConformanceRequest) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{1}
}

func (m *ConformanceRequest) GetPayload() isConformanceRequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *ConformanceRequest) GetProtobufPayload() []byte {
	if x, ok := x.GetPayload().(*ConformanceRequest_ProtobufPayload); ok {
		return x.ProtobufPayload
	}
	return nil
}

func (x *ConformanceRequest) GetJsonPayload() string {
	if x, ok := x.GetPayload().(*ConformanceRequest_JsonPayload); ok {
		return x.JsonPayload
	}
	return ""
}

func (x *ConformanceRequest) GetJspbPayload() string {
	if x, ok := x.GetPayload().(*ConformanceRequest_JspbPayload); ok {
		return x.JspbPayload
	}
	return ""
}

func (x *ConformanceRequest) GetTextPayload() string {
	if x, ok := x.GetPayload().(*ConformanceRequest_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

func (x *ConformanceRequest) GetRequestedOutputFormat() WireFormat {
	if x != nil {
		return x.RequestedOutputFormat
	}
	return WireFormat_UNSPECIFIED
}

func (x *ConformanceRequest) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *ConformanceRequest) GetTestCategory() TestCategory {
	if x != nil {
		return x.TestCategory
	}
	return TestCategory_UNSPECIFIED_TEST
}

func (x *ConformanceRequest) GetJspbEncodingOptions() *JspbEncodingConfig {
	if x != nil {
		return x.JspbEncodingOptions
	}
	return nil
}

func (x *ConformanceRequest) GetPrintUnknownFields() bool {
	if x != nil {
		return x.PrintUnknownFields
	}
	return false
}

type isConformanceRequest_Payload interface {
	isConformanceRequest_Payload()
}

type ConformanceRequest_ProtobufPayload struct {
	ProtobufPayload []byte `protobuf:"bytes,1,opt,name=protobuf_payload,json=protobufPayload,proto3,oneof"`
}

type ConformanceRequest_JsonPayload struct {
	JsonPayload string `protobuf:"bytes,2,opt,name=json_payload,json=jsonPayload,proto3,oneof"`
}

type ConformanceRequest_JspbPayload struct {
	// Google internal only.  Opensource testees just skip it.
	JspbPayload string `protobuf:"bytes,7,opt,name=jspb_payload,json=jspbPayload,proto3,oneof"`
}

type ConformanceRequest_TextPayload struct {
	TextPayload string `protobuf:"bytes,8,opt,name=text_payload,json=textPayload,proto3,oneof"`
}

func (*ConformanceRequest_ProtobufPayload) isConformanceRequest_Payload() {}

func (*ConformanceRequest_JsonPayload) isConformanceRequest_Payload() {}

func (*ConformanceRequest_JspbPayload) isConformanceRequest_Payload() {}

func (*ConformanceRequest_TextPayload) isConformanceRequest_Payload() {}

// Represents a single test case's output.
type ConformanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//	*ConformanceResponse_ParseError
	//	*ConformanceResponse_SerializeError
	//	*ConformanceResponse_RuntimeError
	//	*ConformanceResponse_ProtobufPayload
	//	*ConformanceResponse_JsonPayload
	//	*ConformanceResponse_Skipped
	//	*ConformanceResponse_JspbPayload
	//	*ConformanceResponse_TextPayload
	Result isConformanceResponse_Result `protobuf_oneof:"result"`
}

func (x *ConformanceResponse) Reset() {
	*x = ConformanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conformance_conformance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConformanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConformanceResponse) ProtoMessage() {}

func (x *ConformanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_conformance_conformance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConformanceResponse.ProtoReflect.Descriptor instead.
func (*ConformanceResponse) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{2}
}

func (m *ConformanceResponse) GetResult() isConformanceResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ConformanceResponse) GetParseError() string {
	if x, ok := x.GetResult().(*ConformanceResponse_ParseError); ok {
		return x.ParseError
	}
	return ""
}

func (x *ConformanceResponse) GetSerializeError() string {
	if x, ok := x.GetResult().(*ConformanceResponse_SerializeError); ok {
		return x.SerializeError
	}
	return ""
}

func (x *ConformanceResponse) GetRuntimeError() string {
	if x, ok := x.GetResult().(*ConformanceResponse_RuntimeError); ok {
		return x.RuntimeError
	}
	return ""
}

func (x *ConformanceResponse) GetProtobufPayload() []byte {
	if x, ok := x.GetResult().(*ConformanceResponse_ProtobufPayload); ok {
		return x.ProtobufPayload
	}
	return nil
}

func (x *ConformanceResponse) GetJsonPayload() string {
	if x, ok := x.GetResult().(*ConformanceResponse_JsonPayload); ok {
		return x.JsonPayload
	}
	return ""
}

func (x *ConformanceResponse) GetSkipped() string {
	if x, ok := x.GetResult().(*ConformanceResponse_Skipped); ok {
		return x.Skipped
	}
	return ""
}

func (x *ConformanceResponse) GetJspbPayload() string {
	if x, ok := x.GetResult().(*ConformanceResponse_JspbPayload); ok {
		return x.JspbPayload
	}
	return ""
}

func (x *ConformanceResponse) GetTextPayload() string {
	if x, ok := x.GetResult().(*ConformanceResponse_TextPayload); ok {
		return x.TextPayload
	}
	return ""
}

type isConformanceResponse_Result interface {
	isConformanceResponse_Result()
}

type ConformanceResponse_ParseError struct {
	// This string should be set to indicate parsing failed.  The string can
	// provide more information about the parse error if it is available.
	//
	// Setting this string does not necessarily mean the testee failed the
	// test.  Some of the test cases are intentionally invalid input.
	ParseError string `protobuf:"bytes,1,opt,name=parse_error,json=parseError,proto3,oneof"`
}

type ConformanceResponse_SerializeError struct {
	// If the input was successfully parsed but errors occurred when
	// serializing it to the requested output format, set the error message in
	// this field.
	SerializeError string `protobuf:"bytes,6,opt,name=serialize_error,json=serializeError,proto3,oneof"`
}

type ConformanceResponse_RuntimeError struct {
	// This should be set if some other error occurred.  This will always
	// indicate that the test failed.  The string can provide more information
	// about the failure.
	RuntimeError string `protobuf:"bytes,2,opt,name=runtime_error,json=runtimeError,proto3,oneof"`
}

type ConformanceResponse_ProtobufPayload struct {
	// If the input was successfully parsed and the requested output was
	// protobuf, serialize it to protobuf and set it in this field.
	ProtobufPayload []byte `protobuf:"bytes,3,opt,name=protobuf_payload,json=protobufPayload,proto3,oneof"`
}

type ConformanceResponse_JsonPayload struct {
	// If the input was successfully parsed and the requested output was JSON,
	// serialize to JSON and set it in this field.
	JsonPayload string `protobuf:"bytes,4,opt,name=json_payload,json=jsonPayload,proto3,oneof"`
}

type ConformanceResponse_Skipped struct {
	// For when the testee skipped the test, likely because a certain feature
	// wasn't supported, like JSON input/output.
	Skipped string `protobuf:"bytes,5,opt,name=skipped,proto3,oneof"`
}

type ConformanceResponse_JspbPayload struct {
	// If the input was successfully parsed and the requested output was JSPB,
	// serialize to JSPB and set it in this field. JSPB is google internal only
	// format. Opensource testees can just skip it.
	JspbPayload string `protobuf:"bytes,7,opt,name=jspb_payload,json=jspbPayload,proto3,oneof"`
}

type ConformanceResponse_TextPayload struct {
	// If the input was successfully parsed and the requested output was
	// TEXT_FORMAT, serialize to TEXT_FORMAT and set it in this field.
	TextPayload string `protobuf:"bytes,8,opt,name=text_payload,json=textPayload,proto3,oneof"`
}

func (*ConformanceResponse_ParseError) isConformanceResponse_Result() {}

func (*ConformanceResponse_SerializeError) isConformanceResponse_Result() {}

func (*ConformanceResponse_RuntimeError) isConformanceResponse_Result() {}

func (*ConformanceResponse_ProtobufPayload) isConformanceResponse_Result() {}

func (*ConformanceResponse_JsonPayload) isConformanceResponse_Result() {}

func (*ConformanceResponse_Skipped) isConformanceResponse_Result() {}

func (*ConformanceResponse_JspbPayload) isConformanceResponse_Result() {}

func (*ConformanceResponse_TextPayload) isConformanceResponse_Result() {}

// Encoding options for jspb format.
type JspbEncodingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Encode the value field of Any as jspb array if true, otherwise binary.
	UseJspbArrayAnyFormat bool `protobuf:"varint,1,opt,name=use_jspb_array_any_format,json=useJspbArrayAnyFormat,proto3" json:"use_jspb_array_any_format,omitempty"`
}

func (x *JspbEncodingConfig) Reset() {
	*x = JspbEncodingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conformance_conformance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JspbEncodingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JspbEncodingConfig) ProtoMessage() {}

func (x *JspbEncodingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_conformance_conformance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JspbEncodingConfig.ProtoReflect.Descriptor instead.
func (*JspbEncodingConfig) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{3}
}

func (x *JspbEncodingConfig) GetUseJspbArrayAnyFormat() bool {
	if x != nil {
		return x.UseJspbArrayAnyFormat
	}
	return false
}

var File_conformance_conformance_proto protoreflect.FileDescriptor

var file_conformance_conformance_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x0a,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x53, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x22, 0xf6, 0x03, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x10, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x6a, 0x73, 0x6f, 0x6e,
	0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a,
	0x0c, 0x6a, 0x73, 0x70, 0x62, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6a, 0x73, 0x70, 0x62, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x78, 0x74,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x4f, 0x0a, 0x17, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x57, 0x69, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3e, 0x0a, 0x0d, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0c, 0x74,
	0x65, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x53, 0x0a, 0x15, 0x6a,
	0x73, 0x70, 0x62, 0x5f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4a, 0x73, 0x70, 0x62, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x13, 0x6a, 0x73, 0x70,
	0x62, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12,
	0x70, 0x72, 0x69, 0x6e, 0x74, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xcc, 0x02,
	0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x73, 0x65, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0d, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x10, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x0a, 0x07,
	0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x07, 0x73, 0x6b, 0x69, 0x70, 0x70, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x6a, 0x73, 0x70, 0x62,
	0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x0b, 0x6a, 0x73, 0x70, 0x62, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a,
	0x0c, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x78, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x4e, 0x0a, 0x12,
	0x4a, 0x73, 0x70, 0x62, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x38, 0x0a, 0x19, 0x75, 0x73, 0x65, 0x5f, 0x6a, 0x73, 0x70, 0x62, 0x5f, 0x61,
	0x72, 0x72, 0x61, 0x79, 0x5f, 0x61, 0x6e, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x75, 0x73, 0x65, 0x4a, 0x73, 0x70, 0x62, 0x41, 0x72,
	0x72, 0x61, 0x79, 0x41, 0x6e, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2a, 0x50, 0x0a, 0x0a,
	0x57, 0x69, 0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50,
	0x52, 0x4f, 0x54, 0x4f, 0x42, 0x55, 0x46, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x4f,
	0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53, 0x50, 0x42, 0x10, 0x03, 0x12, 0x0f, 0x0a,
	0x0b, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x04, 0x2a, 0x8f,
	0x01, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x14, 0x0a, 0x10, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x5f, 0x54,
	0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x54,
	0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x49, 0x47,
	0x4e, 0x4f, 0x52, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x50, 0x41, 0x52,
	0x53, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x4a,
	0x53, 0x50, 0x42, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x45,
	0x58, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x10, 0x05,
	0x42, 0x21, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conformance_conformance_proto_rawDescOnce sync.Once
	file_conformance_conformance_proto_rawDescData = file_conformance_conformance_proto_rawDesc
)

func file_conformance_conformance_proto_rawDescGZIP() []byte {
	file_conformance_conformance_proto_rawDescOnce.Do(func() {
		file_conformance_conformance_proto_rawDescData = protoimpl.X.CompressGZIP(file_conformance_conformance_proto_rawDescData)
	})
	return file_conformance_conformance_proto_rawDescData
}

var file_conformance_conformance_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_conformance_conformance_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_conformance_conformance_proto_goTypes = []interface{}{
	(WireFormat)(0),             // 0: conformance.WireFormat
	(TestCategory)(0),           // 1: conformance.TestCategory
	(*FailureSet)(nil),          // 2: conformance.FailureSet
	(*ConformanceRequest)(nil),  // 3: conformance.ConformanceRequest
	(*ConformanceResponse)(nil), // 4: conformance.ConformanceResponse
	(*JspbEncodingConfig)(nil),  // 5: conformance.JspbEncodingConfig
}
var file_conformance_conformance_proto_depIdxs = []int32{
	0, // 0: conformance.ConformanceRequest.requested_output_format:type_name -> conformance.WireFormat
	1, // 1: conformance.ConformanceRequest.test_category:type_name -> conformance.TestCategory
	5, // 2: conformance.ConformanceRequest.jspb_encoding_options:type_name -> conformance.JspbEncodingConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_conformance_conformance_proto_init() }
func file_conformance_conformance_proto_init() {
	if File_conformance_conformance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conformance_conformance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FailureSet); i {
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
		file_conformance_conformance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConformanceRequest); i {
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
		file_conformance_conformance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConformanceResponse); i {
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
		file_conformance_conformance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JspbEncodingConfig); i {
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
	file_conformance_conformance_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ConformanceRequest_ProtobufPayload)(nil),
		(*ConformanceRequest_JsonPayload)(nil),
		(*ConformanceRequest_JspbPayload)(nil),
		(*ConformanceRequest_TextPayload)(nil),
	}
	file_conformance_conformance_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ConformanceResponse_ParseError)(nil),
		(*ConformanceResponse_SerializeError)(nil),
		(*ConformanceResponse_RuntimeError)(nil),
		(*ConformanceResponse_ProtobufPayload)(nil),
		(*ConformanceResponse_JsonPayload)(nil),
		(*ConformanceResponse_Skipped)(nil),
		(*ConformanceResponse_JspbPayload)(nil),
		(*ConformanceResponse_TextPayload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conformance_conformance_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conformance_conformance_proto_goTypes,
		DependencyIndexes: file_conformance_conformance_proto_depIdxs,
		EnumInfos:         file_conformance_conformance_proto_enumTypes,
		MessageInfos:      file_conformance_conformance_proto_msgTypes,
	}.Build()
	File_conformance_conformance_proto = out.File
	file_conformance_conformance_proto_rawDesc = nil
	file_conformance_conformance_proto_goTypes = nil
	file_conformance_conformance_proto_depIdxs = nil
}
