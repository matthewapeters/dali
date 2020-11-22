package dali

import "fmt"

// Div is a page within a Window
type Div struct {
	ID        string
	StyleName string
	Elements  *Elements
	Element
	Binding
}

// Bindings returns the binding
func (p *Div) Bindings() *Binding { return &p.Binding }

//Children returns the Elements
func (p *Div) Children() *Elements { return p.Elements }

//String for Div
func (p *Div) String() string {
	style := ""
	if p.StyleName != "" {
		style = fmt.Sprintf(` style="%s"`, p.StyleName)
	}

	return fmt.Sprintf(`<div id="%s"%s>%s</div>`, p.Name(), style, p.Elements)
}

// NewDiv generates a new Div
func NewDiv(name string) *Div {
	els := Elements{slice: []*Element{}}
	return &Div{
		ID:       name,
		Elements: &els,
	}
}

//Name returns the name of the Div
func (p *Div) Name() string {
	return p.ID
}

//Class of a div is DIV
func (p *Div) Class() string {
	return "DIV"
}
