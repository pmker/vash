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
// source: chat.proto

/*
Package chat is a generated protocol buffer package.

It is generated from these files:
	chat.proto

It has these top-level messages:
	ChatRoom
	ChatMessage
	PutRoomRequest
	PutRoomResponse
	PostMessageRequest
	PostMessageResponse
	DeleteMessageRequest
	DeleteMessageResponse
	ListMessagesRequest
	ListMessagesResponse
	ListRoomsRequest
	ListRoomsResponse
	DeleteRoomRequest
	DeleteRoomResponse
	ChatEvent
	WebSocketMessage
*/
package chat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/pmker/vash/pkg/common/proto/activity"

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

// Client API for ChatService service

type ChatServiceClient interface {
	PutRoom(ctx context.Context, in *PutRoomRequest, opts ...client.CallOption) (*PutRoomResponse, error)
	DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...client.CallOption) (*DeleteRoomResponse, error)
	ListRooms(ctx context.Context, in *ListRoomsRequest, opts ...client.CallOption) (ChatService_ListRoomsClient, error)
	ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...client.CallOption) (ChatService_ListMessagesClient, error)
	PostMessage(ctx context.Context, in *PostMessageRequest, opts ...client.CallOption) (*PostMessageResponse, error)
	DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...client.CallOption) (*DeleteMessageResponse, error)
}

type chatServiceClient struct {
	c           client.Client
	serviceName string
}

func NewChatServiceClient(serviceName string, c client.Client) ChatServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "chat"
	}
	return &chatServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *chatServiceClient) PutRoom(ctx context.Context, in *PutRoomRequest, opts ...client.CallOption) (*PutRoomResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ChatService.PutRoom", in)
	out := new(PutRoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, opts ...client.CallOption) (*DeleteRoomResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ChatService.DeleteRoom", in)
	out := new(DeleteRoomResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ListRooms(ctx context.Context, in *ListRoomsRequest, opts ...client.CallOption) (ChatService_ListRoomsClient, error) {
	req := c.c.NewRequest(c.serviceName, "ChatService.ListRooms", &ListRoomsRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &chatServiceListRoomsClient{stream}, nil
}

type ChatService_ListRoomsClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ListRoomsResponse, error)
}

type chatServiceListRoomsClient struct {
	stream client.Streamer
}

func (x *chatServiceListRoomsClient) Close() error {
	return x.stream.Close()
}

func (x *chatServiceListRoomsClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *chatServiceListRoomsClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *chatServiceListRoomsClient) Recv() (*ListRoomsResponse, error) {
	m := new(ListRoomsResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) ListMessages(ctx context.Context, in *ListMessagesRequest, opts ...client.CallOption) (ChatService_ListMessagesClient, error) {
	req := c.c.NewRequest(c.serviceName, "ChatService.ListMessages", &ListMessagesRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &chatServiceListMessagesClient{stream}, nil
}

type ChatService_ListMessagesClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ListMessagesResponse, error)
}

type chatServiceListMessagesClient struct {
	stream client.Streamer
}

func (x *chatServiceListMessagesClient) Close() error {
	return x.stream.Close()
}

func (x *chatServiceListMessagesClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *chatServiceListMessagesClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *chatServiceListMessagesClient) Recv() (*ListMessagesResponse, error) {
	m := new(ListMessagesResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) PostMessage(ctx context.Context, in *PostMessageRequest, opts ...client.CallOption) (*PostMessageResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ChatService.PostMessage", in)
	out := new(PostMessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, opts ...client.CallOption) (*DeleteMessageResponse, error) {
	req := c.c.NewRequest(c.serviceName, "ChatService.DeleteMessage", in)
	out := new(DeleteMessageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChatService service

type ChatServiceHandler interface {
	PutRoom(context.Context, *PutRoomRequest, *PutRoomResponse) error
	DeleteRoom(context.Context, *DeleteRoomRequest, *DeleteRoomResponse) error
	ListRooms(context.Context, *ListRoomsRequest, ChatService_ListRoomsStream) error
	ListMessages(context.Context, *ListMessagesRequest, ChatService_ListMessagesStream) error
	PostMessage(context.Context, *PostMessageRequest, *PostMessageResponse) error
	DeleteMessage(context.Context, *DeleteMessageRequest, *DeleteMessageResponse) error
}

func RegisterChatServiceHandler(s server.Server, hdlr ChatServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ChatService{hdlr}, opts...))
}

type ChatService struct {
	ChatServiceHandler
}

func (h *ChatService) PutRoom(ctx context.Context, in *PutRoomRequest, out *PutRoomResponse) error {
	return h.ChatServiceHandler.PutRoom(ctx, in, out)
}

func (h *ChatService) DeleteRoom(ctx context.Context, in *DeleteRoomRequest, out *DeleteRoomResponse) error {
	return h.ChatServiceHandler.DeleteRoom(ctx, in, out)
}

func (h *ChatService) ListRooms(ctx context.Context, stream server.Streamer) error {
	m := new(ListRoomsRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ChatServiceHandler.ListRooms(ctx, m, &chatServiceListRoomsStream{stream})
}

type ChatService_ListRoomsStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ListRoomsResponse) error
}

type chatServiceListRoomsStream struct {
	stream server.Streamer
}

func (x *chatServiceListRoomsStream) Close() error {
	return x.stream.Close()
}

func (x *chatServiceListRoomsStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *chatServiceListRoomsStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *chatServiceListRoomsStream) Send(m *ListRoomsResponse) error {
	return x.stream.Send(m)
}

func (h *ChatService) ListMessages(ctx context.Context, stream server.Streamer) error {
	m := new(ListMessagesRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.ChatServiceHandler.ListMessages(ctx, m, &chatServiceListMessagesStream{stream})
}

type ChatService_ListMessagesStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ListMessagesResponse) error
}

type chatServiceListMessagesStream struct {
	stream server.Streamer
}

func (x *chatServiceListMessagesStream) Close() error {
	return x.stream.Close()
}

func (x *chatServiceListMessagesStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *chatServiceListMessagesStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *chatServiceListMessagesStream) Send(m *ListMessagesResponse) error {
	return x.stream.Send(m)
}

func (h *ChatService) PostMessage(ctx context.Context, in *PostMessageRequest, out *PostMessageResponse) error {
	return h.ChatServiceHandler.PostMessage(ctx, in, out)
}

func (h *ChatService) DeleteMessage(ctx context.Context, in *DeleteMessageRequest, out *DeleteMessageResponse) error {
	return h.ChatServiceHandler.DeleteMessage(ctx, in, out)
}
