package tree

import (
	"fmt"
	"testing"
)

// not a test even
func TestString(t *testing.T) {
	tree := add(nil, 4)
	add(tree, 6)
	add(tree, 7)
	add(tree, 7)
	fmt.Println(tree)
}
