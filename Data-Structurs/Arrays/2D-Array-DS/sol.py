#!/bin/python3
  
import math
import os
import random
import re
import sys

# Complete the hourglassSum function below.
def hourglassSum(arr):
    row_len = len(arr)
    col_len = len(arr[0])
    max_val = -50000

    if row_len < 3:
        return 0

    for ri in range(row_len - 2):
        for ci in range(col_len - 2):
            tmp = sum(arr[ri][ci:ci+3]) + arr[ri+1][ci+1] + sum(arr[ri+2][ci:ci+3])
            print(tmp)
            if max_val < tmp:
                max_val = tmp

    return max_val

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    arr = []

    for _ in range(6):
        arr.append(list(map(int, input().rstrip().split())))

    result = hourglassSum(arr)

    fptr.write(str(result) + '\n')

    fptr.close()
