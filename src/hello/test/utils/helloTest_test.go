package utils

import (
	"testing"
)

// go test -v
func TestHelloTestTest(t *testing.T) {
	if HelloTest("test") {
		t.Log("pass")
	} else {
		t.Error("fail")
	}
}

// go test -test.bench=".*"
func BenchmarkHelloTest(b *testing.B) {
	b.StopTimer()
	b.Log("init")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		HelloTest("test")
	}
}
