package dali

import "fmt"

//Span element
type Span struct {
	ID        string
	Text      string
	ClassName string
	StyleName string
	Element
}

//String for span
func (s Span) String() string {
	class := ""
	style := ""
	if s.ClassName != "" {
		class = fmt.Sprintf(` class="%s"`, s.ClassName)
	}
	if s.StyleName != "" {
		style = fmt.Sprintf(` style="%s"`, s.StyleName)
	}
	return fmt.Sprintf(`<span name="%s"%s%s>%s</span>`, s.ID, class, style, s.Text)
}

//Class for span
func (s Span) Class() string { return s.ClassName }

//Style for span
func (s Span) Style() string { return s.StyleName }

//Name returns the name of the Span
func (s Span) Name() string { return s.ID }
