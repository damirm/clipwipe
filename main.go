package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"golang.design/x/clipboard"
)

var defaultParams = []string{
	"utm_source",
	"utm_medium",
	"utm_campaign",
	"utm_content",
	"utm_term",
	"fbclid",
	"gclid",
}

func main() {
	interval := flag.Duration("interval", 500*time.Millisecond, "clipboard polling interval")
	debug := flag.Bool("debug", false, "enable debug output")
	params := flag.String("params", strings.Join(defaultParams, ","), "comma-separated list of query parameters to remove")
	flag.Parse()

	// Parse parameters to remove
	paramsToIgnore := parseParams(*params)

	// Initialize clipboard
	if err := clipboard.Init(); err != nil {
		log.Fatalf("Failed to initialize clipboard: %v", err)
	}

	fmt.Println("Clipboard UTM cleaner started. Watching for URLs with UTM parameters...")
	fmt.Printf("Polling interval: %v\n", *interval)
	fmt.Printf("Parameters to remove: %v\n", paramsToIgnore)
	fmt.Println("Press Ctrl+C to stop.")

	var lastContent string

	for {
		// Read clipboard content
		data := clipboard.Read(clipboard.FmtText)

		if *debug {
			fmt.Printf("[DEBUG] Clipboard content: %q (changed: %v)\n", string(data), string(data) != lastContent)
		}

		if len(data) > 0 {
			currentContent := string(data)
			// Check if content has changed
			if currentContent != lastContent {
				cleanedURL, ok := removeQueryParams(currentContent, paramsToIgnore)
				if ok && cleanedURL != currentContent {
					// Write cleaned URL back to clipboard
					clipboard.Write(clipboard.FmtText, []byte(cleanedURL))
					fmt.Printf("Cleaned: %s\n-> %s\n", currentContent, cleanedURL)
				}
				lastContent = currentContent
			}
		}

		time.Sleep(*interval)
	}
}

// parseParams parses a comma-separated string into a slice of strings
func parseParams(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		trimmed := strings.TrimSpace(p)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

// removeQueryParams removes specified query parameters from a URL.
// Returns the cleaned URL and true if the URL was modified.
// Trims leading/trailing whitespace and newlines from the input.
func removeQueryParams(rawURL string, paramsToRemove []string) (string, bool) {
	// Trim whitespace and newlines
	trimmed := strings.TrimSpace(rawURL)
	if trimmed == "" {
		return rawURL, false
	}

	// Check if string starts with http:// or https://
	if !strings.HasPrefix(trimmed, "http://") && !strings.HasPrefix(trimmed, "https://") {
		return rawURL, false
	}

	// Parse URL
	parsedURL, err := url.Parse(trimmed)
	if err != nil {
		return rawURL, false
	}

	// Get all query parameters
	query := parsedURL.Query()
	modified := false

	// Remove specified parameters
	for _, param := range paramsToRemove {
		if query.Has(param) {
			query.Del(param)
			modified = true
		}
	}

	if modified {
		parsedURL.RawQuery = query.Encode()
		return parsedURL.String(), true
	}

	return rawURL, false
}
