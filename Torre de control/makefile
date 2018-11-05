c++11FLAG = -std=c++11
OTHER_FLAGS = `pkg-config --cflags protobuf grpc`
ifeq ($(SYSTEM),Darwin)
LDFLAGS += -L/usr/local/lib `pkg-config --libs protobuf grpc++ grpc`\
           -lgrpc++_reflection\
           -ldl
else
LDFLAGS += -L/usr/local/lib `pkg-config --libs protobuf grpc++ grpc`\
           -Wl,--no-as-needed -lgrpc++_reflection -Wl,--as-needed\
           -ldl
endif

run: Tower_Server
	./Tower_Server

Tower_Server: TowerProto.pb.o TowerProto.grpc.pb.o Tower_Server.o
	g++ -Wall $^ $(LDFLAGS) -o $@

Tower_Server.o: TowerProto.pb.o TowerProto.grpc.pb.o
	g++ $(c++11FLAG) $(OTHER_FLAGS) -c -o Tower_Server.o Tower_Server.cpp

TowerProto.grpc.pb.o: compile_proto_grpc
	g++ $(c++11FLAG) $(OTHER_FLAGS) -c -o TowerProto.grpc.pb.o TowerProto.grpc.pb.cc

compile_proto_grpc:
	protoc -I=. --grpc_out=. --plugin=protoc-gen-grpc=`which grpc_cpp_plugin` TowerProto.proto

TowerProto.pb.o: compile_proto_cpp
	g++ $(c++11FLAG) $(OTHER_FLAGS) -c -o TowerProto.pb.o TowerProto.pb.cc

compile_proto_cpp:
	protoc -I=. --cpp_out=. TowerProto.proto
