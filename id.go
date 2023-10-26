package adf

import "encoding/xml"

// ADF/XML ID tag
// "The id tag is used by all five categories (prospect, customer, vendor, provider, vehicle)"
type ID struct {
	XMLName xml.Name `xml:"id,omitempty"`
	// The actual value that goes into the xml
	Value string `xml:",chardata"`
	// An attribute for the xml tag
	// valid values are 1-n
	// "The sequence fis a number used to track the history of this piece of data"
	// Example:
	// <id sequence=1 source="Autonation">12345</id>
	// <id sequence=2 source="Cobalt">56789</id>
	Sequence int `xml:"sequence,attr,omitempty"`
	// An attribute for the xml tag
	// "The source is the name of the source that created this id.
	// Different sources may use different ids for the same data as it is passed around."
	// This is a required attribute
	Source string `xml:"source,attr,omitempty"`
}
