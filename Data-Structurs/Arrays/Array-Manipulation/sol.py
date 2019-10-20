#!/bin/python

import math
import os
import random
import re
import sys

# Complete the arrayManipulation function below.
def arrayManipulation(n, queries):
    res = [0] * n
    max_val = -1
    for q in queries:
        res[q[0]-1] += q[2]
        if q[1] < n:
            res[q[1]] += -1*q[2]
    
    t = 0
    for r in res:
        t += r
        if max_val < t:
            max_val = t

    return max_val

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    nm = raw_input().split()

    n = int(nm[0])

    m = int(nm[1])

    queries = []

    for _ in xrange(m):
        queries.append(map(int, raw_input().rstrip().split()))

    result = arrayManipulation(n, queries)

    fptr.write(str(result) + '\n')

    fptr.close()

