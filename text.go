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

//String stringer for TextElement
func (t *TextElement) String() string { return t.text }

//Class for TextElement
func (t *TextElement) Class() string { return "" }

//Style for Text Element
func (t *TextElement) Style() string { return "" }

//Name for Text Element
func (t *TextElement) Name() string { return "" }
