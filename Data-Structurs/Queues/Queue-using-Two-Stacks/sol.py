#!/usr/bin/env python3

import sys

q = int(sys.stdin.readline().split()[0])
qe = 1

class TwoStacksQueue:
    def __init__(self):
        self.s1 = []
        self.s2 = []

    def enqueue(self, x):
        self.s1.append(x)

    def dequeue(self):
        if not self.s1 and not self.s2:
            print('Q is empty')
            sys.exit()


        if not self.s2:
            while self.s1:
                self.s2.append(self.s1.pop())

        return self.s2.pop()

    def peek(self):
        if not self.s2:
            while self.s1:
                self.s2.append(self.s1.pop())

        return self.s2[-1]

queue = TwoStacksQueue()

while True:
    line = sys.stdin.readline().split()
    x = line[0]


    if x == '1':
        # enqueue element x into the end of the queue
        queue.enqueue(line[1])
    elif x == '2':
        # dequeue the element at the front of the queue
        queue.dequeue()
    elif x == '3':
        # print the elemment at the front of the queue
        print(queue.peek())
        pass

    if qe == q:
        break
    qe+=1
