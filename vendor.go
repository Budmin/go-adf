package adf

import "encoding/xml"

// ADF/XML standardization vendor tag
type Vendor struct {
	XMLName xml.Name `xml:"vendor"`

	// ID of the vendor, based on ADF ID specification
	ID         *ID    `xml:"id,omitempty"`
	VendorName string `xml:"vendorname,omitempty"`
	Url        string `xml:"url,omitempty"`

	// Contact information of the vendor, based on ADF ID specification
	Contact *Contact `xml:"contact,omitempty"`
}
