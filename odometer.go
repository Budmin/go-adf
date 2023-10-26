package adf

import "encoding/xml"

// ADF struct used in the Vehicle tag
type Odometer struct {
	XMLName xml.Name `xml:"odometer"`

	// The numerical value of the odometer
	Value int `xml:",chardata"`
	// An attribute of the xml tag
	// Valid Values include unknown, rolledover, replaced, and original
	Status string `xml:"status,attr,omitempty"`
	// An attribute of the xml tag
	// Valid Values include km and mi
	Units string `xml:"units,attr,omitempty"`
}
