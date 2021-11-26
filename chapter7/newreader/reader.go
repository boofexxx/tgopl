package reader

import (
	"fmt"
	"io"
)

type Reader struct {
	s string
}

func (r *Reader) Read(p []byte) (int, error) {
	n := copy(p, []byte(r.s))
	r.s = r.s[n:]
	if len(r.s) == 0 {
		return n, io.EOF
	}
	return n, nil
}

func NewReader(s string) *Reader {
	r := Reader{
		s: s,
	}
	return &r
}

type LimitedReader struct {
	R io.Reader
	N int64
}

func (l *LimitedReader) Read(p []byte) (int, error) {
	fmt.Println(l.N)
	fmt.Println(l.R)
	if l.N <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.N {
		p = p[:l.N]
	}
	fmt.Println(len(p))
	n, err := l.R.Read(p)
	l.N -= int64(n)
	return n, err
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &LimitedReader{R: r, N: n}
}
