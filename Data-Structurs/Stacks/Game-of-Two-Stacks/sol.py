#!/bin/python3

import os
import sys

#
# Complete the twoStacks function below.
#
def twoStacks(x, a, b):
    #
    # Write your code here.
    #

    print (x, a, b)
    a_id = len(a) - 1
    b_id = len(b) - 1
    t = 0
    r = 0

    if a[a_id] > x:
        r = a[a_id]
        return r

    if b[b_id] > x:
        r = b[b_id]
        return r

    while True:
        bigger = a[a_id] > b[b_id]

    return 1

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    g = int(input())

    for g_itr in range(g):
        nmx = input().split()

        n = int(nmx[0])

        m = int(nmx[1])

        x = int(nmx[2])

        a = list(map(int, input().rstrip().split()))

        b = list(map(int, input().rstrip().split()))

        result = twoStacks(x, a, b)

        fptr.write(str(result) + '\n')

    fptr.close()
