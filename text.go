package dali

import "github.com/zserge/lorca"

//TextElement is an element for plain old text - if you want style, use a Span
type TextElement struct {
	text string
	ui   *lorca.UI
}

//Text creates a TextElement
func Text(t string) *TextElement {
	return &TextElement{text: t}
}

//Bindings returns an empty BoundEvents
func (t TextElement) Bindings() *BoundEvents { return &BoundEvents{} }

//Children returns an empty Elements
func (t TextElement) Children() *Elements { return &Elements{slice: []*Element{}} }

//String stringer for TextElement
func (t TextElement) String() string { return t.text }

//Class for TextElement
func (t TextElement) Class() string { return "" }

//Style for Text Element
func (t TextElement) Style() string { return "" }

//Clickable is false on Text
func (t TextElement) Clickable() bool { return false }

//GetUI returns the lorca.UI
func (t TextElement) GetUI() *lorca.UI { return t.ui }

//SetUI sets the lorca.UI
func (t TextElement) SetUI(u *lorca.UI) { t.ui = u }

//Name returns empty string on Text
func (t TextElement) Name() string { return "" }

// ID returns empty string on Text
func (t TextElement) ID() string { return "" }

//SetStyle is a noop function on Text
func (t TextElement) SetStyle(s string) {}

//Value will return the text of Text Element
func (t TextElement) Value() string { return t.text }
