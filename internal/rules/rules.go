package rules

//go:generate go run deedles.dev/transparent/internal/cmd/genrules -out rules_gen.go

import (
	"net/url"
	"regexp"
)

type Provider struct {
	URLPattern        *regexp.Regexp
	CompleteProvider  bool
	Rules             []*regexp.Regexp
	RawRules          []*regexp.Regexp
	ReferralMarketing []*regexp.Regexp
	Exceptions        []*regexp.Regexp
	Redirections      []*regexp.Regexp
	ForceRedirection  bool
}

func (p *Provider) Clear(value string) (cleared string, changed bool) {
	if !p.URLPattern.MatchString(value) {
		return value, false
	}
	for _, exception := range p.Exceptions {
		if exception.MatchString(value) {
			return value, false
		}
	}

	if p.CompleteProvider {
		return "", true
	}

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
