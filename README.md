把json的key 转化成驼峰格式，去重

可以配合改成proto
https://json-to-proto.github.io/


protoc --go_out=plugins=grpc:. proto/test.proto

go mod init proto_test