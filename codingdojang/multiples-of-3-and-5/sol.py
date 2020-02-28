# -*- coding: utf-8 -*-

def sol(x, y, limit):
    s = 0
    if limit % x == 0:
        xc =  (limit - 1) // x
    else:
        xc = limit // x

    for i in range(0, xc):
        s += x + (x * i)

    if limit % y == 0:
        yc =  (limit - 1) // y
    else:
        yc = limit // y

    for i in range(0, yc):
        s += y + (y * i)

    print(s)

if __name__ == '__main__':
    sol(3, 5, 1000)
