#!/usr/bin/env python3
# -*- conding: utf-8 -*-

import sys

try:
    n = int(input())
except Exception as e:
    sys.exit()

z = [1, 0, 1]
o = [0, 1, 1]
def count_fibonacci(n):
    cnt = len(z)
    if cnt <= n:
        for i in range(cnt, n+1):
            z.append(z[i-1] + z[i-2])
            o.append(o[i-1] + o[i-2])

    print('%d %d' % (z[n], o[n]))

for _ in range(n):
    try:
        x = int(input())
        count_fibonacci(x)
    except Exception as e:
        print(e)
        continue
