#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from queue import Queue

graph = {
    'A': ['B'],
    'B': ['A', 'C', 'H'],
    'C': ['B', 'D', 'G'],
    'D': ['C', 'E'],
    'E': ['D', 'F'],
    'F': ['E'],
    'G': ['D'],
    'H': ['B', 'I', 'J', 'M'],
    'I': ['H'],
    'J': ['H', 'K'],
    'K': ['J', 'L'],
    'L': ['K'],
    'M': ['H']
}

def bfs(g, s):
    q = Queue()
    v = set()
    res = []

    q.put(s)

    while not q.empty():
        n = q.get()
        if n not in v:
            v.add(n)
            res.append(n)
            for next_n in g[n]:
                q.put(next_n)

    return res


if __name__ == '__main__':
    print(bfs(graph, 'A'))
