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
	Level     HeaderLevel
	Text      string
	Element
}

func (h *Header) String() string {
	style := ""
	if h.StyleName != "" {
		style = fmt.Sprintf(` style:"%s"`, h.StyleName)
	}
	return fmt.Sprintf(`<H%d %s>%s</H%d>`, h.Level, style, h.Text, h.Level)
}

//NewHeader produces a new header element
func NewHeader(level HeaderLevel, text string) *Header {
	return &Header{
		Text:  text,
		Level: level,
	}
}

//Class returns the header class
func (h *Header) Class() string {
	return fmt.Sprintf(`H%d`, h.Level)
}

//Name of header - nothing
func (h *Header) Name() string { return "" }

//Style returns the style of the Header
func (h *Header) Style() string {
	return fmt.Sprintf(` style="%s"`, h.StyleName)
}
