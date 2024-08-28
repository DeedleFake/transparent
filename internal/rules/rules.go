package rules

//go:generate go run deedles.dev/transparent/internal/cmd/genrules -out rules_gen.go

import "regexp"

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

func (p *Provider) Clear(url string) string {
	if !p.URLPattern.MatchString(url) {
		return url
	}
	if p.CompleteProvider {
		return ""
	}

	for _, rule := range p.Rules {
		url = rule.ReplaceAllString(url, "")
	}

	return url
}
