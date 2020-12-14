package dali

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/png"
	"log"
)

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

//DrawImage an *image.RGBA into the image element
func (c *Canvas) DrawImage(img *image.RGBA, x, y int) error {
	if *c.GetUI() == nil {
		return errors.New("Window Not Yet Started")
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		fmt.Println("writeImageWithTemplate unable to encode image", err)
		log.Fatalln("unable to encode image.")
	}
	// Encode the image data as Base64 (non-binary) encoding
	image64 := base64.StdEncoding.EncodeToString(buffer.Bytes())

	// this expression tells tells the web browser that the image content is here,
	// and does not need to be downloaded from a web resource
	imageDump := fmt.Sprintf("data:image/png;base64,%s", image64)

	// this JavaScript will the source content of the img tag
	// Send the javascript containing the image and the instruction to modify the image
	scriptlet := fmt.Sprintf(`
	var img=new Image();
	img.onload=function(){ document.getElementById("%s").getContext("2d").drawImage(img,%d,%d); };
	img.src = "%s";
	`,
		c.Name(), x, y, imageDump)

	(*c.GetUI()).Eval(scriptlet)
	return nil
}
