package main

/**
* Copyright (c)2020, Matthew A Peters
 */

import (
	"fmt"
	"net/url"
	"os"

	"github.com/zserge/lorca"
)

// Pane is a page within a Window
type Pane struct {
	Class string
	Name  string
	Style string
}

// Window is the main application window
type Window struct {
	Pages map[string]*Pane
	ui    lorca.UI
}

//Render produces the HTML rendering for the Pane
func (p Pane) Render() string {
	return fmt.Sprintf(`<div id="%s" ></div>`, p.Name)
}

// NewWindow creates a new Window
func NewWindow(width, height int, profileDir string, args ...string) Window {

	minimalTemplate := `<html><body></body></html>`

	newui, err := lorca.New("data:text/html,"+url.PathEscape(minimalTemplate), profileDir, width, height, args...)
	if err != nil {
		os.Exit(1)
	}
	return Window{
		ui: newui,
	}
}
