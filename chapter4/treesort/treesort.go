package tree

import "strconv"

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}
	s := ""
	sep := " "
	if t.left != nil {
		s += t.left.String()
	}
	s += strconv.Itoa(t.value) + sep // added sep at the end for simplicity
	if t.right != nil {
		s += t.right.String()
	}
	return s
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		return &tree{value: value}
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
