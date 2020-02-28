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

def dfs(g, s):
    visit = set()
    stack = list()
    res = list()

    stack.append(s)

    while stack:
        node = stack.pop()
        if node not in visit:
            visit.add(node)
            res.append(node)
            for next_n in g[node]:
                stack.append(next_n)

    return res

if __name__ == '__main__':
    print(dfs(graph, 'A'))
