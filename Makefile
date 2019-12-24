.DEFAULT_GOAL := all

.SILENT:

all: uuid-build booking-build

uuid-build:
	protoc -I uuid --go_out=plugins=grpc:uuid uuid/uuid.proto
	echo "UUID COMPLETE"

booking-build:
	protoc -I booking --go_out=plugins=grpc:booking booking/booking.proto
	echo "BOOKING COMPLETE"