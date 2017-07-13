# grpc-fibonacci

protoc -I proto/ -I. -I $GOPATH/src/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ proto/fibonacci.proto --go_out=plugins=grpc:proto/
protoc -I proto/ -I. -I $GOPATH/src/ -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ --plugin=protoc-gen-grpc-gateway=$GOPATH/bin/protoc-gen-grpc-gateway proto/fibonacci.proto --grpc-gateway_out=logtostderr=true:proto/
