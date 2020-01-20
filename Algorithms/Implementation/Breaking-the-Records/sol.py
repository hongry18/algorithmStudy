#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the breakingRecords function below.
def breakingRecords(scores):
    _cur = scores.pop(0)
    _low = _cur
    _high = _cur
    _low_cnt = 0
    _high_cnt = 0

    for score in scores:
        if score == _cur:
            continue

        if (score > _cur) and (score > _high):
            _high = score
            _high_cnt += 1

        if (score < _cur) and (score < _low):
            _low = score
            _low_cnt += 1

        _cur = score

    return (_high_cnt, _low_cnt)

if __name__ == '__main__':
    n = int(input())

    scores = list(map(int, input().rstrip().split()))

    result = breakingRecords(scores)

    print(' '.join(map(str, result)))
