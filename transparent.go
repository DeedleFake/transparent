package transparent

import "deedles.dev/transparent/internal/rules"

// Clear checks url against the known rules for cleaning a URL. If any
// rules apply, the cleaned URL is returned. If no rules apply, the
// originally passed URL and false are returned.
//
// If url is not a valid URL, the results are undefined.
func Clear(url string) (string, bool) {
	for _, provider := range rules.Providers {
		if url, ok := provider.Clear(url); ok {
			return url, true
		}
	}
	return url, false
}
