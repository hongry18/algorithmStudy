#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from sys import stdin, stdout

t = int(stdin.readline())
pass_list = set([1,5,6])

for _ in range(t):
    a, b = map(int, stdin.readline().split(' '))

    if a in pass_list:
        # 1, 5, 6은 아무리 제곱해봐도 같은 숫자가 나오므로 by pass
        stdout.write('%d\n' % a)
        continue

    tmp = 1
    dup_list = []
    for _ in range(b):
        tmp = (tmp * a) % 10
        if tmp in dup_list:
            break
        dup_list.append(tmp)

    r = dup_list[ b % len(dup_list) - 1]
    if r == 0:
        r = 10

    stdout.write('%d\n' % r)
