// pkg/types_string.go

package gammago

// String representation of all types for pretty printing.

import (
	"fmt"
	"strings"
)

func (s Status) String() string {
	return string(s)
}

func (t Team) String() string {
	return fmt.Sprintf(`Team{
  ID: %d
  Name: %s
  League: %s
  Abbreviation: %s
  Alias: %s
  Color: %s
  Logo: %s
}`, t.ID, t.Name, t.League, t.Abbreviation, t.Alias, t.Color, t.Logo)
}

func (s Sport) String() string {
	return fmt.Sprintf(`Sport{
  Sport: %s
  Image: %s
  Resolution: %s
  Ordering: %s
  Tags: %s
  Series: %s
}`, s.Sport, s.Image, s.Resolution, s.Ordering, s.Tags, s.Series)
}

func (m MarketTypes) String() string {
	return fmt.Sprintf(`MarketTypes{
  MarketTypes: [%s]
}`, strings.Join(m.MarketTypes, ", "))
}

func (t Tag) String() string {
	return fmt.Sprintf(`Tag{ID: %s, Label: %s, Slug: %s}`, t.ID, t.Label, t.Slug)
}

func (c Category) String() string {
	return fmt.Sprintf(`Category{ID: %s, Label: %s, Slug: %s, ParentCategory: %s}`,
		c.ID, c.Label, c.Slug, c.ParentCategory)
}

func (c Chat) String() string {
	return fmt.Sprintf(`Chat{
  ID: %s
  ChannelID: %s
  ChannelName: %s
  ChannelImage: %s
  Live: %t
}`, c.ID, c.ChannelID, c.ChannelName, c.ChannelImage, c.Live)
}

func (c Collection) String() string {
	var sb strings.Builder
	sb.WriteString("Collection{\n")
	sb.WriteString(fmt.Sprintf("  ID: %s\n", c.ID))
	sb.WriteString(fmt.Sprintf("  Ticker: %s\n", c.Ticker))
	sb.WriteString(fmt.Sprintf("  Slug: %s\n", c.Slug))
	sb.WriteString(fmt.Sprintf("  Title: %s\n", c.Title))
	sb.WriteString(fmt.Sprintf("  Subtitle: %s\n", c.Subtitle))
	sb.WriteString(fmt.Sprintf("  CollectionType: %s\n", c.CollectionType))
	sb.WriteString(fmt.Sprintf("  Active: %t\n", c.Active))
	sb.WriteString("}")
	return sb.String()
}

func (m Market) String() string {
	var sb strings.Builder
	sb.WriteString("Market{\n")
	sb.WriteString(fmt.Sprintf("  ID: %s\n", m.ID))
	sb.WriteString(fmt.Sprintf("  Question: %s\n", m.Question))
	sb.WriteString(fmt.Sprintf("  Slug: %s\n", m.Slug))
	sb.WriteString(fmt.Sprintf("  MarketType: %s\n", m.MarketType))
	sb.WriteString(fmt.Sprintf("  SportsMarketType: %s\n", m.SportsMarketType))
	sb.WriteString(fmt.Sprintf("  Active: %t\n", m.Active))
	sb.WriteString(fmt.Sprintf("  StartDate: %s\n", m.StartDate.Format("2006-01-02 15:04:05")))
	sb.WriteString(fmt.Sprintf("  EndDate: %s\n", m.EndDate.Format("2006-01-02 15:04:05")))
	sb.WriteString(fmt.Sprintf("  Volume: %s\n", m.Volume))
	sb.WriteString(fmt.Sprintf("  Volume24hr: %.2f\n", m.Volume24hr))
	sb.WriteString(fmt.Sprintf("  Liquidity: %s\n", m.Liquidity))
	sb.WriteString(fmt.Sprintf("  Spread: %.2f\n", m.Spread))
	sb.WriteString(fmt.Sprintf("  Line: %.2f\n", m.Line))

	if len(m.Events) > 0 {
		sb.WriteString("  Events: [\n")
		for _, e := range m.Events {
			sb.WriteString(fmt.Sprintf("    Event{ID: %s, Title: %s, Slug: %s}\n", e.ID, e.Title, e.Slug))
		}
		sb.WriteString("  ]\n")
	}

	if len(m.Categories) > 0 {
		sb.WriteString("  Categories: [")
		for i, cat := range m.Categories {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(cat.Label)
		}
		sb.WriteString("]\n")
	}

	if len(m.Tags) > 0 {
		sb.WriteString("  Tags: [")
		for i, tag := range m.Tags {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(tag.Label)
		}
		sb.WriteString("]\n")
	}

	sb.WriteString("}")
	return sb.String()
}

func (s Series) String() string {
	var sb strings.Builder
	sb.WriteString("Series{\n")
	sb.WriteString(fmt.Sprintf("  ID: %s\n", s.ID))
	sb.WriteString(fmt.Sprintf("  Ticker: %s\n", s.Ticker))
	sb.WriteString(fmt.Sprintf("  Slug: %s\n", s.Slug))
	sb.WriteString(fmt.Sprintf("  Title: %s\n", s.Title))
	sb.WriteString(fmt.Sprintf("  Subtitle: %s\n", s.Subtitle))
	sb.WriteString(fmt.Sprintf("  SeriesType: %s\n", s.SeriesType))
	sb.WriteString(fmt.Sprintf("  Active: %t\n", s.Active))
	sb.WriteString(fmt.Sprintf("  CommentCount: %d\n", s.CommentCount))

	if len(s.Collections) > 0 {
		sb.WriteString(fmt.Sprintf("  Collections: [%d items]\n", len(s.Collections)))
	}

	if len(s.Categories) > 0 {
		sb.WriteString("  Categories: [")
		for i, cat := range s.Categories {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(cat.Label)
		}
		sb.WriteString("]\n")
	}

	if len(s.Tags) > 0 {
		sb.WriteString("  Tags: [")
		for i, tag := range s.Tags {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(tag.Label)
		}
		sb.WriteString("]\n")
	}

	if len(s.Chats) > 0 {
		sb.WriteString(fmt.Sprintf("  Chats: [%d chats]\n", len(s.Chats)))
	}

	sb.WriteString("}")
	return sb.String()
}

func (e Event) String() string {
	var sb strings.Builder
	sb.WriteString("Event{\n")
	sb.WriteString(fmt.Sprintf("  ID: %s\n", e.ID))
	sb.WriteString(fmt.Sprintf("  Ticker: %s\n", e.Ticker))
	sb.WriteString(fmt.Sprintf("  Slug: %s\n", e.Slug))
	sb.WriteString(fmt.Sprintf("  Title: %s\n", e.Title))
	sb.WriteString(fmt.Sprintf("  Subtitle: %s\n", e.Subtitle))
	sb.WriteString(fmt.Sprintf("  Active: %t\n", e.Active))
	sb.WriteString(fmt.Sprintf("  StartDate: %s\n", e.StartDate.Format("2006-01-02 15:04:05")))
	sb.WriteString(fmt.Sprintf("  EndDate: %s\n", e.EndDate.Format("2006-01-02 15:04:05")))
	sb.WriteString(fmt.Sprintf("  Category: %s\n", e.Category))
	sb.WriteString(fmt.Sprintf("  Subcategory: %s\n", e.Subcategory))
	sb.WriteString(fmt.Sprintf("  Volume: %.2f\n", e.Volume))
	sb.WriteString(fmt.Sprintf("  Volume24hr: %.2f\n", e.Volume24hr))
	sb.WriteString(fmt.Sprintf("  Liquidity: %.2f\n", e.Liquidity))
	sb.WriteString(fmt.Sprintf("  NegRisk: %t\n", e.NegRisk))
	sb.WriteString(fmt.Sprintf("  CommentsEnabled: %t\n", e.CommentsEnabled))
	sb.WriteString(fmt.Sprintf("  SpreadsMainLine: %.2f\n", e.SpreadsMainLine))
	sb.WriteString(fmt.Sprintf("  TotalsMainLine: %.2f\n", e.TotalsMainLine))

	if len(e.SubEvents) > 0 {
		sb.WriteString(fmt.Sprintf("  SubEvents: [%d sub-events]\n", len(e.SubEvents)))
		for i, sub := range e.SubEvents {
			if i < 3 { // Show first 3 only
				sb.WriteString(fmt.Sprintf("    - %s (ID: %s)\n", sub.Title, sub.ID))
			} else if i == 3 {
				sb.WriteString(fmt.Sprintf("    ... and %d more\n", len(e.SubEvents)-3))
				break
			}
		}
	}

	if len(e.Markets) > 0 {
		sb.WriteString(fmt.Sprintf("  Markets: [%d markets]\n", len(e.Markets)))
		for i, m := range e.Markets {
			if i < 3 { // Show first 3 only
				sb.WriteString(fmt.Sprintf("    - %s (ID: %s)\n", m.Question, m.ID))
			} else if i == 3 {
				sb.WriteString(fmt.Sprintf("    ... and %d more\n", len(e.Markets)-3))
				break
			}
		}
	}

	if len(e.Series) > 0 {
		sb.WriteString("  Series: [")
		for i, s := range e.Series {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(s.Title)
		}
		sb.WriteString("]\n")
	}

	if len(e.Categories) > 0 {
		sb.WriteString("  Categories: [")
		for i, cat := range e.Categories {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(cat.Label)
		}
		sb.WriteString("]\n")
	}

	if len(e.Collections) > 0 {
		sb.WriteString(fmt.Sprintf("  Collections: [%d collections]\n", len(e.Collections)))
	}

	if len(e.Tags) > 0 {
		sb.WriteString("  Tags: [")
		for i, tag := range e.Tags {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(tag.Label)
		}
		sb.WriteString("]\n")
	}

	if len(e.Chats) > 0 {
		sb.WriteString(fmt.Sprintf("  Chats: [%d chats]\n", len(e.Chats)))
	}

	sb.WriteString("}")
	return sb.String()
}
