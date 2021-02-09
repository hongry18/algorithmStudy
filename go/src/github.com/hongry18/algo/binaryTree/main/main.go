package main

import (
    "os"
    //"fmt"
    "math/rand"
    "github.com/hongry18/algo/binaryTree"
)

func main() {
    tree := &binaryTree.BinaryTree{}

    for i:=0; i<80; i++ {
        tree.Insert(rand.Int63n(200))
    }

    binaryTree.Print(os.Stdout, tree.GetRoot(), 0, 'M')
}
