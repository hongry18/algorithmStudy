#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import random

def heapify(unsorted, i, heap_size):
    print('unsorted', unsorted)
    largest = i
    left = 2 * i + 1
    right = 2 * i + 2

    if left < heap_size and unsorted[left] > unsorted[largest]:
        largest = left

    if right < heap_size and unsorted[right] > unsorted[largest]:
        largest = right

    # checked largest
    if largest != i:
        unsorted[largest], unsorted[i] = unsorted[i], unsorted[largest]
        heapify(unsorted, largest, heap_size)

def heap_sort(unsorted):
    print('before max heap : ', unsorted)
    n = len(unsorted)
    for i in range(n // 2 -1, -1, -1):
        heapify(unsorted, i, n)

    print('*' * 100)
    print('after max heap', unsorted)
    print('*' * 100)

    for i in range(n -1, 0, -1):
        unsorted[0], unsorted[i] = unsorted[i], unsorted[0]
        heapify(unsorted, 0, i)

    return unsorted

l = [random.randint(1, 20) for i in range(20)]

print(l)
print(heap_sort(l))
