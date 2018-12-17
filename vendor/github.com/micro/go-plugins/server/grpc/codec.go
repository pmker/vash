package grpc

import (
	"encoding/json"
	"fmt"

	"github.com/micro/go-micro/codec"
	"github.com/micro/go-micro/codec/jsonrpc"
	"github.com/micro/go-micro/codec/protorpc"
)

type jsonCodec struct{}
type bytesCodec struct{}

var (
	defaultRPCCodecs = map[string]codec.NewCodec{
		"application/json":         jsonrpc.NewCodec,
		"application/json-rpc":     jsonrpc.NewCodec,
		"application/protobuf":     protorpc.NewCodec,
		"application/proto-rpc":    protorpc.NewCodec,
		"application/octet-stream": protorpc.NewCodec,
	}
)

func (jsonCodec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonCodec) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (jsonCodec) Name() string {
	return "json"
}

func (bytesCodec) Marshal(v interface{}) ([]byte, error) {
	b, ok := v.(*[]byte)
	if !ok {
		return nil, fmt.Errorf("failed to marshal: %v is not type of *[]byte", v)
	}
	return *b, nil
}

func (bytesCodec) Unmarshal(data []byte, v interface{}) error {
	b, ok := v.(*[]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal: %v is not type of *[]byte", v)
	}
	*b = data
	return nil
}

func (bytesCodec) Name() string {
	return "bytes"
}
