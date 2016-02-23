package tree

import (
	"fmt"
)

//Tree Binary tree
type Tree struct {
	value       int
	left, right *Tree
}

//Sort sort int slice in place
func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)
}

//NewTree construct tree with int slice
func NewTree(values []int) *Tree {
	var root *Tree
	for _, v := range values {
		root = add(root, v)
	}
	return root
}

func (t *Tree) String() string {
	var result string

	var stack []*Tree
	var node = t
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.left
		}

		if len(stack) > 0 {
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result += fmt.Sprintf(" %d", node.value)
			node = node.right
		}
	}
	return result
}

func appendValues(values []int, t *Tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *Tree, value int) *Tree {
	if t == nil {
		t = new(Tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func visit(t *Tree) {
	fmt.Println(t.value)
}

//PreOrder pre order traversal tree
func PreOrder(t *Tree) {
	if t != nil {
		visit(t)
		PreOrder(t.left)
		PreOrder(t.right)
	}
}

//InOrder in order traversal tree
func InOrder(t *Tree) {
	if t != nil {
		InOrder(t.left)
		visit(t)
		InOrder(t.right)
	}
}

//PostOrder post order traversal tree
func PostOrder(t *Tree) {
	if t != nil {
		PostOrder(t.left)
		PostOrder(t.right)
		visit(t)
	}
}
