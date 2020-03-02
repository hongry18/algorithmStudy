#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from sys import stdin, stdout
from queue import Queue

T = int(input())

def topology(g, D, indegree, W):
    queue = Queue()
    tmp = dict()

    for i in g:
        if indegree[i] == 0:
            tmp[i] = D[i]
            queue.put(i)
        else:
            tmp[i] = 0

    for t in g:
        node = queue.get()
        if node == W:
            return tmp[node]

        for i in g[node]:
            indegree[i] -= 1
            tmp[i] = max(tmp[i], tmp[node]+D[i])
            if indegree[i] == 0:
                queue.put(i)

for _ in range(T):
    N, K = map(int, stdin.readline().split())
    D = (0,) + tuple(map(int, stdin.readline().split()))

    g = dict()
    indegree = dict()

    for i in range(1, N+1):
        g[i] = []
        indegree[i] = 0

    for i in range(K):
        X, Y = map(int, stdin.readline().split())
        g[X].append(Y)
        indegree[Y] += 1

    W = int(stdin.readline())

    print(topology(g, D, indegree, W))
