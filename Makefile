.DEFAULT_GOAL := all

.SILENT:

all: uuid-build booking-build room-build user-build

uuid-build:
	protoc uuid/uuid.proto -I uuid --go_out=plugins=grpc:uuid
	echo "UUID COMPLETE"

booking-build:
	protoc booking/booking.proto -I. --go_out=plugins=grpc:.
	echo "BOOKING COMPLETE"
	
room-build:
	protoc room/room.proto -I. --go_out=plugins=grpc:.
	echo "ROOM COMPLETE"
	
user-build:
	protoc user/user.proto -I user --go_out=plugins=grpc:user
	echo "USER COMPLETE"