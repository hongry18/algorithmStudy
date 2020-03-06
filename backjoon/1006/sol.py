#!/usr/bin/env python3
# -*- conding: utf-8 -*-

from sys import stdin, stdout

def solv(s, a, b, c, e, w, n):
    for i in range(s, n):
        a[i+1] = min(b[i] + 1, c[i] + 1)

        if e[i][0] + e[i][1] <= w:
            a[i+1] = min(a[i+1], a[i] + 1)

        if i>0 and e[i-1][0] + e[i][0] <= w and e[i-1][1] +e[i][1] <= w:
            a[i+1] = min(a[i+1], a[i-1] + 2)

        if (i < n-1):
            b[i+1] = a[i+1] +1
            if (e[i][0] + e[i+1][0]) <= w:
                b[i+1] = min(b[i+1], c[i] + 1)

            c[i+1] = a[i+1] + 1
            if (e[i][1] + e[i+1][1]) <= w:
                c[i+1] = min(c[i+1], b[i] + 1)

if __name__ == '__main__':
    t = int(stdin.readline())

    for _ in range(t):
        n, w = map(int, stdin.readline().split())
        res = 30000

        l1 = list(map(int, stdin.readline().split()))
        l2 = list(map(int, stdin.readline().split()))

        e = [ [l1[i], l2[i]] for i in range(n) ]
        a = [0] * (n+1)
        b = [0] * (n+1)
        c = [0] * (n+1)
        b[0] = 1
        c[0] = 1

        solv(0, a, b, c, e, w, n)
        res = min(res,a[n])

        if n > 1:
            if e[0][0] + e[n-1][0] <= w:
                a[1] = 1
                b[1] = 2
                c[1] = 1 if e[0][1] + e[1][1] <= w else 2
                solv(1, a, b, c, e, w, n)
                res = min(res,c[n-1]+1)

            if e[0][1] + e[n-1][1] <= w:
                a[1] = 1
                b[1] = 1 if e[0][0] + e[1][0] <= w else 2
                c[1] = 2
                solv(1, a, b, c, e, w, n)
                res = min(res,b[n-1]+1)

            if e[0][0] + e[n-1][0] <= w and e[0][1] + e[n-1][1] <= w:
                a[1] = 0
                b[1] = 1
                c[1] = 1
                solv(1, a, b, c, e, w, n)
                res = min(res,a[n-1]+2)

        print(res)
