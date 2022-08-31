// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/rusprofile.proto

/*
Package api is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package api

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_RusProfile_GetRusProfileData_0(ctx context.Context, marshaler runtime.Marshaler, client RusProfileClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RpRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["inn"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "inn")
	}

	protoReq.Inn, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "inn", err)
	}

	msg, err := client.GetRusProfileData(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterRusProfileHandlerFromEndpoint is same as RegisterRusProfileHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterRusProfileHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterRusProfileHandler(ctx, mux, conn)
}

// RegisterRusProfileHandler registers the http handlers for service RusProfile to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterRusProfileHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterRusProfileHandlerClient(ctx, mux, NewRusProfileClient(conn))
}

// RegisterRusProfileHandlerClient registers the http handlers for service RusProfile
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "RusProfileClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "RusProfileClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "RusProfileClient" to call the correct interceptors.
func RegisterRusProfileHandlerClient(ctx context.Context, mux *runtime.ServeMux, client RusProfileClient) error {

	mux.Handle("GET", pattern_RusProfile_GetRusProfileData_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_RusProfile_GetRusProfileData_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_RusProfile_GetRusProfileData_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_RusProfile_GetRusProfileData_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 1, 0, 4, 1, 5, 1}, []string{"get", "inn"}, ""))
)

var (
	forward_RusProfile_GetRusProfileData_0 = runtime.ForwardResponseMessage
)
