.SILENT:
all: go node

go: uuid.proto
	protoc --go_out=plugins=grpc:. uuid.proto

node: uuid.proto
	protoc --js_out=import_style=commonjs,binary:. uuid.proto