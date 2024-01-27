package eod

// Price represents the price at End of Day (EOD).
type Price struct {
	Date          string  `json:"date"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	AdjustedClose float64 `json:"adjusted_close"`
	Volume        int     `json:"volume"`
}
