package dali

//Span element
type Span struct {
	Base
}

// NewSpanElement creates a new Span element
func NewSpanElement(name, id, text string) *Span {
	var textElement Element
	textElement = Text(text)
	return &Span{
		Base: Base{
			ElementName:  name,
			ElementID:    id,
			ElementClass: "span",
			Elements: &Elements{slice: []*Element{
				&textElement,
			}},
		},
	}
}
