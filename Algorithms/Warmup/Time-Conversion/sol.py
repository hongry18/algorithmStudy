#!/usr/bin/env python3

import os
import sys

def timeConversion(s):
    div = s[-2:]
    times = s[:-2].split(':')
    h = int(times[0])

    if div == 'AM' and h > 11 :
        h = h - 12
    elif div == 'PM' and h < 12:
        h = h + 12

    if h == 24:
        h = 0

    times[0] = '{0:02d}'.format(h)
    return ':'.join(times)


if __name__ == '__main__':
    s = input()

    r = timeConversion(s)
    print(r)
