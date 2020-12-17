
protoc -o Proto_Game_CCC.pb.txt Proto_Game_CCC.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. Proto_Game_CCC.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. Proto_Server.proto


pause