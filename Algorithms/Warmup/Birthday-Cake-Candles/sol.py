#!/usr/bin/env python3

def birthdayCakeCandles(ar):
    _m = 0
    _c = 0
    for i in ar:
        if _m < i:
            _m = i
            _c = 0

        if i == _m:
            _c += 1

    return _c

if __name__ == '__main__':
    ar_cnt = int(input())

    ar = list(map(int, input().rstrip().split()))

    r = birthdayCakeCandles(ar)
    print(r)
