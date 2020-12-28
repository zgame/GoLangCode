
protoc -o protoGameSandRock.pb.txt protoGameSandRock.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. protoGameSandRock.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. protoServer.proto


pause