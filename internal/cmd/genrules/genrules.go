package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Provider struct {
	URLPattern        string   `json:"urlPattern"`
	CompleteProvider  bool     `json:"completeProvider"`
	Rules             []string `json:"rules"`
	RawRules          []string `json:"rawRules"`
	ReferralMarketing []string `json:"referralMarketing"`
	Exceptions        []string `json:"exceptions"`
	Redirections      []string `json:"redirections"`
	ForceRedirection  bool     `json:"forceRedirection"`
}

func GetRules() (map[string]Provider, error) {
	const rulesURL = "https://rules2.clearurls.xyz/data.minify.json"

	rsp, err := http.Get(rulesURL)
	if err != nil {
		return nil, fmt.Errorf("get rules file: %w", err)
	}
	defer rsp.Body.Close()

	buf, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("read rules file: %w", err)
	}

	var data struct {
		Providers map[string]Provider `json:"providers"`
	}
	err = json.Unmarshal(buf, &data)
	if err != nil {
		return nil, fmt.Errorf("decode rules file: %w", err)
	}

	return data.Providers, nil
}

func main() {
	rules, err := GetRules()
	if err != nil {
		panic(err)
	}

	for p := range rules {
		fmt.Printf("%q\n", p)
	}
}
