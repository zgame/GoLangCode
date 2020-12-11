
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. Proto_User.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. Proto_Game_CCC.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. Proto_Server.proto


pause