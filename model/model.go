package model

import (
	"fmt"
	"time"
)

// Cryptoresponse is exported, it models the data we receive.
type Cryptoresponse struct {
	ID          string    `json:"id"`
	FullName    string    `json:"fullName"`
	Ask         string    `json:"ask"`
	Bid         string    `json:"bid"`
	Last        string    `json:"last"`
	Open        string    `json:"open"`
	Low         string    `json:"low"`
	High        string    `json:"high"`
	Volume      string    `json:"volume"`
	VolumeQuote string    `json:"volume_quote"`
	FeeCurrency string    `json:"feeCurrency"`
	Timestamp   time.Time `json:"timestamp"`
}

// TextOutput is exported,it formats the data to plain text.
func (c Cryptoresponse) TextOutput() string {
	p := fmt.Sprintf(
		"ID: %s\nFullname : %s\nAsk: %s\nBid: %s\nLast: %s\nOpen: %s\nLow : %s\nHigh: %s\nFeeCurrency: %s\n",
		c.ID[:3], c.FullName, c.Ask, c.Bid, c.Last, c.Open, c.Low, c.High, c.FeeCurrency)
	return p
}
