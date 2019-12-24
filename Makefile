.DEFAULT_GOAL := all

.SILENT:

all: uuid-build booking-build

uuid-build:
	protoc --go_out=plugins=grpc:. uuid/uuid.proto
	echo "UUID COMPLETE"

booking-build:
	protoc --go_out=plugins=grpc:. booking/booking.proto
	echo "BOOKING COMPLETE"