/*
 * Copyright (c) 2018. Abstrium SAS <team (at) pydio.com>
 * This file is part of Pydio Cells.
 *
 * Pydio Cells is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Pydio Cells is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Pydio Cells.  If not, see <http://www.gnu.org/licenses/>.
 *
 * The latest code can be found at <https://pydio.com>.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sync.proto

/*
Package sync is a generated protocol buffer package.

It is generated from these files:
	sync.proto

It has these top-level messages:
	ResyncRequest
	ResyncResponse
*/
package sync

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import jobs "github.com/pmker/vash/common/proto/jobs"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResyncRequest struct {
	Path   string     `protobuf:"bytes,1,opt,name=Path" json:"Path,omitempty"`
	DryRun bool       `protobuf:"varint,2,opt,name=DryRun" json:"DryRun,omitempty"`
	Task   *jobs.Task `protobuf:"bytes,3,opt,name=Task" json:"Task,omitempty"`
}

func (m *ResyncRequest) Reset()                    { *m = ResyncRequest{} }
func (m *ResyncRequest) String() string            { return proto.CompactTextString(m) }
func (*ResyncRequest) ProtoMessage()               {}
func (*ResyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ResyncRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *ResyncRequest) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

func (m *ResyncRequest) GetTask() *jobs.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type ResyncResponse struct {
	Success  bool       `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	JsonDiff string     `protobuf:"bytes,2,opt,name=JsonDiff" json:"JsonDiff,omitempty"`
	Task     *jobs.Task `protobuf:"bytes,3,opt,name=Task" json:"Task,omitempty"`
}

func (m *ResyncResponse) Reset()                    { *m = ResyncResponse{} }
func (m *ResyncResponse) String() string            { return proto.CompactTextString(m) }
func (*ResyncResponse) ProtoMessage()               {}
func (*ResyncResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResyncResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ResyncResponse) GetJsonDiff() string {
	if m != nil {
		return m.JsonDiff
	}
	return ""
}

func (m *ResyncResponse) GetTask() *jobs.Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func init() {
	proto.RegisterType((*ResyncRequest)(nil), "sync.ResyncRequest")
	proto.RegisterType((*ResyncResponse)(nil), "sync.ResyncResponse")
}

func init() { proto.RegisterFile("sync.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x09, 0x44, 0x25, 0x3d, 0x28, 0x83, 0x41, 0x28, 0xca, 0x80, 0xa2, 0x4c, 0x99, 0x12,
	0xa9, 0x48, 0x4c, 0x8c, 0x65, 0x41, 0x0c, 0xc8, 0xed, 0xc6, 0xd4, 0xb8, 0x4e, 0x6a, 0x50, 0x7c,
	0xc1, 0xe7, 0x20, 0xe5, 0xdf, 0xa3, 0x5c, 0x28, 0x52, 0xa7, 0x2e, 0xd6, 0xbd, 0x27, 0xfb, 0x7d,
	0xbe, 0x07, 0x40, 0x83, 0x55, 0x45, 0xe7, 0xd0, 0xa3, 0x08, 0xc7, 0x39, 0x79, 0x6a, 0x8c, 0xdf,
	0xf7, 0x55, 0xa1, 0xb0, 0x2d, 0xbb, 0x61, 0x67, 0xb0, 0x24, 0xed, 0x7e, 0x8c, 0xd2, 0x54, 0x2a,
	0x6c, 0x5b, 0xb4, 0x25, 0xdf, 0x2e, 0x3f, 0xb1, 0x22, 0x3e, 0xa6, 0xd7, 0xd9, 0x07, 0x2c, 0xa4,
	0x1e, 0x13, 0xa4, 0xfe, 0xee, 0x35, 0x79, 0x21, 0x20, 0x7c, 0xdf, 0xfa, 0x7d, 0x1c, 0xa4, 0x41,
	0x3e, 0x97, 0x3c, 0x8b, 0x7b, 0x98, 0xad, 0xdc, 0x20, 0x7b, 0x1b, 0x9f, 0xa7, 0x41, 0x1e, 0xc9,
	0x3f, 0x25, 0x1e, 0x20, 0xdc, 0x6c, 0xe9, 0x2b, 0xbe, 0x48, 0x83, 0xfc, 0x6a, 0x09, 0x05, 0xe7,
	0x8e, 0x8e, 0x64, 0x3f, 0xab, 0xe1, 0xe6, 0x10, 0x4e, 0x1d, 0x5a, 0xd2, 0x22, 0x86, 0xcb, 0x75,
	0xaf, 0x94, 0x26, 0x62, 0x40, 0x24, 0x0f, 0x52, 0x24, 0x10, 0xbd, 0x12, 0xda, 0x95, 0xa9, 0x6b,
	0xa6, 0xcc, 0xe5, 0xbf, 0x3e, 0xc5, 0x59, 0xbe, 0xc1, 0xf5, 0x7a, 0xb0, 0xea, 0xc5, 0xee, 0x3a,
	0x34, 0xd6, 0x8b, 0x67, 0x58, 0x6c, 0x9c, 0x69, 0x1a, 0xed, 0x26, 0xbc, 0xb8, 0x2d, 0xb8, 0xb0,
	0xa3, 0x4d, 0x93, 0xbb, 0x63, 0x73, 0xfa, 0x61, 0x76, 0x56, 0xcd, 0xb8, 0x99, 0xc7, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xdb, 0xa4, 0xc8, 0xff, 0x65, 0x01, 0x00, 0x00,
}