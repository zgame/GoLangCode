
protoc -o protoGameCcc.pb.txt protoGameCcc.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. protoGameCcc.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. protoServer.proto


pause