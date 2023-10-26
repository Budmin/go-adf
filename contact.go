package adf

import "encoding/xml"

// ADF/XML standardization contact tag
// "The contact structure, which is used in consumer, vendor, and provider"
type Contact struct {
	XMLName xml.Name `xml:"contact,omitempty"`
	// An attribute for the xml tag.
	// This is represented as a number, but it's really a boolean
	// 0 or 1 for if this phone number is the primary contact
	PrimaryContact int `xml:"primarycontact,attr,omitempty"`

	Names []Name `xml:"name,omitempty"`

	Emails []Email `xml:"email,omitempty"`

	Phones []Phone `xml:"phone,omitempty"`

	Addresses []Address `xml:"address,omitempty"`
}

// TODO: do we need to include a removing method like RemoveName?
func (c *Contact) AddName(newName Name) {
	c.Names = append(c.Names, newName)
}

// TODO: do we need to include a removing method like RemoveEmail?
func (c *Contact) AddEmail(newEmail Email) {
	c.Emails = append(c.Emails, newEmail)
}

// TODO: do we need to include a removing method like RemovePhone?
func (c *Contact) AddPhone(newPhone Phone) {
	c.Phones = append(c.Phones, newPhone)
}

// TODO: should there be a RemoveAddress?
func (c *Contact) AddAddress(newAddress Address) {
	c.Addresses = append(c.Addresses, newAddress)
}
