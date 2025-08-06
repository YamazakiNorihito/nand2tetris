package component

import (
	"encoding/xml"
	"fmt"
)

type Category string

const (
	CLASS      Category = "class"
	SUBROUTINE Category = "subroutine"
	STATIC     Category = "static"
	FIELD      Category = "field"
	ARGMENT    Category = "arg"
	VAR        Category = "var"
)

type Usage string

const (
	DECLARED Usage = "declared"
	USED     Usage = "used"
)

type Component struct {
	Children []*Component `xml:",any"`
	XMLName  xml.Name     `xml:""`
	Value    string       `xml:",chardata"`

	Name string `xml:"name,attr,omitempty"`
	// identifier attributes
	Category Category `xml:"category,attr,omitempty"`
	Index    string   `xml:"index,attr,omitempty"`
	Usage    Usage    `xml:"usage,attr,omitempty"`
}

func New(name string, value string) *Component {
	return &Component{
		XMLName:  xml.Name{Local: name},
		Value:    value,
		Children: []*Component{},
	}
}

// NewVariableComponent creates a Component for variable nodes (STATIC, FIELD, ARGMENT, VAR)
func NewVariableComponent(name string, value string, category Category, index int, usage Usage) *Component {
	if index < 0 {
		panic("index cannot be negative")
	}
	if category != STATIC && category != FIELD && category != ARGMENT && category != VAR {
		panic("NewVariableComponent: category must be STATIC, FIELD, ARGMENT, or VAR")
	}
	return &Component{
		XMLName: xml.Name{Local: name},
		Value:   value,

		Name:     value,
		Category: category,
		Index:    fmt.Sprint(index),
		Usage:    usage,
		Children: []*Component{},
	}
}

func NewClassComponent(name string, value string) *Component {
	return &Component{
		XMLName:  xml.Name{Local: name},
		Value:    value,
		Name:     value,
		Category: CLASS,
		Children: []*Component{},
	}
}

func NewSubroutineComponent(name string, value string) *Component {
	return &Component{
		XMLName:  xml.Name{Local: name},
		Value:    value,
		Name:     value,
		Category: SUBROUTINE,
		Children: []*Component{},
	}
}
