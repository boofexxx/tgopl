package popcount

import "testing"

const testData = 80

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(testData)
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testData)
	}
}

func BenchmarkPopCountIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountIter(testData)
	}
}

func BenchmarkPopCountClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClear(testData)
	}
}
