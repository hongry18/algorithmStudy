#!/usr/bin/env python3
# -*- coding: utf-8 -*-

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

def dfs(g, start):
    visit = []
    stack = []

    stack.append(start)

    while stack:
        node = stack.pop()
        if node not in visit:
            visit.append(node)
            stack.extend(g[node])

    return visit

if __name__ == '__main__':
    p = dfs(graph, 'A')
    print(p)
