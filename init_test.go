// gammago/init_test.go

package gammago

import (
	"net/http"
	"sync"
	"testing"
	"time"
)

func resetHTTPClient() {
	httpClient = nil
	once = sync.Once{}
}

func TestInitHTTPClient(t *testing.T) {
	t.Run("uses default client when not initialized", func(t *testing.T) {
		resetHTTPClient()

		client := getHTTPClient()

		if client == nil {
			t.Fatal("expected http client, got nil")
		}

		if client.Timeout != 30*time.Second {
			t.Errorf("default timeout = %v, want %v", client.Timeout, 30*time.Second)
		}
	})

	t.Run("initializes custom client with timeout and transport", func(t *testing.T) {
		resetHTTPClient()

		tr := &http.Transport{}
		InitCustomHttpClient(5, tr)

		client := getHTTPClient()

		if client.Timeout != 5*time.Nanosecond {
			t.Errorf("timeout = %v, want %v", client.Timeout, 5*time.Nanosecond)
		}

		if client.Transport != tr {
			t.Errorf("transport mismatch: got %v, want %v", client.Transport, tr)
		}
	})

	t.Run("InitCustomHttpClient only runs once", func(t *testing.T) {
		resetHTTPClient()

		tr1 := &http.Transport{}
		tr2 := &http.Transport{}

		InitCustomHttpClient(5, tr1)
		InitCustomHttpClient(10, tr2) // should be ignored

		client := getHTTPClient()

		if client.Timeout != 5*time.Nanosecond {
			t.Errorf("timeout = %v, want %v", client.Timeout, 5*time.Nanosecond)
		}

		if client.Transport != tr1 {
			t.Errorf("transport was overwritten, singleton violated")
		}
	})
}
