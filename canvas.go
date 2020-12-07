package dali

import "fmt"

// Canvas element
type Canvas struct {
	Width, Height int
	Base
}

//NewCanvas creates a new Canvas
func NewCanvas(width, height int, name string) *Canvas {
	return &Canvas{
		Width:  width,
		Height: height,
		Base:   Base{ID: name},
	}
}

// Children will return an empty Elements
func (c *Canvas) Children() *Elements { return &Elements{slice: []*Element{}} }

func (c *Canvas) String() string {
	style := ""
	if c.Base.Style != "" {
		style = fmt.Sprintf(` style="%s"`, c.Style)
	}
	return fmt.Sprintf(`<canvas id="%s" width="%dpx" height="%dpx"%s></canvas>`, c.ID, c.Width, c.Height, style)
}

// Class of the canvas
func (c *Canvas) Class() string { return "canvas" }

// Clickable is false on Canvas
func (c *Canvas) Clickable() bool { return false }
