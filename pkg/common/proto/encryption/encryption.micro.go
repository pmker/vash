// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: encryption.proto

/*
Package encryption is a generated protocol buffer package.

It is generated from these files:
	encryption.proto

It has these top-level messages:
	Export
	Import
	KeyInfo
	Key
	AddKeyRequest
	AddKeyResponse
	GetKeyRequest
	GetKeyResponse
	AdminListKeysRequest
	AdminListKeysResponse
	AdminDeleteKeyRequest
	AdminDeleteKeyResponse
	AdminExportKeyRequest
	AdminExportKeyResponse
	AdminImportKeyRequest
	AdminImportKeyResponse
	AdminCreateKeyRequest
	AdminCreateKeyResponse
	Params
	NodeKey
	DeleteNodeRequest
	DeleteNodeResponse
	SetNodeParamsRequest
	SetNodeParamsResponse
	GetNodeKeyRequest
	GetNodeKeyResponse
	SetNodeKeyRequest
	SetNodeKeyResponse
	DeleteNodeKeyRequest
	DeleteNodeKeyResponse
	DeleteNodeSharedKeyRequest
	DeleteNodeSharedKeyResponse
*/
package encryption

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Client API for UserKeyStore service

type UserKeyStoreClient interface {
	AddKey(ctx context.Context, in *AddKeyRequest, opts ...client.CallOption) (*AddKeyResponse, error)
	GetKey(ctx context.Context, in *GetKeyRequest, opts ...client.CallOption) (*GetKeyResponse, error)
	AdminListKeys(ctx context.Context, in *AdminListKeysRequest, opts ...client.CallOption) (*AdminListKeysResponse, error)
	AdminCreateKey(ctx context.Context, in *AdminCreateKeyRequest, opts ...client.CallOption) (*AdminCreateKeyResponse, error)
	AdminDeleteKey(ctx context.Context, in *AdminDeleteKeyRequest, opts ...client.CallOption) (*AdminDeleteKeyResponse, error)
	AdminExportKey(ctx context.Context, in *AdminExportKeyRequest, opts ...client.CallOption) (*AdminExportKeyResponse, error)
	AdminImportKey(ctx context.Context, in *AdminImportKeyRequest, opts ...client.CallOption) (*AdminImportKeyResponse, error)
}

type userKeyStoreClient struct {
	c           client.Client
	serviceName string
}

func NewUserKeyStoreClient(serviceName string, c client.Client) UserKeyStoreClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "encryption"
	}
	return &userKeyStoreClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userKeyStoreClient) AddKey(ctx context.Context, in *AddKeyRequest, opts ...client.CallOption) (*AddKeyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserKeyStore.AddKey", in)
	out := new(AddKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userKeyStoreClient) GetKey(ctx context.Context, in *GetKeyRequest, opts ...client.CallOption) (*GetKeyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserKeyStore.GetKey", in)
	out := new(GetKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userKeyStoreClient) AdminListKeys(ctx context.Context, in *AdminListKeysRequest, opts ...client.CallOption) (*AdminListKeysResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserKeyStore.AdminListKeys", in)
	out := new(AdminListKeysResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userKeyStoreClient) AdminCreateKey(ctx context.Context, in *AdminCreateKeyRequest, opts ...client.CallOption) (*AdminCreateKeyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserKeyStore.AdminCreateKey", in)
	out := new(AdminCreateKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userKeyStoreClient) AdminDeleteKey(ctx context.Context, in *AdminDeleteKeyRequest, opts ...client.CallOption) (*AdminDeleteKeyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserKeyStore.AdminDeleteKey", in)
	out := new(AdminDeleteKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userKeyStoreClient) AdminExportKey(ctx context.Context, in *AdminExportKeyRequest, opts ...client.CallOption) (*AdminExportKeyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserKeyStore.AdminExportKey", in)
	out := new(AdminExportKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userKeyStoreClient) AdminImportKey(ctx context.Context, in *AdminImportKeyRequest, opts ...client.CallOption) (*AdminImportKeyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "UserKeyStore.AdminImportKey", in)
	out := new(AdminImportKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserKeyStore service

type UserKeyStoreHandler interface {
	AddKey(context.Context, *AddKeyRequest, *AddKeyResponse) error
	GetKey(context.Context, *GetKeyRequest, *GetKeyResponse) error
	AdminListKeys(context.Context, *AdminListKeysRequest, *AdminListKeysResponse) error
	AdminCreateKey(context.Context, *AdminCreateKeyRequest, *AdminCreateKeyResponse) error
	AdminDeleteKey(context.Context, *AdminDeleteKeyRequest, *AdminDeleteKeyResponse) error
	AdminExportKey(context.Context, *AdminExportKeyRequest, *AdminExportKeyResponse) error
	AdminImportKey(context.Context, *AdminImportKeyRequest, *AdminImportKeyResponse) error
}

func RegisterUserKeyStoreHandler(s server.Server, hdlr UserKeyStoreHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserKeyStore{hdlr}, opts...))
}

type UserKeyStore struct {
	UserKeyStoreHandler
}

func (h *UserKeyStore) AddKey(ctx context.Context, in *AddKeyRequest, out *AddKeyResponse) error {
	return h.UserKeyStoreHandler.AddKey(ctx, in, out)
}

func (h *UserKeyStore) GetKey(ctx context.Context, in *GetKeyRequest, out *GetKeyResponse) error {
	return h.UserKeyStoreHandler.GetKey(ctx, in, out)
}

func (h *UserKeyStore) AdminListKeys(ctx context.Context, in *AdminListKeysRequest, out *AdminListKeysResponse) error {
	return h.UserKeyStoreHandler.AdminListKeys(ctx, in, out)
}

func (h *UserKeyStore) AdminCreateKey(ctx context.Context, in *AdminCreateKeyRequest, out *AdminCreateKeyResponse) error {
	return h.UserKeyStoreHandler.AdminCreateKey(ctx, in, out)
}

func (h *UserKeyStore) AdminDeleteKey(ctx context.Context, in *AdminDeleteKeyRequest, out *AdminDeleteKeyResponse) error {
	return h.UserKeyStoreHandler.AdminDeleteKey(ctx, in, out)
}

func (h *UserKeyStore) AdminExportKey(ctx context.Context, in *AdminExportKeyRequest, out *AdminExportKeyResponse) error {
	return h.UserKeyStoreHandler.AdminExportKey(ctx, in, out)
}

func (h *UserKeyStore) AdminImportKey(ctx context.Context, in *AdminImportKeyRequest, out *AdminImportKeyResponse) error {
	return h.UserKeyStoreHandler.AdminImportKey(ctx, in, out)
}

// Client API for NodeKeyManager service

type NodeKeyManagerClient interface {
	DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...client.CallOption) (*DeleteNodeResponse, error)
	SetNodeParams(ctx context.Context, in *SetNodeParamsRequest, opts ...client.CallOption) (*SetNodeParamsResponse, error)
	GetNodeKey(ctx context.Context, in *GetNodeKeyRequest, opts ...client.CallOption) (*GetNodeKeyResponse, error)
	SetNodeKey(ctx context.Context, in *SetNodeKeyRequest, opts ...client.CallOption) (*SetNodeKeyResponse, error)
	DeleteNodeKey(ctx context.Context, in *DeleteNodeKeyRequest, opts ...client.CallOption) (*DeleteNodeKeyResponse, error)
	DeleteNodeSharedKey(ctx context.Context, in *DeleteNodeSharedKeyRequest, opts ...client.CallOption) (*DeleteNodeSharedKeyResponse, error)
}

type nodeKeyManagerClient struct {
	c           client.Client
	serviceName string
}

func NewNodeKeyManagerClient(serviceName string, c client.Client) NodeKeyManagerClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "encryption"
	}
	return &nodeKeyManagerClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *nodeKeyManagerClient) DeleteNode(ctx context.Context, in *DeleteNodeRequest, opts ...client.CallOption) (*DeleteNodeResponse, error) {
	req := c.c.NewRequest(c.serviceName, "NodeKeyManager.DeleteNode", in)
	out := new(DeleteNodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeKeyManagerClient) SetNodeParams(ctx context.Context, in *SetNodeParamsRequest, opts ...client.CallOption) (*SetNodeParamsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "NodeKeyManager.SetNodeParams", in)
	out := new(SetNodeParamsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeKeyManagerClient) GetNodeKey(ctx context.Context, in *GetNodeKeyRequest, opts ...client.CallOption) (*GetNodeKeyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "NodeKeyManager.GetNodeKey", in)
	out := new(GetNodeKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeKeyManagerClient) SetNodeKey(ctx context.Context, in *SetNodeKeyRequest, opts ...client.CallOption) (*SetNodeKeyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "NodeKeyManager.SetNodeKey", in)
	out := new(SetNodeKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeKeyManagerClient) DeleteNodeKey(ctx context.Context, in *DeleteNodeKeyRequest, opts ...client.CallOption) (*DeleteNodeKeyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "NodeKeyManager.DeleteNodeKey", in)
	out := new(DeleteNodeKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeKeyManagerClient) DeleteNodeSharedKey(ctx context.Context, in *DeleteNodeSharedKeyRequest, opts ...client.CallOption) (*DeleteNodeSharedKeyResponse, error) {
	req := c.c.NewRequest(c.serviceName, "NodeKeyManager.DeleteNodeSharedKey", in)
	out := new(DeleteNodeSharedKeyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NodeKeyManager service

type NodeKeyManagerHandler interface {
	DeleteNode(context.Context, *DeleteNodeRequest, *DeleteNodeResponse) error
	SetNodeParams(context.Context, *SetNodeParamsRequest, *SetNodeParamsResponse) error
	GetNodeKey(context.Context, *GetNodeKeyRequest, *GetNodeKeyResponse) error
	SetNodeKey(context.Context, *SetNodeKeyRequest, *SetNodeKeyResponse) error
	DeleteNodeKey(context.Context, *DeleteNodeKeyRequest, *DeleteNodeKeyResponse) error
	DeleteNodeSharedKey(context.Context, *DeleteNodeSharedKeyRequest, *DeleteNodeSharedKeyResponse) error
}

func RegisterNodeKeyManagerHandler(s server.Server, hdlr NodeKeyManagerHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&NodeKeyManager{hdlr}, opts...))
}

type NodeKeyManager struct {
	NodeKeyManagerHandler
}

func (h *NodeKeyManager) DeleteNode(ctx context.Context, in *DeleteNodeRequest, out *DeleteNodeResponse) error {
	return h.NodeKeyManagerHandler.DeleteNode(ctx, in, out)
}

func (h *NodeKeyManager) SetNodeParams(ctx context.Context, in *SetNodeParamsRequest, out *SetNodeParamsResponse) error {
	return h.NodeKeyManagerHandler.SetNodeParams(ctx, in, out)
}

func (h *NodeKeyManager) GetNodeKey(ctx context.Context, in *GetNodeKeyRequest, out *GetNodeKeyResponse) error {
	return h.NodeKeyManagerHandler.GetNodeKey(ctx, in, out)
}

func (h *NodeKeyManager) SetNodeKey(ctx context.Context, in *SetNodeKeyRequest, out *SetNodeKeyResponse) error {
	return h.NodeKeyManagerHandler.SetNodeKey(ctx, in, out)
}

func (h *NodeKeyManager) DeleteNodeKey(ctx context.Context, in *DeleteNodeKeyRequest, out *DeleteNodeKeyResponse) error {
	return h.NodeKeyManagerHandler.DeleteNodeKey(ctx, in, out)
}

func (h *NodeKeyManager) DeleteNodeSharedKey(ctx context.Context, in *DeleteNodeSharedKeyRequest, out *DeleteNodeSharedKeyResponse) error {
	return h.NodeKeyManagerHandler.DeleteNodeSharedKey(ctx, in, out)
}