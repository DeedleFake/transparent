package transparent

import (
	"iter"
	"net/url"
	"regexp"
	"slices"

	"deedles.dev/transparent/internal/rules"
)

// Clear checks value against the known rules for cleaning a URL. If
// any rules apply, the cleaned URL is returned. If no rules apply,
// the originally passed URL and false are returned.
//
// If value is not a valid URL, the results are undefined.
func Clear(value string) (cleared string, changed bool) {
	if rules.CompleteProviderMatches(value) {
		return "", value != ""
	}

	providers := slices.Values(slices.Collect(rules.MatchingProviders(value)))

	for rule := range allRawRules(providers) {
		old := value
		value = rule.ReplaceAllLiteralString(value, "")
		changed = changed || value != old
	}

	parsed, err := url.Parse(value)
	if err != nil {
		return value, changed
	}
	query := parsed.Query()

	for rule := range allRules(providers) {
		for k := range query {
			if rule.MatchString(k) {
				changed = true
				query.Del(k)
			}
		}

		// TODO: Handle fragments, too.
	}

	parsed.RawQuery = query.Encode()
	return parsed.String(), changed
}

func allRawRules(providers iter.Seq[rules.Provider]) iter.Seq[*regexp.Regexp] {
	return func(yield func(*regexp.Regexp) bool) {
		for p := range providers {
			for _, rule := range p.RawRules {
				if !yield(rule) {
					return
				}
			}
		}
	}
}

func allRules(providers iter.Seq[rules.Provider]) iter.Seq[*regexp.Regexp] {
	return func(yield func(*regexp.Regexp) bool) {
		for p := range providers {
			for _, rule := range p.Rules {
				if !yield(rule) {
					return
				}
			}
		}
	}
}
