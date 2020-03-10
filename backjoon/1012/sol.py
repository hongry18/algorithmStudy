#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from sys import stdin, stdout, setrecursionlimit
setrecursionlimit(10**6)

dx = [0, 1, 0, -1]
dy = [1, 0, -1, 0]
d = []

def dfs(x, y):
    d[x][y] = 0
    for i in range(4):
        nx = x + dx[i]
        ny = y + dy[i]
        if nx < 0 or nx >= m or ny < 0 or ny >= n:
            continue

        if d[nx][ny] == 1:
            dfs(nx, ny)

if __name__ == '__main__':
    t = int(stdin.readline())

    for _ in range(t):
        m, n, k = map(int, stdin.readline().split())
        d.clear()
        d = [ [0]*n for _ in range(m)]

        for _ in range(k):
            x, y = map(int, stdin.readline().split())
            d[x][y] = 1

        cnt = 0
        for x in range(m):
            for y in range(n):
                if d[x][y] == 1:
                    dfs(x, y)
                    cnt += 1

        stdout.write('%d\n' % (cnt))
