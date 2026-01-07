// gammago/types.go

package gammago

import "time"

type Status string

const (
	ACTIVE Status = "ACTIVE"
	CLOSED Status = "CLOSED"
)

type Team struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	League       string `json:"league"`
	Logo         string `json:"logo"`
	Abbreviation string `json:"abbreviation"`
	Alias        string `json:"alias"`
	Color        string `json:"color"`
}

type Sport struct {
	Sport      string `json:"sport"`
	Image      string `json:"image"`
	Resolution string `json:"resolution"`
	Ordering   string `json:"ordering"`
	Tags       string `json:"tags"`
	Series     string `json:"series"`
}

type MarketTypes struct {
	MarketTypes []string `json:"marketTypes"`
}

type Tag struct {
	ID    string `json:"id"`
	Label string `json:"label"`
	Slug  string `json:"slug"`
}

type Category struct {
	ID             string `json:"id"`
	Label          string `json:"label"`
	ParentCategory string `json:"parentCategory"`
	Slug           string `json:"slug"`
}

type Chat struct {
	ID           string `json:"id"`
	ChannelID    string `json:"channelId"`
	ChannelName  string `json:"channelName"`
	ChannelImage string `json:"channelImage"`
	Live         bool   `json:"live"`
}

type Collection struct {
	ID             string `json:"id"`
	Ticker         string `json:"ticker"`
	Slug           string `json:"slug"`
	Title          string `json:"title"`
	Subtitle       string `json:"subtitle"`
	CollectionType string `json:"collectionType"`
	Description    string `json:"description"`
	Tags           string `json:"tags"`
	Image          string `json:"image"`
	Icon           string `json:"icon"`
	HeaderImage    string `json:"headerImage"`
	Active         bool   `json:"active"`
}

type Market struct {
	ID               string     `json:"id"`
	Question         string     `json:"question"`
	ConditionID      string     `json:"conditionId"`
	Slug             string     `json:"slug"`
	EndDate          time.Time  `json:"endDate"`
	Category         string     `json:"category"`
	AMMType          string     `json:"ammType"`
	Liquidity        string     `json:"liquidity"`
	StartDate        time.Time  `json:"startDate"`
	Image            string     `json:"image"`
	Icon             string     `json:"icon"`
	LowerBound       string     `json:"lowerBound"`
	UpperBound       string     `json:"upperBound"`
	Description      string     `json:"description"`
	Outcomes         string     `json:"outcomes"`
	OutcomePrices    string     `json:"outcomePrices"`
	Volume           string     `json:"volume"`
	Active           bool       `json:"active"`
	MarketType       string     `json:"marketType"`
	QuestionID       string     `json:"questionID"`
	Volume24hr       float64    `json:"volume24hr"`
	Volume1wk        float64    `json:"volume1wk"`
	Volume1mo        float64    `json:"volume1mo"`
	Volume1yr        float64    `json:"volume1yr"`
	CLOBTokenIDs     string     `json:"clobTokenIds"`
	TeamAID          string     `json:"teamAID"`
	TeamBID          string     `json:"teamBID"`
	UMABond          string     `json:"umaBond"`
	UMAReward        string     `json:"umaReward"`
	Events           []Event    `json:"events"`
	Categories       []Category `json:"categories"`
	Tags             []Tag      `json:"tags"`
	Spread           float64    `json:"spread"`
	SportsMarketType string     `json:"sportsMarketType"`
	Line             float64    `json:"line"`
	EventStartTime   time.Time  `json:"eventStartTime"`
}

type Series struct {
	ID           string       `json:"id"`
	Ticker       string       `json:"ticker"`
	Slug         string       `json:"slug"`
	Title        string       `json:"title"`
	Subtitle     string       `json:"subtitle"`
	SeriesType   string       `json:"seriesType"`
	Recurrence   string       `json:"recurrence"`
	Description  string       `json:"description"`
	Image        string       `json:"image"`
	Icon         string       `json:"icon"`
	Active       bool         `json:"active"`
	Events       string       `json:"events"`
	Collections  []Collection `json:"collections"`
	Categories   []Category   `json:"categories"`
	Tags         []Tag        `json:"tags"`
	CommentCount int          `json:"commentCount"`
	Chats        []Chat       `json:"chats"`
}

type Event struct {
	ID               string       `json:"id"`
	Ticker           string       `json:"ticker"`
	Slug             string       `json:"slug"`
	Title            string       `json:"title"`
	Subtitle         string       `json:"subtitle"`
	Description      string       `json:"description"`
	ResolutionSource string       `json:"resolutionSource"`
	StartDate        time.Time    `json:"startDate"`
	CreationDate     time.Time    `json:"creationDate"`
	EndDate          time.Time    `json:"endDate"`
	Image            string       `json:"image"`
	Icon             string       `json:"icon"`
	Active           bool         `json:"active"`
	Liquidity        float64      `json:"liquidity"`
	Volume           float64      `json:"volume"`
	SortBy           string       `json:"sortBy"`
	Category         string       `json:"category"`
	Subcategory      string       `json:"subcategory"`
	IsTemplate       bool         `json:"isTemplate"`
	PublishedAt      string       `json:"published_at"`
	CreatedBy        string       `json:"createdBy"`
	UpdatedBy        string       `json:"updatedBy"`
	CreatedAt        time.Time    `json:"createdAt"`
	UpdatedAt        time.Time    `json:"updatedAt"`
	CommentsEnabled  bool         `json:"commentsEnabled"`
	Volume24hr       float64      `json:"volume24hr"`
	Volume1wk        float64      `json:"volume1wk"`
	Volume1mo        float64      `json:"volume1mo"`
	Volume1yr        float64      `json:"volume1yr"`
	FeaturedImage    string       `json:"featuredImage"`
	ParentEvent      string       `json:"parentEvent"`
	NegRisk          bool         `json:"negRisk"`
	NegRiskMarketID  string       `json:"negRiskMarketID"`
	SubEvents        []Event      `json:"subEvents"`
	Markets          []Market     `json:"markets"`
	Series           []Series     `json:"series"`
	Categories       []Category   `json:"categories"`
	Collections      []Collection `json:"collections"`
	Tags             []Tag        `json:"tags"`
	StartTime        time.Time    `json:"startTime"`
	SeriesSlug       string       `json:"seriesSlug"`
	Chats            []Chat       `json:"chats"`
	SpreadsMainLine  float64      `json:"spreadsMainLine"`
	TotalsMainLine   float64      `json:"totalsMainLine"`
}
