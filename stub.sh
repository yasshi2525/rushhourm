#!/bin/bash
rm -f pb/*.pb.go
protoc --go_out=plugins=grpc:. pb/*.proto

mkdir -p mock_pb
rm -f mock_pb/*_mock.go
mockgen github.com/yasshi2525/rushhourm/pb ModelServiceClient > mock_pb/model_mock.go