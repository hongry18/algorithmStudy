#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the miniMaxSum function below.
def miniMaxSum(arr):
    _min = 1e+10
    _max = 0
    for i in range(len(arr)):
        s = 0
        for j in range(len(arr)):
            if j == i:
                continue
            s += arr[j]

        _min = _min if s > _min else s
        _max = _max if s < _max else s

    print(_min, _max)

if __name__ == '__main__':
    #arr = list(map(int, input().rstrip().split()))
    arr = [i for i in range(1,6)]

    miniMaxSum(arr)

