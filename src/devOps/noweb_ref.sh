#!/bin/sh
df \
|sed '1d' \
|awk '{print $5 " " $6}' | sort -n | tail -1 \
|awk '{print $2}'
