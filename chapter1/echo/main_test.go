package main

import (
	"testing"
)

func BenchmarkEchoUsingJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoUsingJoin(testData)
	}
}

func BenchmarkEchoUsingIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoUsingIteration(testData)
	}
}
