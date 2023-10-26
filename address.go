package adf

type Street struct {
	// The actual value of the xml tag
	Value string `xml:",chardata"`
	// An attribute of the xml tag
	// Valid values are 1-5.
	// "Street name. Allowed up to 5 lines"
	Line int `xml:"line,attr"`
}

// ADF representation of an address
type Address struct {
	// An attribute of the xml
	// Valid values are work, home, and delivery
	Type string `xml:"type,attr,omitempty"`

	Streets []Street `xml:"street,omitempty"`

	Apartment string `xml:"apartment,omitempty"`
	City      string `xml:"city,omitempty"`
	// Valid values: "Free-form, 2-char code recommended for N. America"
	RegionCode string `xml:"regioncode,omitempty"`
	PostalCode string `xml:"postalcode,omitempty"`
	// Valid values: ISO 3166 2-letter code
	Country string `xml:"country,omitempty"`
}

// TODO: should there be a RemoveStreet?
func (a *Address) AddStreet(newStreet Street) {
	a.Streets = append(a.Streets, newStreet)
}
