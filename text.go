package dali

//Text is an element for plain old text - if you want style, use a Span
type Text struct {
	text string
	Element
}

func (t *Text) String() string { return t.text }

func (t *Text) Class() string { return "" }

func (t *Text) Style() string { return "" }

func (t *Text) Name() string { return "" }
