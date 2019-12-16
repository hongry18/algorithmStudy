#!/usr/bin/env python3
import unittest

class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

class SingleLinkedList:
    def __init__(self):
        self.head = None

    def print(self, head):
        cur = head
        while cur:
            print(cur.data)
            cur = cur.next

    def append(self, d):
        if self.head is None:
            self.head = Node(d)
            return

        cur = self.head
        while cur.next is not None:
            cur = cur.next

        cur.next = Node(d)

    def add2head(self, d):
        if self.head is None:
            self.head = Node(d)
        else:
            t = self.head
            self.head = Node(d)
            self.head.next = t

    def remove_first(self):
        if self.head is None:
            return

        self.head = self.head.next

    def remove_duplicate(self):
        cur = self.head
        dup = []

        while(cur.next):
            if cur.data in dup:
                cur.data = cur.next.data
                cur.next = cur.next.next
                continue

            dup.append(cur.data)
            cur = cur.next

    def reverse(self):
        cur = self.head
        r = Node(cur.data)
        cur = cur.next

        while(cur):
            t = r
            r = Node(cur.data)
            r.next = t
                
            cur = cur.next

        self.print(r)


class SingleLinkedListTest(unittest.TestCase):

    sll = None

    def setUp(self):
        self.sll = SingleLinkedList()
        self.sll.append(5)
        self.sll.add2head(10)
        self.sll.append(4)
        self.sll.add2head(20)
        self.sll.append(1)
        self.sll.append(2)
        self.sll.append(3)
        self.sll.append(1)
        self.sll.add2head(40)
        self.sll.append(2)
        self.sll.append(2)
        self.sll.append(9)
        self.sll.add2head(30)
        self.sll.add2head(50)
        print('# start {}'.format(self._testMethodName))

    def tearDown(self):
        print('# end {}'.format(self._testMethodName))

    def test_add2head(self):
        self.sll.add2head(70)
        pass

    def test_remove_first(self):
        self.sll.remove_first()

    def test_remove_duplicate(self):
        self.sll.remove_duplicate()

    def test_reverse(self):
        self.sll.reverse()

    def test_print(self):
        self.sll.print(self.sll.head)

if __name__ == '__main__':
    unittest.main()
