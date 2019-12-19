#!/usr/bin/env python3

import random

class Node:

    def __init__(self, data):
        self.left = None
        self.right = None
        self.data = data

class BinaryTree():
    def __init__(self):
        self.root = None

    def insert(self, data):
        if self.root is None:
            self.root = Node(data)
        else:
            cur = self.root

            while True:
                if data < cur.data:
                    if cur.left:
                        cur = cur.left
                    else:
                        cur.left = Node(data)
                        break
                elif data > cur.data:
                    if cur.right:
                        cur = cur.right
                    else:
                        cur.right = Node(data)
                        break
                else:
                    break

# end BinaryTree class

def preorder(root):
    if root is None:
        return
    cur = root

    print(cur.data)
    preorder(cur.left)
    preorder(cur.right)

def inorder(root):
    if root is None:
        return
    cur = root

    inorder(cur.left)
    print(cur.data)
    inorder(cur.right)

def postorder(root):
    if root is None:
        return
    cur = root

    postorder(cur.left)
    postorder(cur.right)
    print(cur.data)

def bf_print(root):
    # breadth first
    if root is None:
        return
    cur = root
    q = [root]

    while q:
        n = q.pop(0)
        print(n.data)
        if n.left:
            q.append(n.left)
        if n.right:
            q.append(n.right)

def df_print(root):
    # depth first
    if root is None:
        return
    cur = root
    s = [root]

    while s:
        n = s.pop()
        print(n.data)
        if n.right:
            s.append(n.right)
        if n.left:
            s.append(n.left)


items = [i for i in range(1, 50)]
random.shuffle(items)

items = [i for i in range(1,16)]
random.shuffle(items)

tree = BinaryTree()
for i in items:
    tree.insert(i)

print(items)

print('preorder')
preorder(tree.root)

print('inorder')
inorder(tree.root)

print('postorder')
postorder(tree.root)

print('depth first')
df_print(tree.root)

print('breadth first')
bf_print(tree.root)
