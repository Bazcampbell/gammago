// gammago/init.go

package gammago

import (
	"net/http"
	"sync"
	"time"
)

var (
	once       sync.Once
	httpClient *http.Client
)

// Initialise http client with timeout
// Singleton pattern
func InitCustomHttpClient(timeout int, transport *http.Transport) {
	once.Do(func() {
		httpClient = &http.Client{
			Timeout:   time.Duration(timeout),
			Transport: transport,
		}
	})
}

// Or use a default if not initialized
func getHTTPClient() *http.Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 30 * time.Second,
		}
	}
	return httpClient
}
