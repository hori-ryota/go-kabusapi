package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := generate(); err != nil {
		log.Fatal(err)
	}
}

func generate() error {
	var dstDir string
	flag.StringVar(&dstDir, "dstDir", ".", "output directory")
	var docURL string
	flag.StringVar(&docURL, "docURL", "https://raw.githubusercontent.com/kabucom/kabusapi/master/reference/kabu_STATION_API.yaml", "URL of kabuステーションAPI YAML document")
	flag.Parse()

	if dstDir == "" {
		dstDir = "."
	}

	docResp, err := http.Get(docURL)
	if err != nil {
		return fmt.Errorf("failed to fetch yaml document: url=%s: %w", docURL, err)
	}
	defer docResp.Body.Close()

	doc, err := ParseKabusAPIDocument(docResp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse document: %w", err)
	}

	if err := GenerateStructs(dstDir, "kabusapi", doc); err != nil {
		return fmt.Errorf("failed to GenerateStructs: %w", err)
	}
	if err := GenerateRequests(dstDir, "kabusapi", doc); err != nil {
		return fmt.Errorf("failed to GenerateRequests: %w", err)
	}
	return nil
}
