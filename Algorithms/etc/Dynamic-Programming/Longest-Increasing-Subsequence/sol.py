#!/usr/bin/env python3
# -*- coding: utf-8 -*-

def sol(arr):
    n = len(arr)

    lis = [1] * n
    for i in range(1, n):
        for j in range(0, i):
            if arr[i] > arr[j] and lis[i] < lis[j] + 1:
                lis[i] = lis[j] + 1


    maximum = 0
    for i in range(n):
        maximum = max(maximum, lis[i])

    print(maximum)

def lis(arr):
    if not arr:
        return 0

    ret = 1
    l = len(arr)
    for i in range(l):
        nxt = []
        for j in range(i+1, l):
            if arr[i] < arr[j]:
                nxt.append(arr[j])

        ret = max(ret, 1 + lis(nxt))

    return ret

i1 = [3, 10, 2, 1, 20]
i2 = [3, 2]
i3 = [50, 3, 10, 7, 40, 80]

sol(i3)

print(lis(i3))
