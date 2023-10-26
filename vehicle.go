package adf

import "encoding/xml"

// ADF/XML standardization vehicle tag
// TODO: several of these string values can be turned into enums
// refer to the specification pdf for more info https://adfxml.info/adf_spec.pdf
type Vehicle struct {
	XMLName  xml.Name `xml:"vehicle"`
	Interest string   `xml:"interest,attr,omitempty"`
	Status   string   `xml:"status,attr,omitempty"`

	ID *ID `xml:"id,omitempty"`

	Year  string `xml:"year"`
	Make  string `xml:"make"`
	Model string `xml:"model"`
	Vin   string `xml:"vin,omitempty"`
	Stock string `xml:"stock,omitempty"`
	Trim  string `xml:"trim,omitempty"`
	// NOTE: According to the spec doors is a text value,
	// but it could also be an int, which may be more useful?
	// need more data to be sure
	Doors        string `xml:"doors,omitempty"`
	BodyStyle    string `xml:"bodystyle,omitempty"`
	Transmission string `xml:"transmission,omitempty"`

	Odometer Odometer `xml:"odometer,omitempty"`

	// Valid values: excellent, good, fair, poor, and unknown
	Condition string `xml:"condition,omitempty"`

	ColorCombination []struct {
		InteriorColor string `xml:"interiorcolor,omitempty"`
		ExteriorColor string `xml:"exteriorcolor,omitempty"`
		Preference    int    `xml:"preference,omitempty"`
	} `xml:"colorcombination,omitempty"`

	// does this need to be it's own standalone struct?
	ImageTag *struct {
		XMLName xml.Name `xml:"imagetag,omitempty"`
		Value   string   `xml:",chardata"`
		Width   string   `xml:"width,attr,omitempty"`
		Height  string   `xml:"height,attr,omitempty"`
		AltText string   `xml:"alttext,attr,omitempty"`
	} `xml:"imagetag,omitempty"`

	Price         *Price `xml:"price,omitempty"`
	PriceComments string `xml:"pricecomments,omitempty"`

	Options []struct {
		OptionName       string `xml:"optionname"`
		ManufacturerCode string `xml:"manufacturercode"`
		Stock            string `xml:"stock"`
		Weighting        int    `xml:"weighting"`
		Price            Price  `xml:"price"`
	} `xml:"option"`

	Finance *struct {
		Method string `xml:"method,omitempty"`

		Amount *struct {
			// not sure if this is int, float, or string, needs investigation
			Value    string `xml:",chardata"`
			Type     string `xml:"type,attr,omitempty"`
			Limit    string `xml:"limit,attr,omitempty"`
			Currency string `xml:"currency,attr,omitempty"`
		} `xml:"amount,omitempty"`

		Balance *struct {
			// not sure if this is int, float, or string. needs investigation
			Value    int    `xml:",chardata"`
			Type     string `xml:"type,attr,omitempty"`
			Currency string `xml:"currency,attr,omitempty"`
		} `xml:"balance,omitempty"`
	} `xml:"finance,omitempty"`

	Comments string `xml:"comments,omitempty"`
}
