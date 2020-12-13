package dali

import "fmt"

// Div is a page within a Window
type Div struct {
	Base
	Elements *Elements
}

//Children returns the Elements
func (p *Div) Children() *Elements { return p.Elements }

//String for Div
func (p *Div) String() string {
	style := ""
	if p.Base.Style != "" {
		style = fmt.Sprintf(` style="%s"`, p.Base.Style)
	}

	return fmt.Sprintf(`<div id="%s"%s>%s</div>`, p.Name(), style, p.Elements)
}

// NewDiv generates a new Div
func NewDiv(name string) *Div {
	els := Elements{slice: []*Element{}}
	return &Div{
		Base:     Base{ID: name},
		Elements: &els,
	}
}

//Class of a div is DIV
func (p *Div) Class() string {
	return "DIV"
}
