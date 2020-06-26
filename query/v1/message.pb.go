// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/query/v1/message.proto

package query

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "go.temporal.io/temporal-proto/common/v1"
	v11 "go.temporal.io/temporal-proto/enums/v1"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type WorkflowQuery struct {
	QueryType string       `protobuf:"bytes,1,opt,name=query_type,json=queryType,proto3" json:"query_type,omitempty"`
	QueryArgs *v1.Payloads `protobuf:"bytes,2,opt,name=query_args,json=queryArgs,proto3" json:"query_args,omitempty"`
}

func (m *WorkflowQuery) Reset()      { *m = WorkflowQuery{} }
func (*WorkflowQuery) ProtoMessage() {}
func (*WorkflowQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e8a4ce3fb63cd22, []int{0}
}
func (m *WorkflowQuery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkflowQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkflowQuery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkflowQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowQuery.Merge(m, src)
}
func (m *WorkflowQuery) XXX_Size() int {
	return m.Size()
}
func (m *WorkflowQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowQuery.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowQuery proto.InternalMessageInfo

func (m *WorkflowQuery) GetQueryType() string {
	if m != nil {
		return m.QueryType
	}
	return ""
}

func (m *WorkflowQuery) GetQueryArgs() *v1.Payloads {
	if m != nil {
		return m.QueryArgs
	}
	return nil
}

type WorkflowQueryResult struct {
	ResultType   v11.QueryResultType `protobuf:"varint,1,opt,name=result_type,json=resultType,proto3,enum=temporal.enums.v1.QueryResultType" json:"result_type,omitempty"`
	Answer       *v1.Payloads        `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	ErrorMessage string              `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (m *WorkflowQueryResult) Reset()      { *m = WorkflowQueryResult{} }
func (*WorkflowQueryResult) ProtoMessage() {}
func (*WorkflowQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e8a4ce3fb63cd22, []int{1}
}
func (m *WorkflowQueryResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkflowQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkflowQueryResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkflowQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowQueryResult.Merge(m, src)
}
func (m *WorkflowQueryResult) XXX_Size() int {
	return m.Size()
}
func (m *WorkflowQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowQueryResult proto.InternalMessageInfo

func (m *WorkflowQueryResult) GetResultType() v11.QueryResultType {
	if m != nil {
		return m.ResultType
	}
	return v11.QUERY_RESULT_TYPE_UNSPECIFIED
}

func (m *WorkflowQueryResult) GetAnswer() *v1.Payloads {
	if m != nil {
		return m.Answer
	}
	return nil
}

func (m *WorkflowQueryResult) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type QueryRejected struct {
	Status v11.WorkflowExecutionStatus `protobuf:"varint,1,opt,name=status,proto3,enum=temporal.enums.v1.WorkflowExecutionStatus" json:"status,omitempty"`
}

func (m *QueryRejected) Reset()      { *m = QueryRejected{} }
func (*QueryRejected) ProtoMessage() {}
func (*QueryRejected) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e8a4ce3fb63cd22, []int{2}
}
func (m *QueryRejected) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRejected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRejected.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRejected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRejected.Merge(m, src)
}
func (m *QueryRejected) XXX_Size() int {
	return m.Size()
}
func (m *QueryRejected) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRejected.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRejected proto.InternalMessageInfo

func (m *QueryRejected) GetStatus() v11.WorkflowExecutionStatus {
	if m != nil {
		return m.Status
	}
	return v11.WORKFLOW_EXECUTION_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterType((*WorkflowQuery)(nil), "temporal.query.v1.WorkflowQuery")
	proto.RegisterType((*WorkflowQueryResult)(nil), "temporal.query.v1.WorkflowQueryResult")
	proto.RegisterType((*QueryRejected)(nil), "temporal.query.v1.QueryRejected")
}

func init() { proto.RegisterFile("temporal/query/v1/message.proto", fileDescriptor_3e8a4ce3fb63cd22) }

var fileDescriptor_3e8a4ce3fb63cd22 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x31, 0x4f, 0xf2, 0x40,
	0x18, 0xee, 0x7d, 0x5f, 0x42, 0xc2, 0x01, 0x26, 0x56, 0x07, 0x42, 0xe4, 0x24, 0x75, 0x21, 0x46,
	0xdb, 0x80, 0x6e, 0x4c, 0x62, 0x1c, 0x4d, 0xb0, 0x98, 0x98, 0xb8, 0x90, 0x0a, 0x67, 0x83, 0xb4,
	0xbd, 0x7a, 0x77, 0x05, 0xbb, 0x39, 0x3a, 0xfa, 0x33, 0xfc, 0x03, 0xfe, 0x07, 0x47, 0x46, 0x46,
	0x29, 0x8b, 0x23, 0x3f, 0xc1, 0xf4, 0xb8, 0x16, 0x4c, 0x1d, 0xdc, 0xde, 0x3c, 0xef, 0xf3, 0xbc,
	0xcf, 0xf3, 0xde, 0xbd, 0x70, 0x9f, 0x63, 0xd7, 0x27, 0xd4, 0x72, 0x8c, 0xc7, 0x00, 0xd3, 0xd0,
	0x18, 0x37, 0x0c, 0x17, 0x33, 0x66, 0xd9, 0x58, 0xf7, 0x29, 0xe1, 0x44, 0xdd, 0x4e, 0x08, 0xba,
	0x20, 0xe8, 0xe3, 0x46, 0xa5, 0x9a, 0x6a, 0xb0, 0x17, 0xb8, 0x2c, 0xd6, 0xac, 0x7a, 0x42, 0x51,
	0xa9, 0x65, 0xdb, 0x13, 0x42, 0x47, 0xf7, 0x0e, 0x99, 0x64, 0x18, 0x7d, 0xe2, 0xba, 0xc4, 0xcb,
	0xb8, 0x6a, 0x23, 0x58, 0xba, 0x91, 0x9a, 0xab, 0x78, 0xb4, 0x5a, 0x85, 0x50, 0x78, 0xf4, 0x78,
	0xe8, 0xe3, 0x32, 0xa8, 0x81, 0x7a, 0xde, 0xcc, 0x0b, 0xe4, 0x3a, 0xf4, 0xb1, 0xda, 0x4a, 0xda,
	0x16, 0xb5, 0x59, 0xf9, 0x5f, 0x0d, 0xd4, 0x0b, 0xcd, 0x3d, 0x3d, 0x8d, 0xbe, 0xb2, 0xd1, 0xc7,
	0x0d, 0xbd, 0x63, 0x85, 0x0e, 0xb1, 0x06, 0x4c, 0x8a, 0xcf, 0xa8, 0xcd, 0xb4, 0x77, 0x00, 0x77,
	0x7e, 0xb8, 0x99, 0x98, 0x05, 0x0e, 0x57, 0xcf, 0x61, 0x81, 0x8a, 0x6a, 0x6d, 0xba, 0xd5, 0xd4,
	0xd6, 0x53, 0xc5, 0x7a, 0xf1, 0xd0, 0x0d, 0x51, 0x9c, 0xc6, 0x84, 0x34, 0xad, 0xd5, 0x53, 0x98,
	0xb3, 0x3c, 0x36, 0xc1, 0xf4, 0x4f, 0xa9, 0x24, 0x57, 0x3d, 0x80, 0x25, 0x4c, 0x29, 0xa1, 0x3d,
	0xf9, 0x2c, 0xe5, 0xff, 0x62, 0xe3, 0xa2, 0x00, 0x2f, 0x57, 0x98, 0xd6, 0x85, 0x25, 0xe9, 0xfc,
	0x80, 0xfb, 0x1c, 0x0f, 0xd4, 0x36, 0xcc, 0x31, 0x6e, 0xf1, 0x80, 0xc9, 0xac, 0x87, 0xbf, 0x64,
	0x4d, 0x16, 0xbd, 0x78, 0xc2, 0xfd, 0x80, 0x0f, 0x89, 0xd7, 0x15, 0x0a, 0x53, 0x2a, 0xdb, 0x2f,
	0x60, 0x3a, 0x47, 0xca, 0x6c, 0x8e, 0x94, 0xe5, 0x1c, 0x81, 0xe7, 0x08, 0x81, 0xb7, 0x08, 0x81,
	0x8f, 0x08, 0x81, 0x69, 0x84, 0xc0, 0x67, 0x84, 0xc0, 0x57, 0x84, 0x94, 0x65, 0x84, 0xc0, 0xeb,
	0x02, 0x29, 0xd3, 0x05, 0x52, 0x66, 0x0b, 0xa4, 0xc0, 0xdd, 0x21, 0xd1, 0x33, 0x97, 0xd2, 0x2e,
	0xca, 0xa8, 0x9d, 0xf8, 0x53, 0x3b, 0xe0, 0xf6, 0xc8, 0xde, 0x60, 0x0d, 0x89, 0x91, 0xd4, 0xc7,
	0xe2, 0xd7, 0xd3, 0x13, 0x6c, 0x89, 0xe2, 0x2e, 0x27, 0xd0, 0x93, 0xef, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x0c, 0xd4, 0x92, 0x2d, 0xa4, 0x02, 0x00, 0x00,
}

func (this *WorkflowQuery) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkflowQuery)
	if !ok {
		that2, ok := that.(WorkflowQuery)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.QueryType != that1.QueryType {
		return false
	}
	if !this.QueryArgs.Equal(that1.QueryArgs) {
		return false
	}
	return true
}
func (this *WorkflowQueryResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkflowQueryResult)
	if !ok {
		that2, ok := that.(WorkflowQueryResult)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ResultType != that1.ResultType {
		return false
	}
	if !this.Answer.Equal(that1.Answer) {
		return false
	}
	if this.ErrorMessage != that1.ErrorMessage {
		return false
	}
	return true
}
func (this *QueryRejected) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryRejected)
	if !ok {
		that2, ok := that.(QueryRejected)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	return true
}
func (this *WorkflowQuery) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&query.WorkflowQuery{")
	s = append(s, "QueryType: "+fmt.Sprintf("%#v", this.QueryType)+",\n")
	if this.QueryArgs != nil {
		s = append(s, "QueryArgs: "+fmt.Sprintf("%#v", this.QueryArgs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *WorkflowQueryResult) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&query.WorkflowQueryResult{")
	s = append(s, "ResultType: "+fmt.Sprintf("%#v", this.ResultType)+",\n")
	if this.Answer != nil {
		s = append(s, "Answer: "+fmt.Sprintf("%#v", this.Answer)+",\n")
	}
	s = append(s, "ErrorMessage: "+fmt.Sprintf("%#v", this.ErrorMessage)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *QueryRejected) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&query.QueryRejected{")
	s = append(s, "Status: "+fmt.Sprintf("%#v", this.Status)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *WorkflowQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkflowQuery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkflowQuery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.QueryArgs != nil {
		{
			size, err := m.QueryArgs.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.QueryType) > 0 {
		i -= len(m.QueryType)
		copy(dAtA[i:], m.QueryType)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.QueryType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WorkflowQueryResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkflowQueryResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkflowQueryResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ErrorMessage) > 0 {
		i -= len(m.ErrorMessage)
		copy(dAtA[i:], m.ErrorMessage)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.ErrorMessage)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Answer != nil {
		{
			size, err := m.Answer.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ResultType != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.ResultType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryRejected) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRejected) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRejected) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorkflowQuery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.QueryType)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.QueryArgs != nil {
		l = m.QueryArgs.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *WorkflowQueryResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ResultType != 0 {
		n += 1 + sovMessage(uint64(m.ResultType))
	}
	if m.Answer != nil {
		l = m.Answer.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.ErrorMessage)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *QueryRejected) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovMessage(uint64(m.Status))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *WorkflowQuery) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WorkflowQuery{`,
		`QueryType:` + fmt.Sprintf("%v", this.QueryType) + `,`,
		`QueryArgs:` + strings.Replace(fmt.Sprintf("%v", this.QueryArgs), "Payloads", "v1.Payloads", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *WorkflowQueryResult) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WorkflowQueryResult{`,
		`ResultType:` + fmt.Sprintf("%v", this.ResultType) + `,`,
		`Answer:` + strings.Replace(fmt.Sprintf("%v", this.Answer), "Payloads", "v1.Payloads", 1) + `,`,
		`ErrorMessage:` + fmt.Sprintf("%v", this.ErrorMessage) + `,`,
		`}`,
	}, "")
	return s
}
func (this *QueryRejected) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&QueryRejected{`,
		`Status:` + fmt.Sprintf("%v", this.Status) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *WorkflowQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WorkflowQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkflowQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryArgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.QueryArgs == nil {
				m.QueryArgs = &v1.Payloads{}
			}
			if err := m.QueryArgs.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WorkflowQueryResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WorkflowQueryResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkflowQueryResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResultType", wireType)
			}
			m.ResultType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResultType |= v11.QueryResultType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Answer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Answer == nil {
				m.Answer = &v1.Payloads{}
			}
			if err := m.Answer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrorMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrorMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryRejected) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRejected: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRejected: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= v11.WorkflowExecutionStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)