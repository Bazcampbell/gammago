// gammago/util_test.go

package gammago

import (
	"net/url"
	"testing"
)

func TestUtils(t *testing.T) {
	t.Run("shouldRetry", func(t *testing.T) {
		tests := []struct {
			name     string
			status   int
			expected bool
		}{
			{"429 Too Many Requests", 429, true},
			{"500 Internal Server Error", 500, true},
			{"502 Bad Gateway", 502, true},
			{"503 Service Unavailable", 503, true},
			{"504 Gateway Timeout", 504, true},

			{"400 Bad Request", 400, false},
			{"401 Unauthorized", 401, false},
			{"403 Forbidden", 403, false},
			{"404 Not Found", 404, false},
			{"422 Unprocessable Entity", 422, false},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := shouldRetry(tt.status, 0); got != tt.expected {
					t.Errorf("shouldRetry(%d) = %v, want %v", tt.status, got, tt.expected)
				}
			})
		}
	})

	t.Run("buildUrl", func(t *testing.T) {
		tests := []struct {
			name     string
			endpoint string
			params   url.Values
			want     string
		}{
			{
				name:     "simple endpoint no params",
				endpoint: "teams",
				params:   nil,
				want:     "https://gamma-api.polymarket.com/teams",
			},
			{
				name:     "endpoint with leading slash",
				endpoint: "/markets",
				params:   nil,
				want:     "https://gamma-api.polymarket.com/markets",
			},
			{
				name:     "endpoint with trailing slash",
				endpoint: "tags/",
				params:   nil,
				want:     "https://gamma-api.polymarket.com/tags",
			},
			{
				name:     "nested endpoint",
				endpoint: "sports/market-types",
				params:   nil,
				want:     "https://gamma-api.polymarket.com/sports/market-types",
			},
			{
				name:     "endpoint with both slashes",
				endpoint: "/tags/slug/abc",
				params:   nil,
				want:     "https://gamma-api.polymarket.com/tags/slug/abc",
			},
			{
				name:     "empty endpoint",
				endpoint: "",
				params:   nil,
				want:     "https://gamma-api.polymarket.com",
			},
			{
				name:     "just slash endpoint",
				endpoint: "/",
				params:   nil,
				want:     "https://gamma-api.polymarket.com",
			},
			{
				name:     "with params",
				endpoint: "events",
				params:   url.Values{"limit": []string{"10"}, "offset": []string{"0"}},
				want:     "https://gamma-api.polymarket.com/events?limit=10&offset=0",
			},
			{
				name:     "with multi-value params",
				endpoint: "teams",
				params:   url.Values{"league": []string{"nfl", "nba"}},
				want:     "https://gamma-api.polymarket.com/teams?league=nfl&league=nba",
			},
			{
				name:     "empty params",
				endpoint: "sports",
				params:   url.Values{},
				want:     "https://gamma-api.polymarket.com/sports",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, _ := buildUrl(tt.endpoint, tt.params)
				if got != tt.want {
					t.Errorf("buildUrl(%q, %v) = %q, want %q", tt.endpoint, tt.params, got, tt.want)
				}
			})
		}
	})
}
