package dali

import "fmt"

// Pane is a page within a Window
type Pane struct {
	ID        string
	StyleName string
	Elements  Elements
	Element
}

// Panes is a map of Pane elements
type Panes struct {
	List []*Pane
}

// Add a Pane the list of Panes
func (ps *Panes) Add(p *Pane) {
	ps.List = append(ps.List, p)
}

//String for Pane
func (p *Pane) String() string {
	style := ""
	if p.StyleName != "" {
		style = fmt.Sprintf(` style="%s"`, p.StyleName)
	}

	return fmt.Sprintf(`<div id="%s"%s>%s</div>`, p.Name(), style, p.Elements)
}

//String for Panes
func (ps *Panes) String() string {
	html := ""
	for _, p := range ps.List {
		html = fmt.Sprintf(`%s%s`, html, p)
	}
	return html
}

//AddElement adds an element to a Pane
func (p *Pane) AddElement(el Element) {
	p.Elements = append(p.Elements, el)
}

// NewPane generates a new Pane
func NewPane(name string) *Pane {
	return &Pane{
		ID:       name,
		Elements: Elements([]Element{}),
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
