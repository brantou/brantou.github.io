#!/bin/bash

tmux new -s htop \; neww -d 'cd ~/ && htop' \; splitw -d  \; selectp -t 1 \; detach
tmux new -s top \; neww -d 'cd ~/ && top' \; splitw -d  \; selectp -t 1 \; detach
