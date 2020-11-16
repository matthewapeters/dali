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

func (s Styles) toString() string {
	style := ""
	for k, v := range s {
		style = fmt.Sprintf("%s:%s;%s", k, v, style)
	}
	return style
}

// Pane is a page within a Window
type Pane struct {
	Class string
	Name  string
	Style Styles
}

// Window is the main application window
type Window struct {
	Pages map[string]*Pane
	ui    lorca.UI
}

//Render produces the HTML rendering for the Pane
func (p Pane) Render() string {
	return fmt.Sprintf(`<div id="%s" styles="%s"></div>`, p.Name, p.Style)
}

// NewWindow creates a new Window
func NewWindow(width, height int, profileDir string, args ...string) (Window, error) {

	minimalTemplate := `<html><body></body></html>`

	newui, err := lorca.New("data:text/html,"+url.PathEscape(minimalTemplate), profileDir, width, height, args...)
	if err != nil {
		return Window{}, err
	}
	w := Window{
		ui: newui,
	}
	return w, nil
}
