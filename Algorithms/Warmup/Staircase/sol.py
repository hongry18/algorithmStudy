#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the staircase function below.
def staircase(n):
    space_cnt = n-1
    for i in range(n):
        p = []
        ncnt = space_cnt - i
        scnt = n - ncnt
        for _i in range(n):
            if _i < ncnt:
                p.append(' ')
            else:
                p.append('#')
        print(''.join(p))

if __name__ == '__main__':
    n = int(input())

    staircase(n)
