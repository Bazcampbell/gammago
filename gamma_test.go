// gammago/gamma_test.go

package gammago

import (
	"net/url"
	"strconv"
	"testing"
	"time"
)

func TestGammaEndpoints(t *testing.T) {
	t.Run("GetTeams URL construction & param handling", func(t *testing.T) {
		tests := []struct {
			name         string
			limit        int
			offset       int
			league       []string
			nameFilter   []string
			abbreviation []string
			wantQuery    string
			wantErr      bool
		}{
			{
				name:      "basic pagination only",
				limit:     25,
				offset:    50,
				wantQuery: "limit=25&offset=50",
				wantErr:   false,
			},
			{
				name:         "with multiple filter values",
				limit:        10,
				offset:       0,
				league:       []string{"nfl", "ncaa"},
				nameFilter:   []string{"patriots", "eagles"},
				abbreviation: []string{"NE"},
				wantQuery:    "abbreviation=NE&league=nfl&league=ncaa&limit=10&name=patriots&name=eagles&offset=0",
				wantErr:      false,
			},
			{
				name:      "negative limit (allowed by client)",
				limit:     -5,
				offset:    0,
				wantQuery: "limit=-5&offset=0",
				wantErr:   false,
			},
			{
				name:       "empty filter slices",
				limit:      100,
				offset:     0,
				league:     []string{},
				nameFilter: nil,
				wantQuery:  "limit=100&offset=0",
				wantErr:    false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// We only care about URL construction here, not the actual HTTP call
				params := url.Values{}
				params.Add("limit", strconv.Itoa(tt.limit))
				params.Add("offset", strconv.Itoa(tt.offset))
				for _, v := range tt.league {
					params.Add("league", v)
				}
				for _, v := range tt.nameFilter {
					params.Add("name", v)
				}
				for _, v := range tt.abbreviation {
					params.Add("abbreviation", v)
				}

				got, err := buildUrl("teams", params)
				if (err != nil) != tt.wantErr {
					t.Errorf("buildUrl error = %v, wantErr %v", err, tt.wantErr)
				}
				if !tt.wantErr {
					u, _ := url.Parse(got)
					if u.RawQuery != tt.wantQuery {
						t.Errorf("query string mismatch\nwant: %s\ngot:  %s", tt.wantQuery, u.RawQuery)
					}
				}
			})
		}
	})

	t.Run("GetTagBySlug input validation", func(t *testing.T) {
		tests := []struct {
			name    string
			slug    string
			wantErr bool
		}{
			{"normal slug", "nfl", false},
			{"slug with hyphens", "ncaa-football", false},
			{"random slug", "mike-hunt-420", true},
			{"empty string", "", true},
			{"only spaces", "   ", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := GetTagBySlug(tt.slug)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetTagBySlug(%q) error = %v, wantErr %v", tt.slug, err, tt.wantErr)
				}
			})
		}
	})

	t.Run("GetEventByID input validation", func(t *testing.T) {
		tests := []struct {
			name    string
			id      string
			wantErr bool
		}{
			{"single id", "12345", false},
			{"letter id", "abd", true},
			{"empty id", "", true},
			{"malformed id", "?&//", true},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := GetEventByID(tt.id)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetEventByID(%v) error = %v, wantErr %v", tt.id, err, tt.wantErr)
				}
			})
		}
	})

	t.Run("date formatting consistency (GetMarketsBetweenDates)", func(t *testing.T) {
		tests := []struct {
			name      string
			start     time.Time
			end       time.Time
			wantStart string
			wantEnd   string
		}{
			{
				name:      "standard UTC date",
				start:     time.Date(2025, 10, 15, 0, 0, 0, 0, time.UTC),
				end:       time.Date(2025, 12, 31, 23, 59, 59, 0, time.UTC),
				wantStart: "2025-10-15T00:00:00Z",
				wantEnd:   "2025-12-31T23:59:59Z",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				params := url.Values{}
				params.Add("start_date_min", tt.start.Format("2006-01-02T15:04:05Z"))
				params.Add("end_date_max", tt.end.Format("2006-01-02T15:04:05Z"))

				got, err := buildUrl("markets", params)
				if err != nil {
					t.Fatalf("unexpected build error: %v", err)
				}

				u, _ := url.Parse(got)
				q := u.Query()
				if q.Get("start_date_min") != tt.wantStart {
					t.Errorf("start_date_min: want %q got %q", tt.wantStart, q.Get("start_date_min"))
				}
				if q.Get("end_date_max") != tt.wantEnd {
					t.Errorf("end_date_max: want %q got %q", tt.wantEnd, q.Get("end_date_max"))
				}
			})
		}
	})
}
