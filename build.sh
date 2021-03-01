#!/bin/bash

set -e
# Build uibackend proto
grpc_tools_node_protoc \
    --js_out=import_style=commonjs,binary:uiBackend/proto/ \
    --grpc_out=grpc_js:uiBackend/proto \
    --proto_path=./protos/uibackend ./protos/uibackend/*.proto

# Build all protos for userInterface
#grpc_tools_node_protoc \
#    --js_out=import_style=commonjs,binary:userInterface/proto/ \
#    --grpc_out=grpc_js:userInterface/proto \
#    --proto_path=./protos ./protos/uibackend/*.proto ./protos/notificationservice/*.proto ./protos/igapi/*.proto

#copy uibackend proto file to uibackend folder and userinterface folder
#cp protos/uibackend/uibackend.proto userInterface/proto/
cp protos/uibackend/uibackend.proto uiBackend/proto/

cp protos/gmailapi/gmailapi.proto notificationService/proto/

cp protos/gmailapi/gmailapi.proto apiService/proto/
cp protos/api/api.proto apiService/proto/
cp protos/user/user.proto apiService/proto/
cp protos/uibackend/uibackend.proto apiService/proto/
cp protos/igapi/igapi.proto apiService/proto/

# Build the user service protos for Node
grpc_tools_node_protoc \
    --js_out=import_style=commonjs,binary:userService/proto/ \
    --grpc_out=grpc_js:userService/proto \
    --proto_path=./protos/user ./protos/user/*.proto

#python3 -m grpc_tools.protoc \
#    --python_out=./projectService/proto \
#    --grpc_python_out=./projectService/proto \
#    --proto_path=./protos/project ./protos/project/*.proto

#grpc_tools_ruby_protoc \
#    --proto_path=./protos \
#    --ruby_out=./taskService/proto \
#    --grpc_out=./taskService/proto \
#    ./protos/task/*.proto ./protos/user/*.proto ./protos/project/*.proto

protoc  --go_out=./notificationService/proto --go_opt=paths=source_relative\
    --go-grpc_out=./notificationService/proto --go-grpc_opt=paths=source_relative\
    --proto_path=./protos ./protos/gmailapi/*.proto ./protos/igapi/*.proto
#echo "protoc for notificationservice"

#protoc -I .\
#    --java_out=./igapi/proto \
#    --java-grpc_out=./igapi/proto \
#    --proto_path=./protos/igapi ./protos/igapi/*.proto
#echo "protoc for igapi"

# Fix python import
# https://github.com/protocolbuffers/protobuf/issues/1491
#pushd projectService/proto
#sed -i -E 's/^import.*_pb2/from . \0/' *.py
#popd

# Generate requirements.txt for python project
#pushd projectService
#pipenv lock -r > requirements.txt
#popd


packs=(
    "api"
    "user"
#    "project"
#    "task"
    "gmailapi"
    "igapi"
    "uibackend"
);

# Generate go protobuf files for all protos
# shellcheck disable=SC2068
for d in ${packs[@]} ; do
    echo "Compiling $d";
#    protoc -I . -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
protoc -I . -I./protos/google/api \
            -I protos \
            --go_out=. --go_opt=paths=source_relative \
            --go-grpc_out=. --go-grpc_opt=paths=source_relative \
            protos/$d/*.proto
done

# Generate api definition with GRPC Gateway
#protoc -I . -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
protoc -I . -I./protos/google/api \
    -I protos \
    --grpc-gateway_out . \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    protos/api/api.proto

# Genrate swagger docs with GRPC Gateway
#protoc -I . -I . -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
protoc -I . -I./protos/google/api \
    -I protos \
    --openapiv2_out . \
    --openapiv2_opt logtostderr=true \
    protos/api/api.proto

# Move the swagger json to api service www folder
mv protos/api/api.swagger.json apiService/www