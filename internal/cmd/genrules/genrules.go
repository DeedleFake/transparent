package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log/slog"
	"net/http"
	"os"
	"text/template"
)

//go:embed output.tmpl
var outputSource string

var tmpl = template.Must(template.New("output").Parse(outputSource))

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
	pkgname := flag.String("package", "rules", "package name of output file")
	outpath := flag.String("out", "", "path to output file or blank for stdout")
	flag.Parse()

	out := os.Stdout
	if *outpath != "" {
		file, err := os.Create(*outpath)
		if err != nil {
			slog.Error("create output file", "err", err)
			os.Exit(1)
		}
		defer file.Close()
		out = file
	}

	providers, err := GetRules()
	if err != nil {
		slog.Error("get rules", "err", err)
		os.Exit(1)
	}

	var buf bytes.Buffer

	data := map[string]any{
		"Package":   *pkgname,
		"Providers": providers,
	}
	err = tmpl.Execute(&buf, data)
	if err != nil {
		slog.Error("generate output", "err", err)
		os.Exit(1)
	}

	output, err := format.Source(buf.Bytes())
	if err != nil {
		slog.Error("format output", "err", err)
		os.Exit(1)
	}

	_, err = out.Write(output)
	if err != nil {
		slog.Error("write output", "err", err)
		os.Exit(1)
	}
}
