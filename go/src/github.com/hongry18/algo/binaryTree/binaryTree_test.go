package binaryTree

import "testing"

func TestNode(t *testing.T) {
    tree := &BinaryTree{}
    tree.Insert(100)

    t.Errorf("%+v", tree.root)
}
