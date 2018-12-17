// Package grpc provides a gRPC options
package grpc

import (
	"context"
	"crypto/tls"

	"github.com/micro/go-micro/client"
	"github.com/micro/grpc-go/encoding"
)

var (
	// DefaultMaxRecvMsgSize maximum message that client can receive
	// (4 MB).
	DefaultMaxRecvMsgSize = 1024 * 1024 * 4
)

type codecsKey struct{}
type tlsAuth struct{}
type maxRecvMsgSizeKey struct{}

// gRPC Codec to be used to encode/decode requests for a given content type
func Codec(contentType string, c encoding.Codec) client.Option {
	return func(o *client.Options) {
		codecs := make(map[string]encoding.Codec)
		if o.Context == nil {
			o.Context = context.Background()
		}
		if v := o.Context.Value(codecsKey{}); v != nil {
			codecs = v.(map[string]encoding.Codec)
		}
		codecs[contentType] = c
		o.Context = context.WithValue(o.Context, codecsKey{}, codecs)
	}
}

// AuthTLS should be used to setup a secure authentication using TLS
func AuthTLS(t *tls.Config) client.Option {
	return func(o *client.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, tlsAuth{}, t)
	}
}

//
// MaxRecvMsgSize set the maximum size of message that client can receive.
//
func MaxRecvMsgSize(s int) client.Option {
	return func(o *client.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, maxRecvMsgSizeKey{}, s)
	}
}
