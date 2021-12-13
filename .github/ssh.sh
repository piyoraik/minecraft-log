#!/bin/bash

expect -c "
set timeout 30
spawn scp -o StrictHostKeyChecking=no -r root@$1:/opt/minecraft_server/logs ./	
expect \"Password:\"
send \"$2\n\"
interact
"
