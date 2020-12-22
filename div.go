package dali

// Div is a page within a Window
type Div struct {
	Base
}

//Children returns the Elements
func (p *Div) Children() *Elements { return p.Elements }

// NewDiv generates a new Div
func NewDiv(name, id string) *Div {
	els := Elements{slice: []*Element{}}
	return &Div{
		Base: Base{ElementID: id, ElementName: name, ElementClass: "div", Elements: &els},
	}
}

//Class of a div is DIV
func (p *Div) Class() string {
	return p.ElementClass
}
