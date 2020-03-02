#!/usr/bin/env python3
# -*- coding: utf-8 -*-

n = int(input())

def in_check(x1, x2, y1, y2, r):
    return (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) < r*r

for _ in range(n):
    """
        x1: 시작 x
        y1: 시작 y
        x2: 도착 x
        y2: 도착 y
    """
    x1, y1, x2, y2 = map(int, input().split())

    m = int(input())

    cnt = 0

    for _ in range(m):
        """
            cx: 원 x
            cy: 원 y
            r: 반지름
        """
        cx, cy, r = map(int, input().split())

        c1 = in_check(cx, x1, cy, y1, r)
        c2 = in_check(cx, x2, cy, y2, r)
        if c1 and c2:
            continue
        elif c1 or c2:
            cnt += 1

    print(cnt)
