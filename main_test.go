package main

import (
	"testing"
)

func TestParseParams(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "default params",
			input:    "utm_source,utm_medium,utm_campaign",
			expected: []string{"utm_source", "utm_medium", "utm_campaign"},
		},
		{
			name:     "with spaces",
			input:    "utm_source, utm_medium , utm_campaign",
			expected: []string{"utm_source", "utm_medium", "utm_campaign"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: nil,
		},
		{
			name:     "single param",
			input:    "utm_source",
			expected: []string{"utm_source"},
		},
		{
			name:     "with empty parts",
			input:    "utm_source,,utm_medium",
			expected: []string{"utm_source", "utm_medium"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseParams(tt.input)
			if len(result) != len(tt.expected) {
				t.Fatalf("parseParams(%q) returned %d elements, expected %d", tt.input, len(result), len(tt.expected))
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("parseParams(%q)[%d] = %q, expected %q", tt.input, i, v, tt.expected[i])
				}
			}
		})
	}
}

func TestRemoveQueryParams(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		paramsToRemove []string
		expectedURL    string
		expectedMod    bool
	}{
		{
			name:           "remove utm params",
			input:          "https://example.com/page?utm_source=google&utm_medium=cpc&id=123",
			paramsToRemove: []string{"utm_source", "utm_medium"},
			expectedURL:    "https://example.com/page?id=123",
			expectedMod:    true,
		},
		{
			name:           "remove all params",
			input:          "https://example.com/page?utm_source=google&utm_campaign=sale",
			paramsToRemove: []string{"utm_source", "utm_campaign"},
			expectedURL:    "https://example.com/page",
			expectedMod:    true,
		},
		{
			name:           "no matching params",
			input:          "https://example.com/page?id=123&name=test",
			paramsToRemove: []string{"utm_source", "utm_medium"},
			expectedURL:    "https://example.com/page?id=123&name=test",
			expectedMod:    false,
		},
		{
			name:           "not a URL",
			input:          "just some text",
			paramsToRemove: []string{"utm_source"},
			expectedURL:    "just some text",
			expectedMod:    false,
		},
		{
			name:           "URL without params",
			input:          "https://example.com/page",
			paramsToRemove: []string{"utm_source"},
			expectedURL:    "https://example.com/page",
			expectedMod:    false,
		},
		{
			name:           "http URL",
			input:          "http://example.com/page?utm_source=facebook",
			paramsToRemove: []string{"utm_source"},
			expectedURL:    "http://example.com/page",
			expectedMod:    true,
		},
		{
			name:           "with whitespace",
			input:          "  https://example.com/page?utm_source=test\n",
			paramsToRemove: []string{"utm_source"},
			expectedURL:    "https://example.com/page",
			expectedMod:    true,
		},
		{
			name:           "multiple params same type",
			input:          "https://example.com/page?utm_source=google&utm_source=bing&id=1",
			paramsToRemove: []string{"utm_source"},
			expectedURL:    "https://example.com/page?id=1",
			expectedMod:    true,
		},
		{
			name:           "preserve other params",
			input:          "https://example.com/page?ref=home&utm_source=google&page=2",
			paramsToRemove: []string{"utm_source"},
			expectedURL:    "https://example.com/page?page=2&ref=home",
			expectedMod:    true,
		},
		{
			name:           "empty input",
			input:          "",
			paramsToRemove: []string{"utm_source"},
			expectedURL:    "",
			expectedMod:    false,
		},
		{
			name:           "fbclid param",
			input:          "https://example.com/page?fbclid=abc123&utm_source=google",
			paramsToRemove: []string{"fbclid", "utm_source"},
			expectedURL:    "https://example.com/page",
			expectedMod:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultURL, modified := removeQueryParams(tt.input, tt.paramsToRemove)
			if modified != tt.expectedMod {
				t.Errorf("removeQueryParams(%q) modified = %v, expected %v", tt.input, modified, tt.expectedMod)
			}
			if resultURL != tt.expectedURL {
				t.Errorf("removeQueryParams(%q) = %q, expected %q", tt.input, resultURL, tt.expectedURL)
			}
		})
	}
}
