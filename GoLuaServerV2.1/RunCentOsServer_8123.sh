firewall-cmd --add-port=6379/tcp	# redis
firewall-cmd --add-port=3306/tcp	# mySql
firewall-cmd --add-port=27017/tcp	# mongodb
firewall-cmd --add-port=9001/tcp	# socket
firewall-cmd --add-port=10001/tcp	#websocket
ulimit -n 65533
./server -SocketPort=9001 -UdpPort=10001 -WebSocketPort=11001   -ServerTypeName=Game