#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import random

def binary_search(arr, target):
    start = 0
    end = len(arr) - 1

    while start <= end:
        mid = (start + end) // 2

        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            start = mid + 1
        else:
            end = mid - 1

    return None

def binary_search_recursion(t, s, e, a):
    """
        Args
            t: target
            s: start
            e: end
            a: array
        Return
            Numeric
    """

    if s > e:
        return None

    m = (s+e) // 2
    if a[m] == t:
        return m
    elif a[m] > t:
        e = m - 1
    else:
        s = m + 1

    return binary_search_recursion(t, s, e, a)

if __name__ == '__main__':
    l = [i**2 for i in range(11)]
    random.shuffle(l)
    print(l)
    target = 9
    i1 = binary_search(l, target)
    i2 = binary_search_recursion(target, 0, 10, l)

    print(i1)
    print(i2)
