#!/bin/bash

expect -c "
set timeout 5
spawn scp -r root@$1:/opt/minecraft_server/logs ./	
expect \"Password:\"
send \"$2\n\"
interact
"