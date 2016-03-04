// Code generated by protoc-gen-grpc-gateway
// source: build.proto
// DO NOT EDIT!

/*
Package ci is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package ci

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/gengo/grpc-gateway/utilities"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = json.Marshal
var _ = utilities.NewDoubleArray

var (
	filter_Builds_Describe_0 = &utilities.DoubleArray{Encoding: map[string]int{"job_name": 0, "number": 1}, Base: []int{1, 1, 2, 0, 0}, Check: []int{0, 1, 1, 2, 3}}
)

func request_Builds_Describe_0(ctx context.Context, client BuildsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq BuildDescribeRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["job_name"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "job_name")
	}

	protoReq.JobName, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	val, ok = pathParams["number"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "number")
	}

	protoReq.Number, err = runtime.Int64(val)

	if err != nil {
		return nil, metadata, err
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_Builds_Describe_0); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Describe(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Builds_List_0(ctx context.Context, client BuildsClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq BuildListRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["job_name"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "job_name")
	}

	protoReq.JobName, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	msg, err := client.List(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterBuildsHandlerFromEndpoint is same as RegisterBuildsHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterBuildsHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				glog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterBuildsHandler(ctx, mux, conn)
}

// RegisterBuildsHandler registers the http handlers for service Builds to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterBuildsHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewBuildsClient(conn)

	mux.Handle("GET", pattern_Builds_Describe_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Builds_Describe_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Builds_Describe_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Builds_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Builds_List_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Builds_List_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Builds_Describe_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5}, []string{"api", "ci", "v0", "builds", "job_name", "number"}, ""))

	pattern_Builds_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "ci", "v0", "builds", "job_name"}, ""))
)

var (
	forward_Builds_Describe_0 = runtime.ForwardResponseMessage

	forward_Builds_List_0 = runtime.ForwardResponseMessage
)
