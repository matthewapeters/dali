package dali

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"strings"
)

// LinkType is a constant of either Function or URL
type LinkType string

// AreaShape is a string type
type AreaShape string

// Coordinates is an array of coordinates
type Coordinates []int

const (
	//Default covers the whole area
	Default = AreaShape("default")
	// Rectangle Map
	Rectangle = AreaShape("rect")
	//Circle is a circular Map
	Circle = AreaShape("circle")
	//Polygon is a polygon
	Polygon = AreaShape("poly")
	//Function is a function link type
	Function = LinkType("function")
	//URL is an HREF link type
	URL = LinkType("url")
)

// String representation of coordinates
func (c Coordinates) String() string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint([]int(c))), ","), "[]")
}

//Area is a component of a Map
type Area struct {
	Shape    AreaShape
	Coords   Coordinates
	Alt      string
	URL      string
	LinkType LinkType
}

//Map of regions on an image that are clickable
type Map struct {
	Name  string
	Areas []Area
}

//String html rendering of a map area
func (a Area) String() string {
	link := ""
	if a.LinkType == URL {
		link = fmt.Sprintf(` href="%s"`, a.URL)
	} else {
		link = fmt.Sprintf(` onclick="%s()"`, a.URL)
	}
	return fmt.Sprintf(`<area shape="%s" coords="%s"%s alt="%s">`, a.Shape, a.Coords, link, a.Alt)
}

//String of Map
func (m Map) String() string {
	html := fmt.Sprintf(`<map name="%s">`, m.Name)
	for _, a := range m.Areas {
		html = fmt.Sprintf("%s%s", html, a)
	}

	return fmt.Sprintf(`%s</map>`, html)
}

// Image element
type Image struct {
	Width, Height int
	URL           string
	Alt           string
	AreaMap       Map
	Base
}

// NewImage generates a new Image object
func NewImage(name string, width, height int, url string) *Image {
	return &Image{
		Base:    Base{ID: name},
		Width:   width,
		Height:  height,
		URL:     url,
		AreaMap: Map{Name: fmt.Sprintf(`%s_map`, name), Areas: []Area{}},
	}
}

//String for image
func (i *Image) String() string {
	alt := ""
	style := ""
	if i.Alt != "" {
		alt = fmt.Sprintf(` alt="%s"`, i.Alt)
	}
	if i.Style != "" {
		style = fmt.Sprintf(` style="%s"`, i.Style)
	}
	areamap := ""
	if len(i.AreaMap.Areas) > 0 {
		areamap = fmt.Sprintf(` usemap="#%s_map"`, i.Name())
	}
	img := fmt.Sprintf(`<image id="%s" width="%d" height="%d" src="%s"%s%s%s>`, i.ID, i.Width, i.Height, i.URL, alt, style, areamap)
	if len(i.AreaMap.Areas) > 0 {
		img += fmt.Sprintf("%s", i.AreaMap)
	}
	return img
}

//Bindings returns nil
func (i *Image) Bindings() BoundEvents { return nil }

//Children returns an empty Elements
func (i *Image) Children() *Elements { return &Elements{slice: []*Element{}} }

//Class of image
func (i *Image) Class() string { return "image" }

//Clickable attribute of image - is true if there are mapped areas
func (i *Image) Clickable() bool { return len(i.AreaMap.Areas) > 0 }

//AddMapArea to the image map
func (i *Image) AddMapArea(shape AreaShape, coords Coordinates, alt string, linkType LinkType, link string) error {
	if shape == Circle && len(coords) != 3 {
		return fmt.Errorf("Circles must have 3 coordinates")
	}
	if shape == Rectangle && len(coords) != 4 {
		return fmt.Errorf("Rectangls must have 4 coordinates")
	}
	if shape == Polygon {
		if len(coords) < 6 {
			return fmt.Errorf("Polygons must have at least 6 coordinates")
		}
		if math.Mod(float64(len(coords)), 2) != 0 {
			return fmt.Errorf("Polygons must have an even number of coordinates")
		}
	}
	if shape == Default {
		coords = Coordinates{0, 0, i.Width, i.Height}
	}

	a := Area{Shape: shape, Coords: coords, Alt: alt, URL: link, LinkType: linkType}
	i.AreaMap.Areas = append(i.AreaMap.Areas, a)
	return nil
}

//Load an *image.RGBA into the image element
func (i *Image) Load(img *image.RGBA) error {
	if *i.GetUI() == nil {
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
	(*i.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").src="%s"`, i.Name(), imageDump))

	return nil
}
