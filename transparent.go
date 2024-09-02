package transparent

import (
	"net/url"

	"deedles.dev/transparent/internal/rules"
)

// Clear checks value against the known rules for cleaning a URL. If
// any rules apply, the cleaned URL is returned. If no rules apply,
// the originally passed URL and false are returned.
//
// If value is not a valid URL, the results are undefined.
func Clear(value string) (cleared string, changed bool) {
	for p := range rules.MatchingProviders(value) {
		if p.CompleteProvider {
			return "", true
		}

		cleared, ok := clearProvider(p, value)
		value = cleared
		changed = changed || ok
	}
	return value, changed
}

func clearProvider(p rules.Provider, value string) (cleared string, changed bool) {
	for _, rule := range p.RawRules {
		old := value
		value = rule.ReplaceAllLiteralString(value, "")
		changed = changed || value != old
	}

	parsed, err := url.Parse(value)
	if err != nil {
		return value, changed
	}
	query := parsed.Query()

	for _, rule := range p.Rules {
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
