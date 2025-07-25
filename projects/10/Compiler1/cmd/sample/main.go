package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Token struct {
	XMLName xml.Name
	Value   string  `xml:",chardata"`
	Tokens  []Token `xml:",any"`
}

type Tokens struct {
	XMLName xml.Name `xml:"tokens"`
	Tokens  []Token  `xml:",any"`
}

func main() {
	// ネストしたXMLのサンプル
	tokens := Tokens{
		Tokens: []Token{
			{
				XMLName: xml.Name{Local: "keyword"},
				Value:   " class ",
				Tokens: []Token{
					{
						XMLName: xml.Name{Local: "identifier"},
						Value:   " Main ",
						Tokens: []Token{
							{
								XMLName: xml.Name{Local: "symbol"},
								Value:   " { ",
							},
						},
					},
				},
			},
		},
	}

	// XMLを出力
	encoder := xml.NewEncoder(os.Stdout)
	encoder.Indent("", "  ")
	if err := encoder.Encode(tokens); err != nil {
		fmt.Println("Error encoding XML:", err)
	}
}
