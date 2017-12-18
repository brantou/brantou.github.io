#!/bin/bash

tmux new -s htop \; \
     send-keys -t htop:0 'cd ~/ && htop' C-m \; \
     splitw -d \; selectp -t 1 \; \
     send-keys -t htop:0.1 'cd ~/ && clear' C-m \; detach
