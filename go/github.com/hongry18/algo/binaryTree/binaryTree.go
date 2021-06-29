package binaryTree

import (
    "io"
    "fmt"
)

type Node struct {
    data int64
    left *Node
    right *Node
}

type BinaryTree struct {
    root *Node
}

func (t *BinaryTree) GetRoot() *Node {
    return t.root
}

func (t *BinaryTree) GetLeft() *Node {
    return t.root.left
}

func (t *BinaryTree) GetRight() *Node {
    return t.root.right
}

func (t *BinaryTree) Insert(data int64) *BinaryTree {
    if t.root == nil {
        t.root = &Node{data: data}
    } else {
        t.root.Insert(data)
    }
    return t
}

func (n *Node) Insert(data int64) {
    if n == nil {
        return
    }

    if data <= n.data {
        if n.left == nil {
            n.left = &Node{data: data}
        } else {
            n.left.Insert(data)
        }

        return
    }

    if n.right == nil {
        n.right = &Node{data: data}
    } else {
        n.right.Insert(data)
    }
}

func Print(w io.Writer, node *Node, ns int, ch rune) {
    if node == nil {
        return
    }


    for i:=0; i<ns; i++ {
        fmt.Fprint(w, " ")
    }

    fmt.Fprintf(w, "%c:%+v\n", ch, node.data)
    Print(w, node.left, ns+2, 'L')
    Print(w, node.right, ns+2, 'R')
}

