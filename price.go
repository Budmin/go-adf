package adf

import "encoding/xml"

// ADF/XML price tag
// "The price tag is used by vehicle, and option within vehicle"
// TODO: almost all of the attributes could be made into enums at some point
type Price struct {
	XMLName xml.Name `xml:"price,omitempty"`
	// TODO: I can't tell if this is an int of a float, needs further investigation
	Value string `xml:",chardata"`
	// Attribute for the xml tag
	// Valid Values are quote, offer, msrp, invoice, call, appraisal, and asking
	Type string `xml:"type,attr,omitempty"`
	// Attribute for the xml tag
	// Valid Values are ISO 4217 3-letter codes (e.g. "USD")
	Currency string `xml:"currency,attr,omitempty"`
	// Attribute for the xml tag
	// Valid Values are absolute, relative, and percentage
	Delta string `xml:"delta,attr,omitempty"`
	// Attribute for the xml tag
	// Valid Values are msrp and invoice
	RelativeTo string `xml:"relativeto,attr,omitempty"`
	// Attribute for the xml tag
	// Relates to the source of the price (e.g. "Kelley Blue Book")
	Source string `xml:"source,attr,omitempty"`
}
