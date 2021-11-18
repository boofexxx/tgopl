package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var testData = `
<html>
<head>
	<title>Title</title>
</head>
<body>
	<img src="some.png">
</body>
</html>
`

// I have no imagination how to test
func TestForEachNode(t *testing.T) {
	doc, err := html.Parse(strings.NewReader(testData))
	if err != nil {
		t.Error(err)
	}
	forEachNode(doc, startElement, endElement)
}
