#!/bin/python3

import math
import os
import random
import re
import sys

class SinglyLinkedListNode:
    def __init__(self, node_data):
        self.data = node_data
        self.next = None

class SinglyLinkedList:
    def __init__(self):
        self.head = None
        self.tail = None

    def insert_node(self, node_data):
        node = SinglyLinkedListNode(node_data)

        if not self.head:
            self.head = node
        else:
            self.tail.next = node


        self.tail = node

def print_singly_linked_list(node, sep, fptr):
    while node:
        fptr.write(str(node.data))

        node = node.next

        if node:
            fptr.write(sep)

# Complete the reverse function below.

#
# For your reference:
#
# SinglyLinkedListNode:
#     int data
#     SinglyLinkedListNode next
#
#
# O(N)
def reverse(head):
    prev_node = SinglyLinkedListNode(head.data)
    head = head.next
    while head:
        cur_node = SinglyLinkedListNode(head.data)
        cur_node.next=prev_node
        prev_node = cur_node
        head = head.next

    return cur_node

# O(2N)
"""
def reverse(head):
    r_node = SinglyLinkedList()
    r_arr = []

    while head:
        r_arr.append(head.data)
        head = head.next

    for i in range(len(r_arr)):
        r_node.insert_node(r_arr.pop())

    return r_node.head

"""

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    tests = int(input())

    for tests_itr in range(tests):
        llist_count = int(input())

        llist = SinglyLinkedList()

        for _ in range(llist_count):
            llist_item = int(input())
            llist.insert_node(llist_item)

        llist1 = reverse(llist.head)

        print_singly_linked_list(llist1, ' ', fptr)
        fptr.write('\n')

    fptr.close()

