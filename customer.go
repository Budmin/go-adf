package adf

import (
	"encoding/xml"
	"time"
)

// ADF/XML standardization customer tag
type Customer struct {
	XMLName xml.Name `xml:"customer"`

	ID ID `xml:"id,omitempty"`

	Contact Contact `xml:"contact,omitempty"`

	TimeFrame *struct {
		Description string `xml:"description,omitempty"`

		// Earliest date customer is interested in.
		// If timeframe tag is present, it is required to specify earliestdate and/or latestdate
		// Valid value is a ISO 8601 formatted datetime
		EarliestDate *time.Time `xml:"earliestdate,omitempty"`
		// Latest date customer is interested in.
		// If timeframe tag is present, it is required to specify earliestdate and/or latestdate
		// Valid valus is a ISO 8601 formatted datetime
		LatestDate *time.Time `xml:"latestdate,omitempty"`
	} `xml:"timeframe,omitempty"`

	Comments string `xml:"comments,omitempty"`
}
