#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from sys import stdin, stdout

def sol(n, m):
    r = 1
    k = m - n
    while m > k:
        r *= m
        m -= 1

    while n > 1:
        r = r // n
        n -= 1

    stdout.write('%d\n' % r)

if __name__ == '__main__':
    t = int(stdin.readline())

    for _ in range(t):
        n, m = map(int, stdin.readline().split())
        sol(n, m)
