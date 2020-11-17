package dali

import "fmt"

// Pane is a page within a Window
type Pane struct {
	Class    string
	Name     string
	Style    string
	Elements Elements
}

// Panes is a map of Pane elements
type Panes map[string]Pane

//String for Pane
func (p Pane) String() string {
	class := ""
	style := ""
	if p.Class != "" {
		class = fmt.Sprintf(` class="%s"`, p.Class)
	}
	if p.Style != "" {
		style = fmt.Sprintf(` style="%s"`, p.Style)
	}

	return fmt.Sprintf(`<div id="%s"%s%s>%s</div>`, p.Name, class, style, p.Elements)
}

//String for Panes
func (ps Panes) String() string {
	html := ""
	for _, p := range ps {
		html = fmt.Sprintf(`%s%s`, p, html)
	}
	return html
}

//AddElement adds an element to a Pane
func (p Pane) AddElement(el Element) {
	p.Elements[el.Name()] = el
}

// NewPane generates a new Pane
func NewPane(name string) Pane {
	return Pane{
		Name:     name,
		Elements: Elements(map[string]Element{}),
	}
}
