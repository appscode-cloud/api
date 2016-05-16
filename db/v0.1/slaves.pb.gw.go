// Code generated by protoc-gen-grpc-gateway
// source: slaves.proto
// DO NOT EDIT!

/*
Package db is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package db

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gengo/grpc-gateway/runtime"
	"github.com/gengo/grpc-gateway/utilities"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = json.Marshal
var _ = utilities.NewDoubleArray

func request_Slaves_Add_0(ctx context.Context, client SlavesClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq SlaveAddRequest
	var metadata runtime.ServerMetadata

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["cluster"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "cluster")
	}

	protoReq.Cluster, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	msg, err := client.Add(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterSlavesHandlerFromEndpoint is same as RegisterSlavesHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterSlavesHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterSlavesHandler(ctx, mux, conn)
}

// RegisterSlavesHandler registers the http handlers for service Slaves to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterSlavesHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewSlavesClient(conn)

	mux.Handle("PUT", pattern_Slaves_Add_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		resp, md, err := request_Slaves_Add_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Slaves_Add_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Slaves_Add_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 1, 0, 4, 1, 5, 3}, []string{"db", "v0.1", "slave", "cluster"}, ""))
)

var (
	forward_Slaves_Add_0 = runtime.ForwardResponseMessage
)
