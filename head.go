package dali

import "fmt"

//HeadElement provides the Head
type HeadElement struct {
	Title    string
	Elements Elements
	Element
}

//String for Head
func (h *HeadElement) String() string {
	return fmt.Sprintf(`<head>%s</head>`, h.Elements)
}

func (h *HeadElement) Class() string { return "Head" }

func (h *HeadElement) Style() string { return "" }

func (h *HeadElement) Name() string { return "" }

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

func (scr *ScriptElement) Class() string { return "script" }

func (scr *ScriptElement) Name() string { return scr.ID }

func (scr *ScriptElement) Style() string { return "" }

type TitleElement struct {
	Text string
	Element
}

func (t *TitleElement) String() string {
	return fmt.Sprintf(`<title>%s</title>`, t.Text)
}
func (t *TitleElement) Class() string { return "title" }

func (t *TitleElement) Name() string { return "" }

func (t *TitleElement) Style() string { return "" }
