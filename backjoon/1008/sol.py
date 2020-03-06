#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from sys import stdin, stdout

if __name__ == '__main__':
    x, y = map(int, stdin.readline().split())

    stdout.write( ('%.9f' % (x / y)).rstrip('0').rstrip('.'))
