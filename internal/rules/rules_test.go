package rules_test

import (
	"testing"

	"deedles.dev/transparent/internal/rules"
)

func TestClear(t *testing.T) {
	const before = "https://9gag.com/something?ref=3&other=2"
	after := rules.Providers[0].Clear(before)
	if after != "https://9gag.com/something?other=2" {
		t.Fatalf("%q -> %q", before, after)
	}
}
