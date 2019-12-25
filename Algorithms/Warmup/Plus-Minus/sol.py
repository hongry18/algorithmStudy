#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the plusMinus function below.
def plusMinus(arr):
    r = [0, 0, 0]
    size = len(arr)

    for i in arr:
        if i > 0:
            r[0] += 1
        elif i < 0:
            r[1] += 1
        else:
            r[2] += 1

    for i in r:
        print('{:.6f}'.format(i / size))

if __name__ == '__main__':
    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    plusMinus(arr)
