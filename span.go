package dali

import "fmt"

//Span element
type Span struct {
	Text string
	Base
}

//String for span
func (s *Span) String() string {
	style := ""
	if s.Style != "" {
		style = fmt.Sprintf(` style="%s"`, s.Style)
	}
	return fmt.Sprintf(`<span id="%s"%s>%s</span>`, s.ID, style, s.Text)
}

//Bindings returns nil
func (s *Span) Bindings() *map[EventType]*Binding { return nil }

//Children returns an empty Elements
func (s *Span) Children() *Elements { return &Elements{slice: []*Element{}} }

//Class for span
func (s *Span) Class() string { return "span" }
