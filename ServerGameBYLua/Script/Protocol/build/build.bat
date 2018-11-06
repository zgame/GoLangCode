protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. user.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. CMD_Common.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. CMD_Game.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. CMD_GameServer.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. CMD_LoginServer.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. CMD_Monitor.proto

pause