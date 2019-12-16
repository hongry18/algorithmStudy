#!/bin/python3

import os
import sys

arr = []

loop_cnt = 0
m = 0
try:
    loop_cnt = int(sys.stdin.readline().split()[0])
except:
    sys.exit(0)

for _ in range(loop_cnt):
    items = sys.stdin.readline().split()

    if items[0] == '1':
        for item_idx in range(len(items[1])):
            arr.append(items[1][item_idx])

    elif items[0] == '2':
        p_idx = int(items[1]) - 1
        pass
    elif items[0] == '3':
        p_idx = int(items[1]) - 1
        print(arr[p_idx])
        pass
    else:
        pass
