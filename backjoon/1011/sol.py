#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from sys import stdin, stdout

def sol(a):
    ck = set([1,2,3])
    if a in ck:
        stdout.write('%d\n' % a )
        return

    n = 2
    minN = 0
    maxN = 0
    cnt = 0

    while True:
        powN = n * n
        minN = powN - n + 1
        maxN = powN + n

        if minN <= a and maxN >= a:
            cnt = 2 * n
            if a <= powN:
                cnt -= 1

            break

        n += 1

    stdout.write('%d\n' % cnt )

if __name__ == '__main__':
    t = int(stdin.readline())

    for _ in range(t):
        x, y = map(int, stdin.readline().split())
        dist = y - x
        sol(dist)
