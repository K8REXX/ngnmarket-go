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
