package dali

import (
	"fmt"

	"github.com/zserge/lorca"
)

// HeaderLevel enums header sizes
type HeaderLevel int

const (
	//H1 header
	H1 = HeaderLevel(1)
	//H2 header
	H2 = HeaderLevel(2)
	//H3 header
	H3 = HeaderLevel(3)
	//H4 header
	H4 = HeaderLevel(4)
)

// BR a break tag
type BR struct {
	StyleName string
	Base
}

// Children will return an empty Elements
func (br *BR) Children() *Elements {
	return &Elements{slice: []*Element{}}
}

//LineBreak generates a BR tag
func LineBreak() *BR {
	return &BR{}
}

func (br *BR) String() string {
	return "<br/>"
}

//Class  of the BR
func (br *BR) Class() string { return "BR" }

//Style of the BR
func (br *BR) Style() string { return br.StyleName }

//Header is a header
type Header struct {
	Level HeaderLevel
	Text  string
	Element
	Base
}

func (h *Header) String() string {
	return fmt.Sprintf(`<H%d %s%s%s>%s</H%d>`, h.Level, h.getName(), h.getID(), h.getStyle(), h.Text, h.Level)
}

//NewHeader produces a new header element
func NewHeader(level HeaderLevel, name, id, text string) *Header {
	return &Header{
		Text: text,
		Base: Base{
			ElementID:   id,
			ElementName: name},
		Level: level,
	}
}

//Value returns the base value
func (h *Header) Value() string { return h.Base.Value() }

//Bindings returns bound functions
func (h *Header) Bindings() *BoundEvents { return &BoundEvents{} }

//Children will return an empty Elements
func (h *Header) Children() *Elements {
	return &Elements{slice: []*Element{}}
}

//Class returns the header class
func (h *Header) Class() string {
	return fmt.Sprintf(`H%d`, h.Level)
}

//Style returns the style of the Header
func (h *Header) Style() string {
	return h.getStyle()
}

// GetUI will return the lorca.UI
func (h *Header) GetUI() *lorca.UI {
	return h.Base.GetUI()
}

// SetUI will set the lorca.UI
func (h *Header) SetUI(u *lorca.UI) {
	h.Base.SetUI(u)
}

//Name returns the base element name
func (h *Header) Name() string { return h.Base.Name() }

//SetStyle sets the base style
func (h *Header) SetStyle(s string) { h.Base.ElementStyle = s }
