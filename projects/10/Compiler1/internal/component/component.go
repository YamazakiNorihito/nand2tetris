package component

import "encoding/xml"

type Component struct {
	XMLName  xml.Name     `xml:""`
	Value    string       `xml:",chardata"`
	Children []*Component `xml:",any"`
}

func New(name string, value string) *Component {
	return &Component{
		XMLName:  xml.Name{Local: name},
		Value:    value,
		Children: []*Component{},
	}
}
