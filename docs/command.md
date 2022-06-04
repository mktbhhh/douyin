客户端生成
kitex -type protobuf -module douyin idl/core.proto
kitex -type protobuf -module douyin idl/user.proto

服务端生成
cd cmd/user
kitex -type protobuf -module douyin -service user -use douyin/kitex_gen -I ../../idl user.proto

cd cmd/core
kitex -type protobuf -module douyin -service core -use douyin/kitex_gen -I ../../idl core.proto

