#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import math

n = int(input())

for _ in range(n):
    x1, y1, r1, x2, y2, r2 = map(int, input().split())

    d = math.sqrt(pow(x1-x2, 2) + pow(y1-y2, 2))
    rm = abs(r1-r2)
    rp = abs(r1+r2)

    if x1 == x2 and y1 == y2:
        if r1 == r2:
            print(-1)
        else:
            print(0)
    elif rm < d and rp > d:
        print(2)
    elif rm == d or rp == d:
        print(1)
    else:
        print(0)
