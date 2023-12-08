#!/bin/bash

git log --author="Xiang Shi"  --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "KV   : ADD: %s, DEL: %s, TOTAL: %s\n", add, subs, loc }'