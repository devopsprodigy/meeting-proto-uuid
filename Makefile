.SILENT:
all: go node

go: uuid.proto
	echo "GO START"
	protoc --go_out=plugins=grpc:. uuid.proto
	echo "GO COMPLETE"

node: uuid.proto
	echo "NODE START"
	protoc --js_out=import_style=commonjs,binary:. uuid.proto
	echo "NODE COMPLETE"
