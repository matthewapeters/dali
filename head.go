package dali

import "fmt"

//HeadElement provides the Head
type HeadElement struct {
	Title    string
	Elements *Elements
	Base
}

// Bindings returns nil
func (h *HeadElement) Bindings() BoundEvents { return nil }

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
	Base
}

// Bindings returns nil
func (scr *ScriptElement) Bindings() BoundEvents { return nil }

func (scr *ScriptElement) String() string {
	src := ""
	if scr.URL != "" {
		src = fmt.Sprintf(` src="%s"`, scr.URL)
	}
	name := ""
	if scr.ID != "" {
		name = fmt.Sprintf(` id="%s"`, scr.Name())
	}
	return fmt.Sprintf(`<script type="text/javascript" %s%s>
	<!--
	%s
	//-->
	</script>`, src, name, scr.Text)
}

//Children returns an empty Elements
func (scr *ScriptElement) Children() *Elements { return &Elements{slice: []*Element{}} }

//Class of script
func (scr *ScriptElement) Class() string { return "script" }

// Style of script
func (scr *ScriptElement) Style() string { return "" }

//TitleElement for createing window titles
type TitleElement struct {
	Text string
	Base
}

//Bindings returns nil
func (t *TitleElement) Bindings() BoundEvents { return nil }

//Children will return an empty Elements
func (t *TitleElement) Children() *Elements { return &Elements{slice: []*Element{}} }

//String stringer for Title
func (t *TitleElement) String() string {
	return fmt.Sprintf(`<title>%s</title>`, t.Text)
}

//Class for title
func (t *TitleElement) Class() string { return "title" }

//Style for title
func (t *TitleElement) Style() string { return "" }

//BodyElement for holding the body of the page
type BodyElement struct {
	Elements *Elements
	Base
}

func (b *BodyElement) String() string {
	style := ""
	if b.Style != "" {
		style = fmt.Sprintf(` style="%s"`, b.Style)
	}
	onLoad := ""
	if b.BoundEvents != nil {
		for e, bnd := range *b.BoundEvents {
			onLoad += fmt.Sprintf(` %s="%s()"`, e, bnd.FunctionName)
		}
	}
	return fmt.Sprintf(`<body%s%s>%s</body>`, onLoad, style, b.Elements)
}

//Children return the Elements
func (b *BodyElement) Children() *Elements { return b.Elements }

// Bindings returns the Binding
func (b *BodyElement) Bindings() BoundEvents { return b.BoundEvents }

//NewBodyElement creates a body element
func NewBodyElement(onLoad string) *BodyElement {
	els := Elements{slice: []*Element{}}
	var bindings BoundEvents
	if onLoad != "" {
		bindings = &map[EventType]*Binding{LoadEvent: &Binding{FunctionName: "body_on_load"}}
	}

	return &BodyElement{
		Elements: &els,
		Base: Base{
			ID:          "BODY",
			BoundEvents: bindings,
		},
	}
}
