#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the countingValleys function below.
def countingValleys(n, s):
    ch = 0
    ph = 0
    visit = 0
    for c in s:
        ph = ch
        if c == 'U':
            ch += 1
        elif c == 'D':
            ch -= 1

        if ph == 0 and ch == -1:
            visit += 1

    return visit

if __name__ == '__main__':

    n = int(input())

    s = input()

    result = countingValleys(n, s)
    print(result)
