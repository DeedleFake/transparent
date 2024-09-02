package rules

//go:generate go run deedles.dev/transparent/internal/cmd/genrules -out rules_gen.go

import (
	"iter"
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

func (p Provider) Matches(value string) bool {
	if !p.URLPattern.MatchString(value) {
		return false
	}
	for _, exception := range p.Exceptions {
		if exception.MatchString(value) {
			return false
		}
	}
	return true
}

func MatchingProviders(value string) iter.Seq[Provider] {
	return func(yield func(Provider) bool) {
		for _, p := range Providers {
			if p.Matches(value) {
				if !yield(p) {
					return
				}
			}
		}
	}
}
