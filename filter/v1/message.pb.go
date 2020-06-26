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
// source: temporal/filter/v1/message.proto

package filter

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	v1 "go.temporal.io/temporal-proto/enums/v1"
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

type WorkflowExecutionFilter struct {
	WorkflowId string `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId      string `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (m *WorkflowExecutionFilter) Reset()      { *m = WorkflowExecutionFilter{} }
func (*WorkflowExecutionFilter) ProtoMessage() {}
func (*WorkflowExecutionFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b7ad9ad1a3cff40, []int{0}
}
func (m *WorkflowExecutionFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkflowExecutionFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkflowExecutionFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkflowExecutionFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionFilter.Merge(m, src)
}
func (m *WorkflowExecutionFilter) XXX_Size() int {
	return m.Size()
}
func (m *WorkflowExecutionFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionFilter.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionFilter proto.InternalMessageInfo

func (m *WorkflowExecutionFilter) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *WorkflowExecutionFilter) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

type WorkflowTypeFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *WorkflowTypeFilter) Reset()      { *m = WorkflowTypeFilter{} }
func (*WorkflowTypeFilter) ProtoMessage() {}
func (*WorkflowTypeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b7ad9ad1a3cff40, []int{1}
}
func (m *WorkflowTypeFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkflowTypeFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkflowTypeFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkflowTypeFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTypeFilter.Merge(m, src)
}
func (m *WorkflowTypeFilter) XXX_Size() int {
	return m.Size()
}
func (m *WorkflowTypeFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTypeFilter.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTypeFilter proto.InternalMessageInfo

func (m *WorkflowTypeFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StartTimeFilter struct {
	EarliestTime int64 `protobuf:"varint,1,opt,name=earliest_time,json=earliestTime,proto3" json:"earliest_time,omitempty"`
	LatestTime   int64 `protobuf:"varint,2,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"`
}

func (m *StartTimeFilter) Reset()      { *m = StartTimeFilter{} }
func (*StartTimeFilter) ProtoMessage() {}
func (*StartTimeFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b7ad9ad1a3cff40, []int{2}
}
func (m *StartTimeFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StartTimeFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StartTimeFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StartTimeFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartTimeFilter.Merge(m, src)
}
func (m *StartTimeFilter) XXX_Size() int {
	return m.Size()
}
func (m *StartTimeFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_StartTimeFilter.DiscardUnknown(m)
}

var xxx_messageInfo_StartTimeFilter proto.InternalMessageInfo

func (m *StartTimeFilter) GetEarliestTime() int64 {
	if m != nil {
		return m.EarliestTime
	}
	return 0
}

func (m *StartTimeFilter) GetLatestTime() int64 {
	if m != nil {
		return m.LatestTime
	}
	return 0
}

type StatusFilter struct {
	Status v1.WorkflowExecutionStatus `protobuf:"varint,1,opt,name=status,proto3,enum=temporal.enums.v1.WorkflowExecutionStatus" json:"status,omitempty"`
}

func (m *StatusFilter) Reset()      { *m = StatusFilter{} }
func (*StatusFilter) ProtoMessage() {}
func (*StatusFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b7ad9ad1a3cff40, []int{3}
}
func (m *StatusFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StatusFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StatusFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StatusFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusFilter.Merge(m, src)
}
func (m *StatusFilter) XXX_Size() int {
	return m.Size()
}
func (m *StatusFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusFilter.DiscardUnknown(m)
}

var xxx_messageInfo_StatusFilter proto.InternalMessageInfo

func (m *StatusFilter) GetStatus() v1.WorkflowExecutionStatus {
	if m != nil {
		return m.Status
	}
	return v1.WORKFLOW_EXECUTION_STATUS_UNSPECIFIED
}

func init() {
	proto.RegisterType((*WorkflowExecutionFilter)(nil), "temporal.filter.v1.WorkflowExecutionFilter")
	proto.RegisterType((*WorkflowTypeFilter)(nil), "temporal.filter.v1.WorkflowTypeFilter")
	proto.RegisterType((*StartTimeFilter)(nil), "temporal.filter.v1.StartTimeFilter")
	proto.RegisterType((*StatusFilter)(nil), "temporal.filter.v1.StatusFilter")
}

func init() { proto.RegisterFile("temporal/filter/v1/message.proto", fileDescriptor_9b7ad9ad1a3cff40) }

var fileDescriptor_9b7ad9ad1a3cff40 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x31, 0x4b, 0x3b, 0x31,
	0x18, 0xc6, 0x2f, 0xfd, 0xff, 0x2d, 0x18, 0xab, 0xc2, 0x41, 0x51, 0x1c, 0x5e, 0xcb, 0xb9, 0x14,
	0xc1, 0x1c, 0xd5, 0xd1, 0xad, 0xa0, 0xd0, 0x41, 0xa8, 0xd7, 0x42, 0xc1, 0xa5, 0x44, 0x9b, 0x96,
	0xe0, 0xdd, 0xa5, 0xe4, 0x72, 0xad, 0x6e, 0xee, 0x2e, 0x7e, 0x0c, 0x3f, 0x8a, 0x63, 0xc7, 0x8e,
	0x36, 0x5d, 0x1c, 0xfb, 0x11, 0xe4, 0x72, 0x97, 0x56, 0xe8, 0xf6, 0xe6, 0x79, 0x7f, 0xcf, 0x43,
	0x78, 0x1f, 0x5c, 0x53, 0x2c, 0x1a, 0x0b, 0x49, 0x43, 0x7f, 0xc8, 0x43, 0xc5, 0xa4, 0x3f, 0x69,
	0xf8, 0x11, 0x4b, 0x12, 0x3a, 0x62, 0x64, 0x2c, 0x85, 0x12, 0xae, 0x6b, 0x09, 0x92, 0x13, 0x64,
	0xd2, 0x38, 0xd9, 0xb8, 0x58, 0x9c, 0x46, 0x49, 0x66, 0x9a, 0x0a, 0xf9, 0x3c, 0x0c, 0xc5, 0x34,
	0x77, 0x79, 0xf7, 0xf8, 0xa8, 0x57, 0x28, 0x37, 0x2f, 0xec, 0x29, 0x55, 0x5c, 0xc4, 0xb7, 0xc6,
	0xef, 0x9e, 0xe2, 0x3d, 0x0b, 0xf7, 0xf9, 0xe0, 0x18, 0xd5, 0x50, 0x7d, 0x37, 0xc0, 0x56, 0x6a,
	0x0d, 0xdc, 0x2a, 0x2e, 0xcb, 0x34, 0xce, 0x76, 0x25, 0xb3, 0xdb, 0x91, 0x69, 0xdc, 0x1a, 0x78,
	0x75, 0xec, 0xda, 0xc8, 0xee, 0xeb, 0x98, 0x15, 0x69, 0x2e, 0xfe, 0x1f, 0xd3, 0x88, 0x15, 0x31,
	0x66, 0xf6, 0x7a, 0xf8, 0xb0, 0xa3, 0xa8, 0x54, 0x5d, 0x1e, 0x59, 0xec, 0x0c, 0xef, 0x33, 0x2a,
	0x43, 0xce, 0x12, 0xd5, 0x57, 0xbc, 0xe0, 0xff, 0x05, 0x15, 0x2b, 0x66, 0x68, 0xf6, 0xb3, 0x90,
	0xaa, 0x35, 0x52, 0x32, 0x08, 0xce, 0xa5, 0x0c, 0xf0, 0x02, 0x5c, 0xe9, 0x28, 0xaa, 0xd2, 0xa4,
	0x48, 0x6d, 0xe2, 0x72, 0x62, 0xde, 0x26, 0xee, 0xe0, 0xf2, 0x9c, 0xac, 0x8f, 0x65, 0x0e, 0x43,
	0x26, 0x0d, 0xb2, 0x75, 0x86, 0x3c, 0x21, 0x28, 0x9c, 0xcd, 0x77, 0x34, 0x5b, 0x80, 0x33, 0x5f,
	0x80, 0xb3, 0x5a, 0x00, 0x7a, 0xd3, 0x80, 0x3e, 0x35, 0xa0, 0x2f, 0x0d, 0x68, 0xa6, 0x01, 0x7d,
	0x6b, 0x40, 0x3f, 0x1a, 0x9c, 0x95, 0x06, 0xf4, 0xb1, 0x04, 0x67, 0xb6, 0x04, 0x67, 0xbe, 0x04,
	0x07, 0x57, 0xb9, 0x20, 0xdb, 0xcd, 0x34, 0x2b, 0x77, 0x79, 0x79, 0xed, 0xac, 0x85, 0x36, 0x7a,
	0x20, 0xa3, 0x3f, 0x18, 0x17, 0xbe, 0x9d, 0x2f, 0x4c, 0x4d, 0x9b, 0xd2, 0xaf, 0xf3, 0xe9, 0xb1,
	0x6c, 0xf4, 0xab, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0xb6, 0x59, 0xe9, 0x18, 0x02, 0x00,
	0x00,
}

func (this *WorkflowExecutionFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkflowExecutionFilter)
	if !ok {
		that2, ok := that.(WorkflowExecutionFilter)
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
	if this.WorkflowId != that1.WorkflowId {
		return false
	}
	if this.RunId != that1.RunId {
		return false
	}
	return true
}
func (this *WorkflowTypeFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkflowTypeFilter)
	if !ok {
		that2, ok := that.(WorkflowTypeFilter)
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
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *StartTimeFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StartTimeFilter)
	if !ok {
		that2, ok := that.(StartTimeFilter)
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
	if this.EarliestTime != that1.EarliestTime {
		return false
	}
	if this.LatestTime != that1.LatestTime {
		return false
	}
	return true
}
func (this *StatusFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*StatusFilter)
	if !ok {
		that2, ok := that.(StatusFilter)
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
func (this *WorkflowExecutionFilter) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&filter.WorkflowExecutionFilter{")
	s = append(s, "WorkflowId: "+fmt.Sprintf("%#v", this.WorkflowId)+",\n")
	s = append(s, "RunId: "+fmt.Sprintf("%#v", this.RunId)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *WorkflowTypeFilter) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&filter.WorkflowTypeFilter{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StartTimeFilter) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&filter.StartTimeFilter{")
	s = append(s, "EarliestTime: "+fmt.Sprintf("%#v", this.EarliestTime)+",\n")
	s = append(s, "LatestTime: "+fmt.Sprintf("%#v", this.LatestTime)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *StatusFilter) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&filter.StatusFilter{")
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
func (m *WorkflowExecutionFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkflowExecutionFilter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkflowExecutionFilter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RunId) > 0 {
		i -= len(m.RunId)
		copy(dAtA[i:], m.RunId)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.RunId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.WorkflowId) > 0 {
		i -= len(m.WorkflowId)
		copy(dAtA[i:], m.WorkflowId)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.WorkflowId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WorkflowTypeFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkflowTypeFilter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkflowTypeFilter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StartTimeFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StartTimeFilter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StartTimeFilter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestTime != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.LatestTime))
		i--
		dAtA[i] = 0x10
	}
	if m.EarliestTime != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.EarliestTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *StatusFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StatusFilter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StatusFilter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *WorkflowExecutionFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.WorkflowId)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.RunId)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *WorkflowTypeFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *StartTimeFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EarliestTime != 0 {
		n += 1 + sovMessage(uint64(m.EarliestTime))
	}
	if m.LatestTime != 0 {
		n += 1 + sovMessage(uint64(m.LatestTime))
	}
	return n
}

func (m *StatusFilter) Size() (n int) {
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
func (this *WorkflowExecutionFilter) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WorkflowExecutionFilter{`,
		`WorkflowId:` + fmt.Sprintf("%v", this.WorkflowId) + `,`,
		`RunId:` + fmt.Sprintf("%v", this.RunId) + `,`,
		`}`,
	}, "")
	return s
}
func (this *WorkflowTypeFilter) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WorkflowTypeFilter{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StartTimeFilter) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StartTimeFilter{`,
		`EarliestTime:` + fmt.Sprintf("%v", this.EarliestTime) + `,`,
		`LatestTime:` + fmt.Sprintf("%v", this.LatestTime) + `,`,
		`}`,
	}, "")
	return s
}
func (this *StatusFilter) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StatusFilter{`,
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
func (m *WorkflowExecutionFilter) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WorkflowExecutionFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkflowExecutionFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkflowId", wireType)
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
			m.WorkflowId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunId", wireType)
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
			m.RunId = string(dAtA[iNdEx:postIndex])
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
func (m *WorkflowTypeFilter) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WorkflowTypeFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkflowTypeFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
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
func (m *StartTimeFilter) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StartTimeFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StartTimeFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EarliestTime", wireType)
			}
			m.EarliestTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EarliestTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestTime", wireType)
			}
			m.LatestTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestTime |= int64(b&0x7F) << shift
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
func (m *StatusFilter) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: StatusFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StatusFilter: illegal tag %d (wire type %d)", fieldNum, wire)
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
				m.Status |= v1.WorkflowExecutionStatus(b&0x7F) << shift
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