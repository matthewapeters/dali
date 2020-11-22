package dali

import "fmt"

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
	Element
}

// Children will return an empty Elements
func (br *BR) Children() *Elements {
	return &Elements{slice: []*Element{}}
}

// Bindings returns nil
func (br *BR) Bindings() *Binding { return nil }

//LineBreak generates a BR tag
func LineBreak() *BR {
	return &BR{}
}

func (br *BR) String() string {
	return "<br/>"
}

//Name of the BR
func (br *BR) Name() string { return "" }

//Class  of the BR
func (br *BR) Class() string { return "BR" }

//Style of the BR
func (br *BR) Style() string { return br.StyleName }

//Header is a header
type Header struct {
	StyleName string
	ID        string
	Level     HeaderLevel
	Text      string
	Element
}

func (h *Header) String() string {
	style := ""
	name := ""
	if h.StyleName != "" {
		style = fmt.Sprintf(` style:"%s"`, h.StyleName)
	}
	if h.ID != "" {
		name = fmt.Sprintf(` id="%s"`, h.Name())
	}
	return fmt.Sprintf(`<H%d %s%s>%s</H%d>`, h.Level, name, style, h.Text, h.Level)
}

//NewHeader produces a new header element
func NewHeader(level HeaderLevel, name, text string) *Header {
	return &Header{
		Text:  text,
		ID:    name,
		Level: level,
	}
}

//Bindings return nil
func (h *Header) Bindings() *Binding { return nil }

//Children will return an empty Elements
func (h *Header) Children() *Elements {
	return &Elements{slice: []*Element{}}
}

//Class returns the header class
func (h *Header) Class() string {
	return fmt.Sprintf(`H%d`, h.Level)
}

//Name of header
func (h *Header) Name() string { return h.ID }

//Style returns the style of the Header
func (h *Header) Style() string {
	return fmt.Sprintf(` style="%s"`, h.StyleName)
}
