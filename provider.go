package adf

import "encoding/xml"

// ADF/XML standardization provider tag
type Provider struct {
	XMLName xml.Name `xml:"provider"`

	ID *ID `xml:"id,omitempty"`

	Names  []Name  `xml:"name,omitempty"`
	Emails []Email `xml:"email,omitempty"`
	Phones []Phone `xml:"phone,omitempty"`

	Service string `xml:"service,omitempty"`
	Url     string `xml:"url,omitempty"`

	Contact *Contact `xml:"contact,omitempty"`
}

// TODO: do we need to include a removing method like RemoveName?
func (p *Provider) AddName(newName Name) {
	p.Names = append(p.Names, newName)
}

// TODO: do we need to include a removing method like RemoveEmail?
func (p *Provider) AddEmail(newEmail Email) {
	p.Emails = append(p.Emails, newEmail)
}

// TODO: do we need to include a removing method like RemovePhone?
func (p *Provider) AddPhone(newPhone Phone) {
	p.Phones = append(p.Phones, newPhone)
}
