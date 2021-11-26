package reader

import (
	"bytes"
	"fmt"
	"testing"

	"golang.org/x/net/html"
)

func TestNewReader(t *testing.T) {
	s := "hello 世界"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(NewReader(s))
	if n != int64(len(s)) || err != nil {
		t.Errorf("n=%d err=%s", n, err)
	}
	if b.String() != s {
		t.Errorf("\"%s\" != \"%s\"", b.String(), s)
	}
}

func TestReaderWithHTML(t *testing.T) {
	s := "<html><head><title>some</title></head><body>hi</body></html>"
	_, err := html.Parse(NewReader(s))
	if err != nil {
		t.Error(err)
	}
}

// fails
func TestLimitedReader(t *testing.T) {
	lr := LimitReader(NewReader("something"), 5)
	var buf []byte
	_, err := lr.Read(buf)
	if err != nil {
		t.Fatal(err)
	}
	if string(buf) != "somet" {
		t.Fail()
	}
	fmt.Println("buf:", buf)
}
