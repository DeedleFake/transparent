package transparent

import "deedles.dev/transparent/internal/rules"

// Clear checks url against the known rules for cleaning a URL. If any
// rules apply, the cleaned URL is returned. If no rules apply, the
// originally passed URL and false are returned.
//
// If url is not a valid URL, the results are undefined.
func Clear(url string) (cleared string, changed bool) {
	for _, provider := range rules.Providers {
		if url == "" {
			return url, changed
		}

		cleared, ok := provider.Clear(url)
		url = cleared
		changed = changed || ok
	}
	return url, changed
}
