package main

import (
	"bytes"
	"fmt"
	"strings"
)

func basename1(s string) string {
	for i := len(s); i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s); i <= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma1(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	for i := 0; i < n; i++ {
		if n % 3 == 2 && i == 2 {
			buf.WriteByte(',')
		} else if n % 3 == 1 && i == 1 {
			buf.WriteByte(',')
		} else if i % 3 == 0 && i != 0 && i!= 3 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf, "%c", s[i])
	}
	return buf.String()
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(comma2("13123111"))
}
