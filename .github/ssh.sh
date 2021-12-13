#!/bin/sh

server=$1
password=$2

expect -c "
set timeout 5
spawn scp -o StrictHostKeyChecking=no -r root@${server}:/opt/minecraft_server/logs ./	
expect \"Password:\"
send \"${password}\n\"
interact
"
