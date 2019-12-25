.DEFAULT_GOAL := all

.SILENT:

all: uuid-build booking-build room-build user-build

uuid-build:
	protoc -I uuid --go_out=plugins=grpc:uuid uuid/uuid.proto
	echo "UUID COMPLETE"

booking-build:
	protoc -I booking --go_out=plugins=grpc:booking booking/booking.proto
	echo "BOOKING COMPLETE"
	
room-build:
	protoc -I room --go_out=plugins=grpc:room room/room.proto
	echo "ROOM COMPLETE"
	
user-build:
	protoc -I user --go_out=plugins=grpc:user user/user.proto
	echo "USER COMPLETE"