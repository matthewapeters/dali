package dali

import "fmt"

// Canvas element
type Canvas struct {
	Width, Height int
	ID            string
	StyleName     string
	Element
}

//NewCanvas creates a new Canvas
func NewCanvas(width, height int, name string) *Canvas {
	return &Canvas{
		ID:     name,
		Width:  width,
		Height: height,
	}
}

// Children will return an empty Elements
func (c *Canvas) Children() *Elements { return &Elements{slice: []*Element{}} }

func (c *Canvas) String() string {
	style := ""
	if c.StyleName != "" {
		style = fmt.Sprintf(` style="%s"`, c.Style())
	}
	return fmt.Sprintf(`<canvas id="%s" width="%dpx" height="%dpx"%s></canvas>`, c.ID, c.Width, c.Height, style)
}

// Bindings returns nil
func (c *Canvas) Bindings() *Binding { return nil }

// Name of Canvas
func (c *Canvas) Name() string { return c.ID }

// Class of the canvas
func (c *Canvas) Class() string { return "canvas" }

// Clickable is false on Canvas
func (c *Canvas) Clickable() bool { return false }

//Style of the Canvas
func (c *Canvas) Style() string { return c.StyleName }
