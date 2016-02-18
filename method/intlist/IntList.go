package main

//IntList int list
type IntList struct {
    Value int
    Tail *IntList
}

//Sum method
func (list *IntList) Sum() int  {
    if list == nil {
        return 0
    }
    
    return list.Value + list.Tail.Sum()
}