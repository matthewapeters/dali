package dali

import "fmt"

//Span element
type Span struct {
	ID        string
	Text      string
	StyleName string
	Element
}

//String for span
func (s *Span) String() string {
	style := ""
	if s.StyleName != "" {
		style = fmt.Sprintf(` style="%s"`, s.StyleName)
	}
	return fmt.Sprintf(`<span name="%s"%s>%s</span>`, s.ID, style, s.Text)
}

//Bindings returns nil
func (s *Span) Bindings() *Binding { return nil }

//Children returns an empty Elements
func (s *Span) Children() *Elements { return &Elements{slice: []*Element{}} }

//Class for span
func (s *Span) Class() string { return "span" }

//Style for span
func (s *Span) Style() string { return s.StyleName }

//Name returns the name of the Span
func (s *Span) Name() string { return s.ID }
