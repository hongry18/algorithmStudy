#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the compareTriplets function below.
def compareTriplets(a, b):
    res = [0, 0]
    a_len = len(a)
    b_len = len(b)
    l_len = a_len if a_len > b_len else b_len

    for i in range(l_len):
        if a[i] > b[i]:
            res[0] += 1
        elif b[i] > a[i]:
            res[1] += 1
        else:
            continue

    return res

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    a = list(map(int, input().rstrip().split()))

    b = list(map(int, input().rstrip().split()))

    result = compareTriplets(a, b)

    fptr.write(' '.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
