#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys

def sol(a):
    pass

x = int(sys.argv[1])

dp = [0 for _ in range(x+1)]

div = [2,3]

for i in range(2, x+1):
    dp[i] = dp[i-1] + 1
    for di in range(2):
        if not i % div[di] and dp[i] > dp[i // div[di]] + 1:
            dp[i] = dp[i // div[di]] + 1

"""
for i in range(2, x+1):
    dp[i] = dp[i-1] + 1
    if not i % 2 and dp[i] > dp[i//2] + 1:
        dp[i] = dp[i//2] + 1
    if not i % 3 and dp[i] > dp[i//3] + 1:
        dp[i] = dp[i//3] + 1
"""

print(dp[x])
print(dp)
