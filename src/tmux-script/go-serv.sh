#!/bin/bash

servs = ("api" "cashier" "notify" "magic" "privacy")
rootPath = $GOPATH/src/git.wdwd.com/nova/api

for serv in "${servs[@]}"
do
    tmux has-session -t $serv
    if [ $? = 0]; then
        continue
    fi

    servPath = $rootPath"/"$serv
    tmux new -s $serv -n serv \; \
         send-keys -t $serv:0 'cd $servPath && bee run' \; \
         splitw -d \; selectp -t 1 \; \
         send-keys -t $serv:0.1 'cd $servPath' C-m \; detach
done
