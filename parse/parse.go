package parse

import "encoding/xml"

type Results struct {
	Shops []Shop `xml:"shop"`
}

type Shop struct {
	ID        string `xml:"id"`
	Name      string `xml:"name"`
	LogoImage string `xml:"logo_image"`
	Urls      Urls   `xml:"urls"`
}
type Urls struct {
	Pc string `xml:"pc"`
}

func Parse(xmlBytes []byte) ([]Shop, error) {
	var rs Results
	if err := xml.Unmarshal(xmlBytes, &rs); err != nil {
		return nil, err
	}
	return rs.Shops, nil
}
