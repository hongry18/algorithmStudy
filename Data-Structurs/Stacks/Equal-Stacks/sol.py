#!/bin/python3

import os
import sys

#
# Complete the equalStacks function below.
#

def equalStacks2(h1, h2, h3):
    #
    # Write your code here.
    #
    r1 = []
    r2 = []
    r3 = []

    s = [0, 0, 0]
    r = 0

    for _ in range(len(h1)):
        d = h1.pop()
        r1.append(d)
        s[0] += d

    for _ in range(len(h2)):
        d = h2.pop()
        r2.append(d)
        s[1] += d

    for _ in range(len(h3)):
        d = h3.pop()
        r3.append(d)
        s[2] += d

    while True:
        if s[0] == s[1] and s[1] == s[2]:
            r = s[0]
            break

        if not r1 or not r2 or not r3:
            break

        max_id = s.index( max(s) )

        if max_id == 0:
            m = r1.pop()
        elif max_id == 1:
            m = r2.pop()
        else:
            m = r3.pop()

        s[max_id] -= m


    return r

def equalStacks(h1, h2, h3):
    #
    # Write your code here.
    #
    r1 = [h1.pop() for _ in range(len(h1))]
    r2 = [h2.pop() for _ in range(len(h2))]
    r3 = [h3.pop() for _ in range(len(h3))]

    s = [sum(r1), sum(r2), sum(r3)]
    r = 0

    while True:
        if s[0] == s[1] and s[1] == s[2]:
            r = s[0]
            break

        if not r1 or not r2 or not r3:
            break

        max_id = s.index( max(s) )

        if max_id == 0:
            m = r1.pop()
        elif max_id == 1:
            m = r2.pop()
        else:
            m = r3.pop()

        s[max_id] -= m

    return r

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n1N2N3 = input().split()

    n1 = int(n1N2N3[0])

    n2 = int(n1N2N3[1])

    n3 = int(n1N2N3[2])

    h1 = list(map(int, input().rstrip().split()))

    h2 = list(map(int, input().rstrip().split()))

    h3 = list(map(int, input().rstrip().split()))

    result = equalStacks(h1, h2, h3)

    fptr.write(str(result) + '\n')

    fptr.close()

