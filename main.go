package dali

/**
* Copyright (c)2020, Matthew A Peters
 */

import (
	"fmt"
	"net/url"

	"github.com/zserge/lorca"
)

// Styles is a map of style elements and values
type Styles map[string]string

//String for Styles
func (s Styles) String() string {
	style := ""
	for k, v := range s {
		style = fmt.Sprintf("%s:%s;%s", k, v, style)
	}
	return style
}

//Element is an interface for describing an HTML element
type Element interface {
	String() string
	Class() string
	Style() string
	Name() string
	Clickable() bool
	Styles() Styles
}

//Elements is a slice of Elements
type Elements struct {
	slice []*Element
}

//String for Elements
func (els *Elements) String() string {
	html := ""
	for _, el := range els.slice {
		html = fmt.Sprintf(`%s%s`, html, *el)
	}

	return html
}

// AddElement appends an element to the slice of elements
func (els *Elements) AddElement(e Element) {
	newSlice := append(els.slice, &e)
	els.slice = newSlice
}

// Window is the main application window
type Window struct {
	Width, Height int
	Style         StyleSheet
	html          string
	ui            lorca.UI
	ProfileDir    string
	Elements      Elements
	Args          []string
}

// StyleSheet references an external stylesheet to load
type StyleSheet struct {
	URL string
}

//String for StyleSheet
func (style StyleSheet) String() string {
	if style.URL == "" {
		return ""
	}
	return fmt.Sprintf(`<link rel="stylesheet" href="%s">`, style.URL)
}

// NewWindow creates a new Window
func NewWindow(width, height int, profileDir string, styleSheet string, args ...string) *Window {

	minimalTemplate := `<html>%s<body>%s</body></html>`

	w := Window{
		Width:      width,
		Height:     height,
		html:       minimalTemplate,
		Style:      StyleSheet{URL: styleSheet},
		ui:         nil,
		Args:       args,
		ProfileDir: profileDir,
	}
	return &w
}

//String for Window
func (w *Window) String() string {
	return fmt.Sprintf(w.html, w.Style, w.Elements)

}

// Start extracts the application HTML and starts the UI
func (w *Window) Start() error {
	newui, err := lorca.New("data:text/html,"+url.PathEscape(fmt.Sprintf("%s", w)), w.ProfileDir, w.Width, w.Height, w.Args...)
	if err != nil {
		return err
	}
	w.ui = newui
	return nil
}

//Close wraps lorca.UI.Close()
func (w *Window) Close() {
	w.ui.Close()
}

//GetUI is a temporary wrapper for retrieving the lorca.UI
func (w *Window) GetUI() lorca.UI {
	return w.ui
}
