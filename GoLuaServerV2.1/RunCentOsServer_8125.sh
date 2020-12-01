firewall-cmd --add-port=6379/tcp	# redis
firewall-cmd --add-port=3306/tcp	# mySql
firewall-cmd --add-port=27017/tcp	# mongodb
firewall-cmd --add-port=8125/tcp	# socket
firewall-cmd --add-port=8091/tcp	#websocket
ulimit -n 65533
./BYGameServerLua -WebSocketPort=8091 -SocketPort=8125