package adf

// ADF representation of an Email Address
type Email struct {
	// The actual value that goes into the xml tag
	Value string `xml:",chardata"`
	// This is represented as a number, but it's really a boolean
	// 0 or 1 for if this phone number is the preferred contact
	PreferredContact int `xml:"preferredcontact,attr,omitempty"`
}
