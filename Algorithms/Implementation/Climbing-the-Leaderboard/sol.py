#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the climbingLeaderboard function below.
def climbingLeaderboard(scores, alice):
    ranks = []
    scores = list(dict.fromkeys(scores))

    for s in alice:
        print(scores.index(s))

    return ranks

if __name__ == '__main__':
    scores_count = int(input())

    scores = list(map(int, input().rstrip().split()))

    alice_count = int(input())

    alice = list(map(int, input().rstrip().split()))

    result = climbingLeaderboard(scores, alice)
    print(result)
