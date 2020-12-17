sudo ulimit -n 65533
./client -SocketPort=9001  -UdpPort=10001 -WebSocketPort=11001  -ClientStart=1 -ClientEnd=1 -ServerAddress=10.96.8.121