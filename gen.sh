#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

RETVAL=0
ROOT=$PWD

ALIAS="Mgoogle/api/annotations.proto=github.com/gengo/grpc-gateway/third_party/googleapis/google/api,"
ALIAS+="Mapi/dtypes/types.proto=github.com/appscode/api/dtypes,"
ALIAS+="Mapi/ci/v0.1/slave.proto=github.com/appscode/api/ci/v0.1,"
ALIAS+="Mapi/kubernetes/v0.1/clusters.proto=github.com/appscode/api/kubernetes/v0.1,"
ALIAS+="Mapi/ssh/v0.1/ssh.proto=github.com/appscode/api/ssh/v0.1,"
ALIAS+="Mgoogle/protobuf/any.proto=github.com/golang/protobuf/ptypes/any"

clean() {
	rm -rf ./*/*/*.pb.go ./*/*/*.pb.gw.go
	rm -rf ./**/*.pb.go ./**/*.pb.gw.go

	rm -rf ./*/*/*.py ./*/*/*.py
	rm -rf ./**/*.py ./**/*.py
}

gen_proto() {
  rm -rf *.pb.go
  protoc -I /usr/local/include -I . \
         -I ${GOPATH}/src/github.com/appscode \
         -I ${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
         -I ${GOPATH}/src/github.com/google/googleapis/google \
         --go_out=plugins=grpc,${ALIAS}:. *.proto
}

gen_gateway_proto() {
  rm -rf *.pb.gw.go
  protoc -I /usr/local/include -I . \
         -I ${GOPATH}/src/github.com/appscode \
         -I ${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
         -I ${GOPATH}/src/github.com/google/googleapis/google \
         --grpc-gateway_out=logtostderr=true,${ALIAS}:. *.proto
}

gen_swagger_def() {
  rm -rf *.swagger.json
  protoc -I /usr/local/include -I . \
         -I ${GOPATH}/src/github.com/appscode \
         -I ${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
         -I ${GOPATH}/src/github.com/google/googleapis/google \
         --swagger_out=logtostderr=true,${ALIAS}:. *.proto
}

gen_server_protos() {
	echo "Generating server protobuf files"
    for d in */ ; do
        pushd ${d}
        if [ -f *.proto ]; then
          gen_proto
        fi
        if [ -d */ ]; then
          for dd in */ ; do
            pushd ${dd}
            gen_proto
            popd
          done
        fi
        popd
    done
}

gen_proxy_protos() {
    echo "Generating gateway protobuf files"
    for d in */ ; do
        pushd ${d}
        if [ -f *.proto ]; then
          gen_gateway_proto
        fi
         if [ -d */ ]; then
          for dd in */ ; do
            pushd ${dd}
            gen_gateway_proto
            popd
          done
        fi
        popd
    done
}

gen_swagger_defs() {
    echo "Generating swagger api definition files"
    for d in */ ; do
        pushd ${d}
        if [ -f *.proto ]; then
          gen_swagger_def
        fi
         if [ -d */ ]; then
          for dd in */ ; do
            pushd ${dd}
            gen_swagger_def
            popd
          done
        fi
        popd
    done
}

gen_py() {
  rm -rf *.py
  protoc -I /usr/local/include -I . \
         -I ${GOPATH}/src/github.com/appscode \
         -I ${GOPATH}/src/github.com/gengo/grpc-gateway/third_party/googleapis \
         -I ${GOPATH}/src/github.com/google/googleapis/google \
         --python_out=plugins=grpc,${ALIAS}:. *.proto
}

gen_python_proto() {
  for d in */ ; do
        pushd ${d}
        if [ -f *.proto ]; then
          gen_py
        fi
        if [ -d */ ]; then
          for dd in */ ; do
            pushd ${dd}
            gen_py
            popd
          done
        fi
        popd
    done
}

compile() {
    echo "compiling files"
    go install ./...
}

gen_protos() {
    clean
    gen_server_protos
    gen_proxy_protos
    gen_swagger_defs
    compile
}

if [ $# -eq 0 ]; then
	gen_protos
	exit $RETVAL
fi

case "$1" in
  compile)
      compile
      ;;
	server)
		gen_server_protos
		;;
	proxy)
		gen_proxy_protos
		;;
	all)
		gen_protos
		;;
	clean)
	  clean
	  ;;
	py)
	  gen_python_proto
	  ;;
	*)  echo $"Usage: $0 {compile|server|proxy|all|clean}"
		RETVAL=1
		;;
esac
exit $RETVAL