package main

import (
	"fmt"
)

func main() {
    list := IntList {
        1, &IntList {
            2, &IntList{
                3, nil,
            },
        },
    }
    
    sum := list.Sum()
    fmt.Println(sum)
}