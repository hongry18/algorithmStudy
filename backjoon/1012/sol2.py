#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from sys import stdin, stdout, setrecursionlimit
setrecursionlimit(10**6)

dx = [-1, 1, 0, 0]
dy = [0, 0, 1, -1]

def dfs(x, y):
    l[y][x] = 0

    for i in range(4):
        nx = x + dx[i]
        ny = y + dy[i]

        if nx < 0 or nx >= m or ny < 0 or ny >= n:
            # array bound check
            continue

        if l[ny][nx] == 1:
            dfs(nx, ny)

if __name__ == '__main__':
    t = int(stdin.readline())

    for _ in range(t):
        m, n, k = map(int, stdin.readline().split())
        l = [ [0] * m for _ in range(n) ]
        cnt = 0

        for _ in range(k):
            x, y = map(int, stdin.readline().split())
            l[y][x] = 1

        for y in range(n):
            for x in range(m):
                if l[y][x] == 1:
                    dfs(x, y)
                    cnt += 1

        stdout.write('%d\n' % (cnt))
