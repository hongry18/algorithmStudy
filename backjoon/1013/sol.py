#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from sys import stdin, stdout
import re

state_map = [
    [1, 2]
    ,[-1, 5]
    ,[3, -1]
    ,[4, -1]
    ,[4, 6]
    ,[1, 2]
    ,[1, 7]
    ,[8, 7]
    ,[4, 0]
]

matched = set([5, 6, 7])

def is_matched(s):
    _len = len(s)
    if _len < 2:
        stdout.write('NO\n')
        return

    state = 0
    for _i in range(_len):
        ch = 0 if s[_i] == '0' else 1
        state = state_map[state][ch]
        if state == -1:
            break

    if state in matched:
        stdout.write('YES\n')
    else:
        stdout.write('NO\n')

def is_matched2(s):
    if len(s) < 2:
        stdout.write('NO\n')
        return

    m = re.fullmatch('(100+1+|01)+', s)
    if m:
        stdout.write('YES\n')
    else:
        stdout.write('NO\n')

if __name__ == '__main__':
    n = int(stdin.readline())

    for _ in range(n):
        s = stdin.readline().rstrip()
        is_matched2(s)
