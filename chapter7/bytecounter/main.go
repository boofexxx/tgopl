package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (int, error) {
	*b += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(bytes.NewBuffer(p))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		*w++
	}
	return len(p), nil
}

type LineCounter int

func (w *LineCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(bytes.NewBuffer(p))
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		*w++
	}
	return len(p), nil
}

type Counter struct {
	words int
	lines int
}

func (c *Counter) Write(p []byte) (int, error) {
	isWord := false
	for _, b := range p {
		if unicode.IsSpace(rune(b)) {
			if b == '\n' {
				c.lines++
			}
			if isWord {
				isWord = false
			}
		} else {
			if !isWord {
				c.words++
			}
			isWord = true
		}
	}
	if p[len(p)-1] != '\n' {
		c.lines++
	}
	return len(p), nil
}

type CounterWriter struct {
	writer  io.Writer
	counter int64
}

func (cw *CounterWriter) Write(p []byte) (int, error) {
	n, err := cw.writer.Write(p)
	cw.counter += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CounterWriter{
		writer: w,
		counter: 0,
	}
	return cw.writer, &cw.counter
}

func main() {
	var c Counter
	c.Write([]byte("something helo\nhello\n how are you"))
	fmt.Println(c)
}
