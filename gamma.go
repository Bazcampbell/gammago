// gammago/gamma.go

package gammago

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
)

const BASE_URL = "https://gamma-api.polymarket.com"

// GetTeams gets teams with optional filters
func GetTeams(limit, offset int, league, name, abbreviation []string) ([]Team, error) {
	params := url.Values{}
	params.Add("order", "id")
	params.Add("limit", strconv.Itoa(limit))
	params.Add("offset", strconv.Itoa(offset))

	for _, l := range league {
		params.Add("league", l)
	}
	for _, n := range name {
		params.Add("name", n)
	}
	for _, a := range abbreviation {
		params.Add("abbreviation", a)
	}

	reqUrl, _ := buildUrl("teams", params)

	return genericGet[[]Team](reqUrl)
}

// GetSports gets all sports
func GetSports() ([]Sport, error) {
	params := url.Values{}
	params.Add("order", "id")

	reqUrl, _ := buildUrl("sports", params)

	return genericGet[[]Sport](reqUrl)
}

// GetMarketTypes gets market types
func GetMarketTypes() (MarketTypes, error) {
	params := url.Values{}
	params.Add("order", "id")

	reqUrl, _ := buildUrl("sports/market-types", params)

	return genericGet[MarketTypes](reqUrl)
}

// GetTags gets tags with pagination
func GetTags(limit, offset int) ([]Tag, error) {
	params := url.Values{}
	params.Add("order", "id")
	params.Add("limit", strconv.Itoa(limit))
	params.Add("offset", strconv.Itoa(offset))

	reqUrl, _ := buildUrl("tags", params)

	return genericGet[[]Tag](reqUrl)
}

// GetTagBySlug gets a specific tag by its slug
func GetTagBySlug(slug string) (Tag, error) {
	params := url.Values{}
	params.Add("order", "id")

	reqUrl, _ := buildUrl(fmt.Sprintf("tags/slug/%s", slug), params)
	return genericGet[Tag](reqUrl)
}

// GetRelatedTagsBySlug gets related tags for a given tag slug
func GetRelatedTagsByTagId(id int) ([]Tag, error) {
	params := url.Values{}
	params.Add("order", "id")

	reqUrl, _ := buildUrl(fmt.Sprintf("tags/%d/related-tags/tags", id), params)
	return genericGet[[]Tag](reqUrl)
}

// GetEventsByTag gets events by tag ID
func GetEventsByTag(tagID int, includeRelated bool) ([]Event, error) {
	params := url.Values{}
	params.Add("order", "id")
	params.Add("tag_id", strconv.Itoa(tagID))
	params.Add("related_tags", strconv.FormatBool(includeRelated))

	reqUrl, _ := buildUrl("events", params)
	return genericGet[[]Event](reqUrl)
}

// GetEventByID gets events by their IDs
func GetEventByID(id string) (Event, error) {
	reqUrl, _ := buildUrl(fmt.Sprintf("events/%s", id), nil)
	return genericGet[Event](reqUrl)
}

// GetEventsBeforeDate gets ALL events ending before a specific date
func GetEventsBeforeDate(
	limit,
	offset,
	volumeMin,
	tagId int,
	endDate time.Time,
	status Status,
) ([]Event, error) {
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("offset", strconv.Itoa(offset))
	params.Add("tag_id", strconv.Itoa(tagId))
	params.Add("order", "volume")
	params.Add("end_date_max", endDate.Format("2006-01-02T15:04:05Z"))
	params.Add("volume_min", strconv.Itoa(volumeMin))

	if status == ACTIVE {
		params.Add("active", "true")
		params.Add("closed", "false")
	}

	if status == CLOSED {
		params.Add("active", "false")
		params.Add("closed", "true")
	}

	reqUrl, _ := buildUrl("events", params)
	return genericGet[[]Event](reqUrl)
}

// GetEventsBetweenDates gets events starting and ending between two dates
func GetEventsBetweenDates(
	limit,
	offset,
	volumeMin,
	tagId int,
	endDate time.Time,
	startDate time.Time,
	status Status,
) ([]Event, error) {
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("offset", strconv.Itoa(offset))
	params.Add("tag_id", strconv.Itoa(tagId))
	params.Add("order", "volume")
	params.Add("end_date_max", endDate.Format("2006-01-02T15:04:05Z"))
	params.Add("start_date_min", startDate.Format("2006-01-02T15:04:05Z"))
	params.Add("volume_min", strconv.Itoa(volumeMin))

	if status == ACTIVE {
		params.Add("active", "true")
		params.Add("closed", "false")
	}

	if status == CLOSED {
		params.Add("active", "false")
		params.Add("closed", "true")
	}

	reqUrl, _ := buildUrl("events", params)
	return genericGet[[]Event](reqUrl)
}

// GetMarketsBetweenDates gets markets between specified dates
func GetMarketsBetweenDates(limit, offset int, startDate, endDate time.Time) ([]Market, error) {
	params := url.Values{}
	params.Add("order", "id")
	params.Add("limit", strconv.Itoa(limit))
	params.Add("offset", strconv.Itoa(offset))
	params.Add("start_date_min", startDate.Format("2006-01-02T15:04:05Z"))
	params.Add("end_date_max", endDate.Format("2006-01-02T15:04:05Z"))

	reqUrl, _ := buildUrl("markets", params)
	return genericGet[[]Market](reqUrl)
}

// GetMarketByID gets a market by its ID
func GetMarketByID(marketID int) ([]Market, error) {
	params := url.Values{}
	params.Add("order", "id")
	params.Add("id", strconv.Itoa(marketID))

	reqUrl, _ := buildUrl("markets", params)
	return genericGet[[]Market](reqUrl)
}
