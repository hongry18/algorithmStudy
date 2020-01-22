#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the sockMerchant function below.
def sockMerchant(n, ar):
    dup = list()
    dup_cnt = 0

    for i in ar:
        if i in dup:
            dup_cnt += 1
            dup.remove(i)
        else:
            dup.append(i)

    return dup_cnt

if __name__ == '__main__':
    n = int(input())

    ar = list(map(int, input().rstrip().split()))

    result = sockMerchant(n, ar)
