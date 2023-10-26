package adf

// ADF representation for a Phone Number
type Phone struct {
	// The actual value that's in the xml tag
	Value string `xml:",chardata"`
	// Attribute for the xml tag
	// Valid values are phone, fax, cellphone, and pager
	Type string `xml:"type,attr,omitempty"`
	// Attribute for the xml tag
	// Valid values are morning, afternoon, evening, nopreference, and day
	Time string `xml:"time,attr,omitempty"`
	// This is represented as a number, but it's really a boolean
	// 0 or 1 for if this phone number is the preferred contact
	PreferredContact int `xml:"preferredcontact,attr,omitempty"`
}
