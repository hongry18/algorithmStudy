#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys

def sol(a):
    f = [0, 1]

    for i in range(2, a+1):
        f.append(f[i-1] + f[i-2])

    print(f[a])

x = int(sys.argv[1])
sol(x)
