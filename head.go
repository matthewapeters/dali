package dali

import "fmt"

//HeadElement provides the Head
type HeadElement struct {
	Title    string
	Elements *Elements
	Element
}

//String for Head
func (h *HeadElement) String() string {
	return fmt.Sprintf(`<head>%s</head>`, h.Elements)
}

//Class of the HeadElement
func (h *HeadElement) Class() string { return "head" }

// Style - not applicable
func (h *HeadElement) Style() string { return "" }

// Name of head - not applicable
func (h *HeadElement) Name() string { return "" }

// NewHeadElement to create a new Head Element
func NewHeadElement() *HeadElement {
	els := Elements{slice: []*Element{}}

	return &HeadElement{
		Elements: &els,
	}
}

// ScriptElement is for scripts
type ScriptElement struct {
	URL  string
	Text string
	ID   string
	Element
}

func (scr *ScriptElement) String() string {
	src := ""
	if scr.URL != "" {
		src = fmt.Sprintf(` src="%s"`, scr.URL)
	}
	name := ""
	if scr.ID != "" {
		name = fmt.Sprintf(` id="%s"`, scr.Name())
	}
	return fmt.Sprintf(`<script%s%s>%s</script>`, src, name, scr.Text)
}

//Class of script
func (scr *ScriptElement) Class() string { return "script" }

// Name of script
func (scr *ScriptElement) Name() string { return scr.ID }

// Style of script
func (scr *ScriptElement) Style() string { return "" }

//TitleElement for createing window titles
type TitleElement struct {
	Text string
	Element
}

//String stringer for Title
func (t *TitleElement) String() string {
	return fmt.Sprintf(`<title>%s</title>`, t.Text)
}

//Class for title
func (t *TitleElement) Class() string { return "title" }

// Name for title
func (t *TitleElement) Name() string { return "" }

//Style for title
func (t *TitleElement) Style() string { return "" }

//BodyElement for holding the body of the page
type BodyElement struct {
	Elements *Elements
	Element
}

func (b *BodyElement) String() string {
	return fmt.Sprintf(`<body>%s</body>`, b.Elements)
}

//NewBodyElement creates a body element
func NewBodyElement() *BodyElement {
	els := Elements{slice: []*Element{}}
	return &BodyElement{Elements: &els}
}
