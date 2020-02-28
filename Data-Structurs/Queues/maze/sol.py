#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from queue import Queue

def maze(g, s, e):
    """
    Args:
        g : graph
        s : start node
        e : end node

    Returns:
        String
    """

    q = Queue()
    done = set()

    q.put(s)
    done.add(s)

    while q.qsize() > 0:
        node = q.get()
        v = node[-1]
        if v == e:
            return node

        for x in g[v]:
            if x not in done:
                q.put(node + x)
                done.add(x)

    return '?'

graph = {
    'a' : ['e'],
    'b' : ['c', 'f'],
    'c' : ['b', 'd'],
    'd' : ['c'],
    'e' : ['a', 'i'],
    'f' : ['b', 'g', 'j'],
    'g' : ['f', 'h'],
    'h' : ['g', 'l'],
    'i' : ['e', 'm'],
    'j' : ['f', 'k', 'n'],
    'k' : ['j', 'o'],
    'l' : ['h', 'p'],
    'm' : ['i', 'n'],
    'n' : ['m', 'j'],
    'o' : ['k'],
    'p' : ['l']
}
print(maze(graph, 'a', 'p'))
