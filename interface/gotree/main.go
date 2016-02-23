package main

import (
    "github.com/jiayanju/gobible/interface/gotree/tree"
	"fmt"
)

func main() {
    values := []int {1, 9, 4, 2, 18, 10, 7}
    
    root := tree.NewTree(values)
    
    fmt.Println("In Order")
    tree.InOrder(root)
    
    fmt.Println("Pre Order")
    tree.PreOrder(root)
    
    fmt.Println("Post Order")
    tree.PostOrder(root)
    
        
    tree.Sort(values)
    
    fmt.Println(values)
}