#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from sys import stdin, stdout
from queue import Queue

dx = [1, -1, 0, 0]
dy = [0, 0, 1, -1]

def bfs(i, j):
    q.put([i, j])
    # set start point
    ck[j][i] = True
    dist[j][i] = 1

    while not q.empty():
        x, y = q.get()

        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]

            if nx < 0 or nx >= m or ny < 0 or ny >= n:
                continue
            if ck[ny][nx] == 1:
                continue
            if _map[ny][nx] == 0:
                continue

            q.put([nx, ny])
            dist[ny][nx] = dist[y][x] + 1
            ck[ny][nx] = True

    return dist[n-1][m-1]


if __name__ == '__main__':
    n, m = map(int, stdin.readline().split())

    _map = []
    q = Queue()
    ck = [[False for _ in range(m)] for _ in range(n)]
    dist = [[0 for _ in range(m)] for _ in range(n)]

    for _ in range(n):
        y = stdin.readline()
        row = []
        for i in range(len(y)-1):
            row.append(int(y[i]))

        _map.append(row)

    print(bfs(0, 0))
