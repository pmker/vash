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

// Code generated by protoc-gen-micro. DO NOT EDIT.
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
import _ "github.com/pmker/vash/pkg/common/proto/jobs"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SyncEndpoint service

type SyncEndpointClient interface {
	TriggerResync(ctx context.Context, in *ResyncRequest, opts ...client.CallOption) (*ResyncResponse, error)
}

type syncEndpointClient struct {
	c           client.Client
	serviceName string
}

func NewSyncEndpointClient(serviceName string, c client.Client) SyncEndpointClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "sync"
	}
	return &syncEndpointClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *syncEndpointClient) TriggerResync(ctx context.Context, in *ResyncRequest, opts ...client.CallOption) (*ResyncResponse, error) {
	req := c.c.NewRequest(c.serviceName, "SyncEndpoint.TriggerResync", in)
	out := new(ResyncResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SyncEndpoint service

type SyncEndpointHandler interface {
	TriggerResync(context.Context, *ResyncRequest, *ResyncResponse) error
}

func RegisterSyncEndpointHandler(s server.Server, hdlr SyncEndpointHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&SyncEndpoint{hdlr}, opts...))
}

type SyncEndpoint struct {
	SyncEndpointHandler
}

func (h *SyncEndpoint) TriggerResync(ctx context.Context, in *ResyncRequest, out *ResyncResponse) error {
	return h.SyncEndpointHandler.TriggerResync(ctx, in, out)
}
