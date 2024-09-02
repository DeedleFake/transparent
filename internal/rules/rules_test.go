package rules_test

import (
	"regexp"
	"testing"

	"deedles.dev/transparent/internal/rules"
)

var testProvider = rules.Provider{
	URLPattern:        regexp.MustCompile(`^https?:\/\/(?:[a-z0-9-]+\.)*?amazon(?:\.[a-z]{2,}){1,}`),
	CompleteProvider:  false,
	Rules:             []*regexp.Regexp{regexp.MustCompile(`^p[fd]_rd_[a-z]*$`), regexp.MustCompile(`^qid$`), regexp.MustCompile(`^srs?$`), regexp.MustCompile(`^__mk_[a-z]{1,3}_[a-z]{1,3}$`), regexp.MustCompile(`^spIA$`), regexp.MustCompile(`^ms3_c$`), regexp.MustCompile(`^[a-z%0-9]*ie$`), regexp.MustCompile(`^refRID$`), regexp.MustCompile(`^colii?d$`), regexp.MustCompile(`^[^a-z%0-9]adId$`), regexp.MustCompile(`^qualifier$`), regexp.MustCompile(`^_encoding$`), regexp.MustCompile(`^smid$`), regexp.MustCompile(`^field-lbr_brands_browse-bin$`), regexp.MustCompile(`^ref_?$`), regexp.MustCompile(`^th$`), regexp.MustCompile(`^sprefix$`), regexp.MustCompile(`^crid$`), regexp.MustCompile(`^keywords$`), regexp.MustCompile(`^cv_ct_[a-z]+$`), regexp.MustCompile(`^linkCode$`), regexp.MustCompile(`^creativeASIN$`), regexp.MustCompile(`^ascsubtag$`), regexp.MustCompile(`^aaxitk$`), regexp.MustCompile(`^hsa_cr_id$`), regexp.MustCompile(`^sb-ci-[a-z]+$`), regexp.MustCompile(`^rnid$`), regexp.MustCompile(`^dchild$`), regexp.MustCompile(`^camp$`), regexp.MustCompile(`^creative$`), regexp.MustCompile(`^s$`), regexp.MustCompile(`^content-id$`), regexp.MustCompile(`^dib$`), regexp.MustCompile(`^dib_tag$`)},
	RawRules:          []*regexp.Regexp{regexp.MustCompile(`\/ref=[^/?]*`)},
	ReferralMarketing: []*regexp.Regexp{regexp.MustCompile(`tag`), regexp.MustCompile(`ascsubtag`)},
	Exceptions:        []*regexp.Regexp{regexp.MustCompile(`^https?:\/\/(?:[a-z0-9-]+\.)*?amazon(?:\.[a-z]{2,}){1,}\/gp\/.*?(?:redirector.html|cart\/ajax-update.html|video\/api\/)`), regexp.MustCompile(`^https?:\/\/(?:[a-z0-9-]+\.)*?amazon(?:\.[a-z]{2,}){1,}\/(?:hz\/reviews-render\/ajax\/|message-us\?|s\?)`)},
}

func TestProvioder(t *testing.T) {
	const url = "https://amazon.com/something?ref=3&pf_rd_q=5&other=2"
	if !testProvider.Matches(url) {
		t.Fatal("didn't match")
	}
}
