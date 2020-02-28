# -*- coding: utf-8 -*-

import sys

try:
    T = int(input())
except e:
    print(e)
    sys.exit()

for test_case in range(1, T + 1):
    arr = list(map(int, input().split(' ')))
    s = 0
    for item in arr:
        if item % 2 != 0:
            s += item

    print('#{} {}'.format( (test_case+1), s))
