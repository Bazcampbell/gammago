// gammago/util.go

package gammago

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	maxRetries = 3
	baseDelay  = 800 * time.Millisecond
)

// Send a GET request to a given URL
// Add parameter headers
// Attempt to unmarshal the response into T
func genericGet[T any](url string) (T, error) {
	var result T
	var lastErr error

	client := getHTTPClient()

	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			// exponential backoff + jitter
			delay := baseDelay * time.Duration(1<<attempt)
			jitter := time.Duration(time.Now().UnixNano()%100) * time.Millisecond
			time.Sleep(delay + jitter)
		}

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return result, fmt.Errorf("create request failed: %w", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			lastErr = err
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			lastErr = fmt.Errorf("read body failed: %w", err)
			continue
		}

		// defer won't hit on retries
		resp.Body.Close()

		// retry certain statuses
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			lastErr = fmt.Errorf("status code: %d, %s", resp.StatusCode, body)

			if !shouldRetry(resp.StatusCode, attempt) {
				break
			}
			continue
		}

		if err = json.Unmarshal(body, &result); err != nil {
			lastErr = fmt.Errorf("json unmarshal failed: %w", err)
			continue
		}

		return result, nil
	}

	if lastErr == nil {
		lastErr = fmt.Errorf("request failed after %d attempts (unknown reason)", maxRetries)
	}

	return result, fmt.Errorf("%w (after %d attempts)", lastErr, maxRetries)
}

func shouldRetry(status int, attempt int) bool {
	if attempt >= maxRetries-1 {
		return false // last attempt anyway
	}

	// always retry these
	if status == 429 || (status >= 500 && status <= 599) {
		return true
	}

	// usually don't retry these other status codes
	return false
}

// buildUrl constructs a full URL from base + endpoint + query params.
// Returns error on invalid base URL or malformed input.
func buildUrl(endpoint string, params url.Values) (string, error) {
	base := strings.TrimRight(BASE_URL, "/") // normalize base

	u, err := url.Parse(base)
	if err != nil {
		return "", fmt.Errorf("invalid base URL %q: %w", BASE_URL, err)
	}

	// Clean endpoint: remove leading/trailing slashes
	endpoint = strings.Trim(endpoint, "/")

	// Only append if there's meaningful endpoint content
	if endpoint != "" {
		u.Path = u.Path + "/" + endpoint
	}

	// Add query params only if we have any
	if len(params) > 0 {
		u.RawQuery = params.Encode()
	}

	return u.String(), nil
}
