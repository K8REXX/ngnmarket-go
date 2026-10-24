package model

import "time"

// ForexRates is the object inside "data".
type ForexRates struct {
	Target string      `json:"target"` // the base currency, "NGN"
	Date   *time.Time  `json:"date"`
	Rates  []ForexRate `json:"rates"`
	Meta   Meta        `json:"-"`
}

// ForexRate is one currency's rate against NGN.
type ForexRate struct {
	Currency           string     `json:"currency"`
	Rate               float64    `json:"rate"`
	InverseRate        float64    `json:"inverse_rate"`
	DailyChange        float64    `json:"daily_change"`
	DailyChangePercent float64    `json:"daily_change_percent"`
	LastUpdated        *time.Time `json:"last_updated"`
}

// NGXTradingSession holds the NGX trading hours.
type NGXTradingSession struct {
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
	Timezone  string `json:"timezone"`
}

// NGXMarketStatus is the response from GET /market/status.
type NGXMarketStatus struct {
	IsOpen   bool              `json:"is_open"`
	Status   string            `json:"status"` // "open" or "closed"
	Reason   string            `json:"reason"` // open, weekend, holiday, pre_market, after_hours
	Session  NGXTradingSession `json:"session"`
	Holiday  *NGXHoliday       `json:"holiday"`   // nil unless today is a public holiday
	NextOpen *NGXNextOpen      `json:"next_open"` // nil while the market is open
	ClosesIn *NGXClosesIn      `json:"closes_in"` // nil unless the market is open
	AsOf     time.Time         `json:"as_of"`
	Meta     Meta              `json:"-"` // copy this line from ForexRates
}

type NGXHoliday struct {
	Name string `json:"name"`
	Date string `json:"date"`
}

type NGXNextOpen struct {
	Date     string    `json:"date"`
	Datetime time.Time `json:"datetime"`
	Label    string    `json:"label"`
}

type NGXClosesIn struct {
	Hours        int `json:"hours"`
	Minutes      int `json:"minutes"`
	TotalMinutes int `json:"total_minutes"`
}

// NGXMarketSnapshot is the response from GET /market/snapshot.
type NGXMarketSnapshot struct {
	Date                  string            `json:"date"`
	ASI                   float64           `json:"asi"`
	ASIChange             float64           `json:"asi_change"`
	ASIChangePercent      float64           `json:"asi_change_percent"`
	YTDASIChangePercent   *float64          `json:"ytd_asi_change_percent"`
	Deals                 int64             `json:"deals"`
	Volume                int64             `json:"volume"`
	ValueTraded           float64           `json:"value_traded"`
	TurnoverRate          *float64          `json:"turnover_rate"`
	MarketCap             NGXMarketCap      `json:"market_cap"`
	Breadth               NGXMarketBreadth  `json:"breadth"`
	TotalListedSecurities *int              `json:"total_listed_securities"`
	Session               NGXTradingSession `json:"session"`
	UpdatedAt             time.Time         `json:"updated_at"`
	Meta                  Meta              `json:"-"`
}

type NGXMarketCap struct {
	Equity float64 `json:"equity"`
	Bonds  float64 `json:"bonds"`
	ETFs   float64 `json:"etfs"`
	Total  float64 `json:"total"`
}

type NGXMarketBreadth struct {
	Advancers   int      `json:"advancers"`
	Decliners   int      `json:"decliners"`
	Unchanged   int      `json:"unchanged"`
	Total       int      `json:"total"`
	AdvDecRatio *float64 `json:"adv_dec_ratio"` // nil when decliners is zero
}

// NGXIndex is one entry from GET /indices.
type NGXIndex struct {
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	ChangePct float64 `json:"change_pct"`
}
