package adf

// The common name structure in several ADF fields
type Name struct {
	// The actual value of the XML tag
	Value string `xml:",chardata"`
	// An attribute pertaining to the part of a persons name this tag references.
	// Valid Values are first, last, middle, suffix, and full
	Part string `xml:"part,attr,omitempty"`
	// An attribute pertaining to the type of contact
	// Valid Values are individual and business
	Type string `xml:"type,attr,omitempty"`
}
