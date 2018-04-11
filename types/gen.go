package types

//go:generate protoc --proto_path=$GOPATH/src:. --twirp_out=. --gogofaster_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types:. ./listOpts.proto ./uuid.proto ./nil.proto ./atTimestamps.proto
