package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
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

	// NOTE: 下記のバグの修正のためyamlを置換する。直ったら戻す
	// [【不具合】schema定義のenum系descriptionが複数行定義になっていない · Issue \#235 · kabucom/kabusapi](https://github.com/kabucom/kabusapi/issues/235)
	body := new(bytes.Buffer)
	scanner := bufio.NewScanner(docResp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "description:") {
			if !strings.Contains(line, "description: |-") {
				lineHead := strings.Index(line, "description:")
				fmt.Fprintln(body, strings.Repeat(" ", lineHead)+"description: |-")
				fmt.Fprintln(body, strings.Repeat(" ", lineHead+2)+line[lineHead+len("description: "):])
				continue
			}
		}
		fmt.Fprintln(body, scanner.Text())
	}
	doc, err := ParseKabusAPIDocument(body)

	// doc, err := ParseKabusAPIDocument(docResp.Body)
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
