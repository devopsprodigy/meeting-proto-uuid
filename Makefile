.DEFAULT_GOAL := all

.SILENT:

all: uuid-build booking-build

uuid-build:
	protoc -I pkg/uuid --go_out=plugins=grpc:pkg/uuid pkg/uuid/uuid.proto
	echo "UUID COMPLETE"

booking-build:
	protoc -I pkg/booking --go_out=plugins=grpc:pkg/booking pkg/booking/booking.proto
	echo "BOOKING COMPLETE"