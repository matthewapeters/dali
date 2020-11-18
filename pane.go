package dali

import "fmt"

// Pane is a page within a Window
type Pane struct {
	ID        string
	StyleName string
	Elements  *Elements
	Element
}

//String for Pane
func (p *Pane) String() string {
	style := ""
	if p.StyleName != "" {
		style = fmt.Sprintf(` style="%s"`, p.StyleName)
	}

	return fmt.Sprintf(`<div id="%s"%s>%s</div>`, p.Name(), style, p.Elements)
}

// NewPane generates a new Pane
func NewPane(name string) *Pane {
	els := Elements{slice: []*Element{}}
	return &Pane{
		ID:       name,
		Elements: &els,
	}
}

//Name returns the name of the Pane
func (p *Pane) Name() string {
	return p.ID
}

//Class of a pane is DIV
func (p *Pane) Class() string {
	return "DIV"
}
