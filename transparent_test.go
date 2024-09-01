package transparent_test

import (
	"testing"

	"deedles.dev/transparent"
)

func TestClear(t *testing.T) {
	const before = "https://amazon.com/something?ref=3&pf_rd_q=5&other=2"
	after, ok := transparent.Clear(before)
	if !ok {
		t.Fatalf("%q -> %q", before, after)
	}
	if after != "https://amazon.com/something?other=2" {
		t.Fatalf("%q -> %q", before, after)
	}
}

func BenchmarkClear(b *testing.B) {
	const url = "https://x.com/user/status/12389123719273?t=klasdhklashdask&s=33"
	for range b.N {
		_, _ = transparent.Clear(url)
	}
}
