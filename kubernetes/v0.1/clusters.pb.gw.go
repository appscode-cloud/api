// Code generated by protoc-gen-grpc-gateway
// source: clusters.proto
// DO NOT EDIT!

/*
Package kubernetes is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package kubernetes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/appscode/api/dtypes"
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

func request_Clusters_Describe_0(ctx context.Context, client ClustersClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClusterDescribeRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["name"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "name")
	}

	protoReq.Name, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	msg, err := client.Describe(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clusters_Create_0(ctx context.Context, client ClustersClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClusterCreateRequest
	var metadata runtime.ServerMetadata

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Create(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clusters_Scale_0(ctx context.Context, client ClustersClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClusterScaleRequest
	var metadata runtime.ServerMetadata

	if err := json.NewDecoder(req.Body).Decode(&protoReq); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Scale(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

var (
	filter_Clusters_Delete_0 = &utilities.DoubleArray{Encoding: map[string]int{"name": 0}, Base: []int{1, 1, 0}, Check: []int{0, 1, 2}}
)

func request_Clusters_Delete_0(ctx context.Context, client ClustersClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClusterDeleteRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["name"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "name")
	}

	protoReq.Name, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_Clusters_Delete_0); err != nil {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Delete(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clusters_List_0(ctx context.Context, client ClustersClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq dtypes.VoidRequest
	var metadata runtime.ServerMetadata

	msg, err := client.List(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clusters_StartupScript_0(ctx context.Context, client ClustersClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClusterStartupScriptRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["role"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "role")
	}

	protoReq.Role, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	msg, err := client.StartupScript(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_Clusters_Config_0(ctx context.Context, client ClustersClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ClusterConfigRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["name"]
	if !ok {
		return nil, metadata, grpc.Errorf(codes.InvalidArgument, "missing parameter %s", "name")
	}

	protoReq.Name, err = runtime.String(val)

	if err != nil {
		return nil, metadata, err
	}

	msg, err := client.Config(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterClustersHandlerFromEndpoint is same as RegisterClustersHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterClustersHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterClustersHandler(ctx, mux, conn)
}

// RegisterClustersHandler registers the http handlers for service Clusters to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterClustersHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewClustersClient(conn)

	mux.Handle("GET", pattern_Clusters_Describe_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Clusters_Describe_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Clusters_Describe_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("PUT", pattern_Clusters_Create_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Clusters_Create_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Clusters_Create_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_Clusters_Scale_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Clusters_Scale_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Clusters_Scale_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("DELETE", pattern_Clusters_Delete_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Clusters_Delete_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Clusters_Delete_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clusters_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Clusters_List_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Clusters_List_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clusters_StartupScript_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Clusters_StartupScript_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Clusters_StartupScript_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_Clusters_Config_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(ctx)
		closeNotifier, ok := w.(http.CloseNotifier)
		if ok {
			go func() {
				<-closeNotifier.CloseNotify()
				cancel()
			}()
		}
		resp, md, err := request_Clusters_Config_0(runtime.AnnotateContext(ctx, req), client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, w, req, err)
			return
		}

		forward_Clusters_Config_0(ctx, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_Clusters_Describe_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "kubernetes", "v0.1", "clusters", "name"}, ""))

	pattern_Clusters_Create_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "kubernetes", "v0.1", "clusters"}, ""))

	pattern_Clusters_Scale_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "kubernetes", "v0.1", "clusters"}, ""))

	pattern_Clusters_Delete_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "kubernetes", "v0.1", "clusters", "name"}, ""))

	pattern_Clusters_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "kubernetes", "v0.1", "clusters"}, ""))

	pattern_Clusters_StartupScript_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "kubernetes", "v0.1", "cluster-startup-script", "role"}, ""))

	pattern_Clusters_Config_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 1, 0, 4, 1, 5, 5}, []string{"api", "kubernetes", "v0.1", "clusters", "config", "name"}, ""))
)

var (
	forward_Clusters_Describe_0 = runtime.ForwardResponseMessage

	forward_Clusters_Create_0 = runtime.ForwardResponseMessage

	forward_Clusters_Scale_0 = runtime.ForwardResponseMessage

	forward_Clusters_Delete_0 = runtime.ForwardResponseMessage

	forward_Clusters_List_0 = runtime.ForwardResponseMessage

	forward_Clusters_StartupScript_0 = runtime.ForwardResponseMessage

	forward_Clusters_Config_0 = runtime.ForwardResponseMessage
)
