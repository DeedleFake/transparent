package rules_test

import (
	"regexp"
	"testing"

	"deedles.dev/transparent/internal/rules"
)

var testProvider = rules.Provider{
	URLPattern:       regexp.MustCompile(`^https?:\/\/(?:[a-z0-9-]+\.)*?9gag\.com`),
	CompleteProvider: false,
	Rules:            []*regexp.Regexp{regexp.MustCompile(`^ref$`)},
	Exceptions:       []*regexp.Regexp{regexp.MustCompile(`^https?:\/\/comment-cdn\.9gag\.com\/.*?comment-list.json\?`)},
}

func TestClear(t *testing.T) {
	const before = "https://9gag.com/something?ref=3&other=2"
	after, ok := testProvider.Clear(before)
	if !ok {
		t.Fatalf("%q -> %q", before, after)
	}
	if after != "https://9gag.com/something?other=2" {
		t.Fatalf("%q -> %q", before, after)
	}
}
