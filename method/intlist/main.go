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
    
    list1 := IntList{}
    list1.Value = 1
    list1.Tail = &IntList{}
    
    list1.Tail.Value = 2
    
    item3 := IntList{}
    item3.Value = 3;
    item3.Tail = nil;
    
    list1.Tail.Tail = &item3
    
    sum = list1.Sum()
    fmt.Println(sum)
}