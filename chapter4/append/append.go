package main

import "fmt"

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse2(s *[10]string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotate(s []string, n int) {
	if n >= len(s) {
		return
	}
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func rotate2(s []string, n int) {
}

func adjacentRemove(s []string) []string {
	fmt.Println(len(s))
	j := len(s) - 1
	for i := 0; i < j; i++ {
		if s[i] == s[i + 1] {
			copy(s[i:], s[i+1:])
			i--
			j--
			fmt.Println(j)
		}
	}
	return s[:j+1]
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func appendInt(x []string, y ...string) []string {
	var z []string
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]string, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []string, i int) []string {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	s := []string{"first", "second", "first", "second", "second", "second", "fourth", "fourth", "fifth", "fifth"}
	fmt.Println(adjacentRemove(s))
}
