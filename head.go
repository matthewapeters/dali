package dali

import "fmt"

//HeadElement provides the Head
type HeadElement struct {
	Title    string
	Elements *Elements
	Element
}

// Bindings returns nil
func (h *HeadElement) Bindings() *Binding { return nil }

//Children returns the Elements
func (h *HeadElement) Children() *Elements { return h.Elements }

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

// Bindings returns nil
func (scr *ScriptElement) Bindings() *Binding { return nil }

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

//Children returns an empty Elements
func (scr *ScriptElement) Children() *Elements { return &Elements{slice: []*Element{}} }

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

//Bindings returns nil
func (t *TitleElement) Bindings() *Binding { return nil }

//Children will return an empty Elements
func (t *TitleElement) Children() *Elements { return &Elements{slice: []*Element{}} }

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
	Binding *Binding
}

func (b *BodyElement) String() string {
	onLoad := ""
	if b.Binding != nil {
		onLoad = fmt.Sprintf(` onload="%s()"`, b.Binding.FunctionName)
	}
	return fmt.Sprintf(`<body%s>%s</body>`, onLoad, b.Elements)
}

//Children return the Elements
func (b *BodyElement) Children() *Elements { return b.Elements }

// Bindings returns the Binding
func (b *BodyElement) Bindings() *Binding { return b.Binding }

//NewBodyElement creates a body element
func NewBodyElement(onLoad string) *BodyElement {
	els := Elements{slice: []*Element{}}
	var binding *Binding
	if onLoad != "" {
		binding = &Binding{FunctionName: onLoad}
	}

	return &BodyElement{
		Elements: &els,
		Binding:  binding,
	}
}
