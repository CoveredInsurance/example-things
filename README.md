Example Things
====

Example code for using gRPC with Go and Node.js


### Compiling Go

```sh
protoc \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --proto_path=./things \
    --go_out=plugins=grpc:./go \
    --grpc-gateway_out=logtostderr=true:./go \
    --swagger_out=logtostderr=true:./swagger \
    ./things/*.proto
```


### Compiling Java


Using [protoc](https://developers.google.com/protocol-buffers/docs/javatutorial#compiling-your-protocol-buffers):

```sh
protoc \
    -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --proto_path=./things \
    --java_out=./java \
    ./things/*.proto
```

Using Gradle, see:

* https://grpc.io/docs/quickstart/java.html
* https://github.com/google/protobuf-gradle-plugin


### Unit Tests
