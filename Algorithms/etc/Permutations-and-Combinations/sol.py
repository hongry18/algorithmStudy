#!/usr/bin/env python3
# -*- coding: utf-8 -*-

def permutation(a, r):
    """
    Args:
        a - target array
        r - count
    """

    a = sorted(a)
    used = [0 for _ in range(len(a))]

    def generate(c):
        """
        Args:
            c - chosen
            u - used
        """

        if len(c) == r:
            print(c)
            return

        for i in range(len(a)):
            if not used[i] and (i == 0 or a[i-1] != a[i] or used[i-1]):
                c.append(a[i])
                used[i] = 1
                generate(c)
                c.pop()
                used[i] = 0

    generate([])

def combination(a, r):
    """
    Args:
        a - target array
        r - count
    """

    a = sorted(a)
    used = [0 for _ in range(len(a))]

    def generate(c):
        """
        Args:
            c - chosen
        """
        if len(c) == r:
            print(c)
            return

        start = a.index(c[-1]) + 1 if c else 0
        for n in range(start, len(a)):
            if used[n] == 0 and (n == 0 or a[n-1] != a[n] or used[n-1]):
                c.append(a[n])
                used[n] = 1
                generate(c)
                c.pop()
                used[n] = 0

    generate([])

if __name__ == '__main__':
    permutation('ABCD', 2)
    permutation([1, 2, 3, 4, 5], 3)

    combination('ABCDE', 2)
    combination([1, 2, 3, 4, 5], 3)

    permutation('AABC', 2)
