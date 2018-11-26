firewall-cmd --add-port=6379/tcp	# redis
firewall-cmd --add-port=3306/tcp	# mysql
firewall-cmd --add-port=27017/tcp	# mongodb
firewall-cmd --add-port=8123/tcp	# socket
firewall-cmd --add-port=8124/tcp
firewall-cmd --add-port=8089/tcp	#websocket
firewall-cmd --add-port=8090/tcp
./BYGameServerLua