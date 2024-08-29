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

func (p *Provider) Clear(value string) (string, bool) {
	if !p.URLPattern.MatchString(value) {
		return value, false
	}
	for _, exception := range p.Exceptions {
		if exception.MatchString(value) {
			return value, true
		}
	}

	if p.CompleteProvider {
		return "", true
	}

	for _, rule := range p.RawRules {
		value = rule.ReplaceAllLiteralString(value, "")
	}

	parsed, err := url.Parse(value)
	if err != nil {
		return value, true
	}

	for _, rule := range p.Rules {
		query := parsed.Query()
		for k := range query {
			if rule.MatchString(k) {
				query.Del(k)
			}
		}
		parsed.RawQuery = query.Encode()

		// TODO: Handle fragments, too.
	}

	return parsed.String(), true
}
