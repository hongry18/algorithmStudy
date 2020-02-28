#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# BFS Algorithm

import queue

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

def bfs(g, start):
    visit = []
    q = []

    q.append(start)

    while q:
        node = q.pop(0)
        if node not in visit:
            visit.append(node)
            q.extend(g[node])

    return visit

def q_bfs(g, start):
    visit = set()
    res = []
    q = queue.Queue()

    q.put(start)

    while q.qsize() > 0:
        node = q.get()
        if node not in visit:
            visit.add(node)
            res.append(node)
            for next_node in g[node]:
                q.put(next_node)

    return res

if __name__ == '__main__':
    p = bfs(graph, 'A')
    print(p)
    p = q_bfs(graph, 'A')
    print(p)
