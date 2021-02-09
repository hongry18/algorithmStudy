#!/bin/usr/env python3
# -*- coding: utf-8 -*-

from sys import stdin, stdout

dx = [-1, -1, 0, 1, 1]
dy = [1, 0, 1, 0, 1]
# left top, left, top, right, right top

def sol(g):
    cnt = 0
    s = []

    while s:

        for i in range(5):
            if x == m and y == n:
                break
            nx = x + dx[i]
            ny = y + dy[i]

            if nx < 0 or ny < 0 or nx >= m or ny >= n:
                continue

if __name__ == '__main__':
    c = int(stdin.readline())

    for _ in range(c):
        n, m = map(int, stdin.readline().split())
        row = []
        for y in range(n):
            m_item = stdin.readline()
            col = []
            for x in range(m):
                if m_item[x] == '.':
                    col.append(1)
                else:
                    col.append(0)

            row.append(col)

        print(row)
        sol(row)
