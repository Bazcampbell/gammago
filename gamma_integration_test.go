//go:build integration
// +build integration

package gammago

import (
	"testing"
	"time"
)

func TestGammaAPIRealCalls_Marshalling(t *testing.T) {
	// These are smoke tests against the real API
	// They verify:
	// 1. URL builds correctly
	// 2. HTTP call succeeds (or at least gets a response)
	// 3. JSON unmarshals without fatal errors
	// 4. Result has some meaningful data

	t.Run("GetSports - basic list", func(t *testing.T) {
		sports, err := GetSports()
		if err != nil {
			t.Fatalf("GetSports failed: %v", err)
		}
		if len(sports) == 0 {
			t.Error("expected non-empty list of sports, got empty")
		}
		t.Logf("got %d sports", len(sports))
	})

	t.Run("GetTeams - small page", func(t *testing.T) {
		teams, err := GetTeams(5, 0, nil, nil, nil)
		if err != nil {
			t.Fatalf("GetTeams failed: %v", err)
		}
		if len(teams) == 0 {
			t.Error("expected some teams, got zero")
		}
		t.Logf("got %d teams", len(teams))
	})

	t.Run("GetTagBySlug - known existing slug", func(t *testing.T) {
		tag, err := GetTagBySlug("nfl")
		if err != nil {
			t.Fatalf("GetTagBySlug failed: %v", err)
		}
		if tag.ID == "" || tag.Slug == "" {
			t.Errorf("got empty/invalid tag: %+v", tag)
		}
		t.Logf("got tag: %s (id=%s)", tag.Label, tag.ID)
	})

	t.Run("GetMarketsBetweenDates - recent markets", func(t *testing.T) {
		start := time.Now().AddDate(0, -1, 0) // last month
		end := time.Now()

		markets, err := GetMarketsBetweenDates(10, 0, start, end)
		if err != nil {
			t.Fatalf("GetMarketsBetweenDates failed: %v", err)
		}
		if len(markets) == 0 {
			t.Log("warning: no markets found in last month — API might be quiet")
		} else {
			t.Logf("got %d markets", len(markets))
		}
	})

	t.Run("GetTags - pagination smoke", func(t *testing.T) {
		tags, err := GetTags(10, 0)
		if err != nil {
			t.Fatalf("GetTags failed: %v", err)
		}
		if len(tags) == 0 {
			t.Error("expected some tags, got zero")
		}
		t.Logf("got %d tags", len(tags))
	})

	t.Run("GetTags - first tag is Sports", func(t *testing.T) {
		tags, err := GetTags(10, 0)
		if err != nil {
			t.Fatalf("GetTags failed: %v", err)
		}

		if len(tags) == 0 {
			t.Fatal("expected at least one tag, got empty")
		}

		got := tags[0]

		want := Tag{
			ID:    "1",
			Label: "Sports",
			Slug:  "sports",
		}

		if got.ID != want.ID || got.Slug != want.Slug || got.Label != want.Label {
			t.Errorf("unexpected first tag\nwant: %+v\ngot:  %+v", want, got)
		}

		t.Logf("First tag matches expectation: %s (%s)", got.Label, got.Slug)
	})
}
