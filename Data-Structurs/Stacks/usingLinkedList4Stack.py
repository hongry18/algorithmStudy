#!/usr/bin/env python3

class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

class Stack:
    def __init__(self):
        self.head = None

    def push(self, data):
        node = Node(data)
        node.next = self.head
        self.head = node

    def pop(self):
        if self.is_empty():
            return

        cur = self.head
        self.head = self.head.next
        return cur.data

    def peak(self):
        if self.is_empty():
            return -1

        return self.head.data

    def is_empty(self):
        if self.head:
            return False
        return True

    def print_all(self):
        r = []
        cur = self.head
        while cur:
            r.append(cur.data)
            cur = cur.next

        print(r)

s = Stack()

s.push(1)
s.print_all()
s.push(3)
s.print_all()
s.push(2)
s.print_all()

for i in range(12, 83):
    s.push(i)

s.pop()
s.print_all()
