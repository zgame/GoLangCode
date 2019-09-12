
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. CMD_Game_TB.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. CMD_Game.proto
protoc --plugin=protoc-gen-lua="plugin\protoc-gen-lua.bat" --lua_out=. CMD_GameServer.proto

pause