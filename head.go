package dali

import (
	"fmt"

	"github.com/zserge/lorca"
)

//HeadElement provides the Head
type HeadElement struct {
	Title    string
	Elements *Elements
	Base
}

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

func (scr *ScriptElement) String() string {
	src := ""
	if scr.URL != "" {
		src = fmt.Sprintf(` src="%s"`, scr.URL)
	}
	name := ""
	if scr.ID() != "" {
		name = fmt.Sprintf(` id="%s"`, scr.ID())
	}
	return fmt.Sprintf(`<script type="text/javascript" %s%s>
	<!--
	%s
	//-->
	</script>`, src, name, scr.Text)
}

//Bindings on ScriptElement returns an empty BoundEvents
func (scr *ScriptElement) Bindings() *BoundEvents { return &BoundEvents{} }

//Children returns an empty Elements
func (scr *ScriptElement) Children() *Elements { return &Elements{slice: []*Element{}} }

//Class of script
func (scr *ScriptElement) Class() string { return "script" }

// Style of script
func (scr *ScriptElement) Style() string { return "" }

// GetUI gets the lorca.UI
func (scr *ScriptElement) GetUI() *lorca.UI { return scr.UI }

// SetUI sets the lorca.UI
func (scr *ScriptElement) SetUI(u *lorca.UI) { scr.UI = u }

// Name returns the name of the script element
func (scr *ScriptElement) Name() string { return scr.ElementName }

//SetStyle is noop on Script Element
func (scr *ScriptElement) SetStyle(s string) {}

//Value returns empty string on Script Element
func (scr *ScriptElement) Value() string { return "" }

//TitleElement for createing window titles
type TitleElement struct {
	Text string
	Base
}

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
	if b.Style() != "" {
		style = fmt.Sprintf(` style="%s"`, b.Style())
	}
	return fmt.Sprintf(`<body%s%s>%s</body>`, b.BoundEvents, style, b.Elements)
}

//Children return the Elements
func (b *BodyElement) Children() *Elements { return b.Elements }

//NewBodyElement creates a body element
func NewBodyElement(onLoad string) *BodyElement {
	els := Elements{slice: []*Element{}}
	var bindings BoundEvents
	if onLoad != "" {
		bindings = BoundEvents{LoadEvent: &Binding{FunctionName: "body_on_load"}}
	} else {
		bindings = BoundEvents{}
	}

	return &BodyElement{
		Elements: &els,
		Base: Base{
			ElementID:   "BODY",
			BoundEvents: &bindings,
		},
	}
}
