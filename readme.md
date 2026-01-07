Gamma API Client (Polymarket)

A lightweight Go client for interacting with the Polymarket Gamma API.

This package provides simple, typed helper functions for fetching teams, sports, tags, events, and markets from:

https://gamma-api.polymarket.com

It is designed to be:

- Minimal
- Dependency-free
- Safe to use out of the box
- Optionally configurable via a custom HTTP client


Installation

```bash
go get github.com/Bazcampbell/gamma-api-go
```

Quick Start

You can use the package immediately without any setup.

Example:

```go
package main

import (
    "log"

    gamma "github.com/Bazcampbell/gamma-api-go"
)

func main() {
    teams, err := gamma.GetTeams(
        50,
        0,
        []string{"nfl"},
        nil,
        nil,
    )
    if err != nil {
        log.Fatal(err)
    }
}
```
No explicit initialisation is required.


Optional: Custom HTTP Client

By default, the package uses a standard http.Client with a 30 second timeout.

If you want to customize timeouts, transports, proxies, or TLS settings, you can initialize a custom client once.

Example:
```go
gamma.InitCustomHttpClient(
    int(10*time.Second),
    &http.Transport{
        MaxIdleConns:         100,
        IdleConnTimeout:     90 * time.Second,
        TLSHandshakeTimeout: 10 * time.Second,
    },
)
```
Notes:
- Uses a singleton pattern
- Safe to call multiple times, but only the first call has effect
- Completely optional

Pretty Printing

All types implement the fmt.Stringer interface with formatted output for easy debugging and logging:

```go
event, err := gamma.GetEventByID([]string{"123456"})
if err != nil {
    log.Fatal(err)
}

// Automatically pretty-printed
fmt.Println(event)
// Output:
// Event{
//   ID: 123456
//   Title: Lakers vs Celtics
//   Slug: lakers-vs-celtics
//   Active: true
//   StartDate: 2026-01-15 19:30:00
//   EndDate: 2026-01-15 22:00:00
//   Markets: [3 markets]
//     - Will Lakers win? (ID: abc123)
//     - Total points over 220.5? (ID: def456)
//     - Point spread -5.5 (ID: ghi789)
//   ...
// }
```

API Endpoints

Base URL:
https://gamma-api.polymarket.com


Teams

GET /teams

GetTeams(limit, offset, league, name, abbreviation)

Query parameters:
- limit
- offset
- league (repeatable)
- name (repeatable)
- abbreviation (repeatable)


Sports

GET /sports

GetSports()


Market Types

GET /sports/market-types

GetMarketTypes()


Tags

GET /tags

GetTags(limit, offset)


Tag by Slug

GET /tags/slug/{slug}

GetTagBySlug(slug)


Related Tags

GET /tags/{id}/related-tags/tags

GetRelatedTagsByTagId(id)


Events by Tag

GET /events

GetEventsByTag(tagID, includeRelated)

Query parameters:
- tag_id
- related_tags
- order=id


Events by ID

GET /events

GetEventByID(ids)

Query parameter:
- id (comma-separated)


Events Before Date

GET /events

GetEventsBeforeDate(limit, offset, volumeMin, tagId, endDate, status)

Query parameters:
- limit
- offset
- tag_id
- volume_min
- end_date_max
- active / closed
- order=volume


Markets Between Dates

GET /markets

GetMarketsBetweenDates(limit, offset, startDate, endDate)


Market by ID

GET /markets

GetMarketByID(marketID)


Date Formatting

All timestamps are formatted as:

2006-01-02T15:04:05Z


Design Notes

- Thin wrapper over the Gamma REST API
- No retries
- No rate limiting
- No caching
- Safe for concurrent use
