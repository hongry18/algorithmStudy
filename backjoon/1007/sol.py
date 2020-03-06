#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from sys import stdin, stdout
import math

ax = []
ay = []
c = []
r = [0]

def solv(n, cnt, idx):

    if cnt == n / 2:
        x = 0
        y = 0
        for i in range(0, n):
            if c[i]:
                x -= ax[i]
                y -= ay[i]
            else:
                x += ax[i]
                y += ay[i]

        r[0] = min(r[0], math.sqrt(x*x + y*y))
        return

    if idx == n:
        return

    solv(n, cnt, idx+1)
    c[idx] = True
    solv(n, cnt+1, idx+1)
    c[idx] = False

if __name__ == '__main__':
    t = int(stdin.readline())

    for _ in range(t):
        n = int(stdin.readline())

        ax.clear()
        ay.clear()
        c = [False] * 21
        r[0] = float(2**31-1)

        for _ in range(n):
            x, y = map(int, stdin.readline().split())
            ax.append(x)
            ay.append(y)

        solv(n, 0, 0)
        stdout.write('%0.6f\n' % r[0])

