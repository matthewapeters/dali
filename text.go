package dali

//TextElement is an element for plain old text - if you want style, use a Span
type TextElement struct {
	text string
	Element
}

//Text creates a TextElement
func Text(t string) *TextElement {
	return &TextElement{text: t}
}

func (t *TextElement) String() string { return t.text }

func (t *TextElement) Class() string { return "" }

func (t *TextElement) Style() string { return "" }

func (t *TextElement) Name() string { return "" }
