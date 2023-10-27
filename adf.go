package adf

import (
	"encoding/xml"
	"time"
)

// our struct which contains the complete
// array of all Users in the file
type Adf struct {
	XMLName  xml.Name `xml:"adf"`
	Prospect struct {
		Status string `xml:"status,attr,omitempty"`
		ID     ID     `xml:"id"`

		RequestDate time.Time `xml:"requestdate"`

		Vehicles []Vehicle   `xml:"vehicle"`
		Customer Customer  `xml:"customer"`
		Vendor   Vendor    `xml:"vendor"`
		Provider *Provider `xml:"provider,omitempty"`
	} `xml:"prospect"`
}
