// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: watchtowerrpc/watchtower.proto

/*
Package watchtowerrpc is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package watchtowerrpc

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

func request_Watchtower_GetInfo_0(ctx context.Context, marshaler runtime.Marshaler, client WatchtowerClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetInfoRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetInfo(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_Watchtower_GetInfo_0(ctx context.Context, marshaler runtime.Marshaler, server WatchtowerServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetInfoRequest
	var metadata runtime.ServerMetadata

	msg, err := server.GetInfo(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterWatchtowerHandlerServer registers the http handlers for service Watchtower to "mux".
// UnaryRPC     :call WatchtowerServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterWatchtowerHandlerServer(ctx context.Context, mux *runtime.ServeMux, server WatchtowerServer) error {

	mux.Handle("GET", pattern_Watchtower_GetInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_Watchtower_GetInfo_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Watchtower_GetInfo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterWatchtowerHandlerFromEndpoint is same as RegisterWatchtowerHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterWatchtowerHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterWatchtowerHandler(ctx, mux, conn)
}

// RegisterWatchtowerHandler registers the http handlers for service Watchtower to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterWatchtowerHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterWatchtowerHandlerClient(ctx, mux, NewWatchtowerClient(conn))
}

// RegisterWatchtowerHandlerClient registers the http handlers for service Watchtower
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "WatchtowerClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "WatchtowerClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "WatchtowerClient" to call the correct interceptors.
func RegisterWatchtowerHandlerClient(ctx context.Context, mux *runtime.ServeMux, client WatchtowerClient) error {

	mux.Handle("GET", pattern_Watchtower_GetInfo_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_Watchtower_GetInfo_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_Watchtower_GetInfo_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Watchtower_GetInfo_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v2", "watchtower", "server"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_Watchtower_GetInfo_0 = runtime.ForwardResponseMessage
)