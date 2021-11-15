package main

import (
	"crypto/sha256"
	"fmt"
	"tgopl/chapter2/popcount"
)

func bitdiff(sha1, sha2 [32]byte) int {
	n := 0
	for i := range sha1 {
		n += popcount.PopCount(uint64(sha1[i] & sha2[i]))
	}
	return n
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(bitdiff(c1, c2))
}
