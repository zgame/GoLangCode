#ulimit -n 65533
nohup ./SandRockServer -SocketPort=9001 -UdpPort=10001 -WebSocketPort=11001   -ServerTypeName=Game >/dev/null 2>/dev/null &