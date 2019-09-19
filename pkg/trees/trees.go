package trees

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New returns a new, random binary tree holding the values k, 2k, ..., 10k.
func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

/*
Walk walks the tree t (in-order visit) sending all values from the tree
to the channel ch.
*/
func Walk(t *Tree, ch chan int) {
	inOrderVisit(t, func(t *Tree) { ch <- t.Value })
	// closing the channel signals the receiver that there is no more data to be sent on it
	close(ch)
}

type visit func(*Tree)

func inOrderVisit(t *Tree, v visit) {
	if t != nil {
		inOrderVisit(t.Left, v)
		v(t)
		inOrderVisit(t.Right, v)
	}
}

/*
Same determines whether the trees t1 and t2 contain the same values.
*/
func Same(t1, t2 *Tree) bool {
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var v1 []int
	var v2 []int
	for v := range ch1 {
		v1 = append(v1, v)
	}
	for v := range ch2 {
		v2 = append(v2, v)
	}

	return intArrayEquals(v1, v2)
}

func intArrayEquals(a1 []int, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i, v := range a1 {
		if v != a2[i] {
			return false
		}
	}
	return true
}
