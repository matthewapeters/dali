package dali

import "fmt"

// Canvas element
type Canvas struct {
	Width, Height int
	Base
	Elements *Elements
}

//NewCanvas creates a new Canvas
func NewCanvas(name string, width, height int) *Canvas {
	return &Canvas{
		Width:  width,
		Height: height,
		Base: Base{ID: name,
			BoundEvents: &BoundEvents{}},
		Elements: &Elements{},
	}
}

// Children will return an empty Elements
func (c *Canvas) Children() *Elements { return c.Elements }

func (c *Canvas) String() string {
	style := ""
	if c.Base.Style != "" {
		style = fmt.Sprintf(` style="%s"`, c.Style)
	}
	return fmt.Sprintf(`<canvas id="%s" width="%dpx" height="%dpx"%s>%s</canvas>`, c.ID, c.Width, c.Height, style, c.Elements)
}

// Class of the canvas
func (c *Canvas) Class() string { return "canvas" }
