#!/bin/bash

CMDS=( "htop" "top" )
CMDPATH="~"

for cmd in "${CMDS[@]}"
do
    tmux new -s $cmd \; \
         send-keys -t $cmd:0 'cd '$CMDPATH' && '$cmd C-m \; \
         splitw -d \; selectp -t 1 \; \
         send-keys -t $cmd:0.1 'cd '$CMDPATH' && clear' C-m \; detach
done
