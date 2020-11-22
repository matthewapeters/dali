package dali

import (
	"fmt"
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
	ID            string
	Width, Height int
	URL           string
	StyleName     string
	Alt           string
	AreaMap       Map
	Element
}

// NewImage generates a new Image object
func NewImage(name string, width, height int, url string) *Image {
	return &Image{
		ID:      name,
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
	if i.StyleName != "" {
		style = fmt.Sprintf(` style="%s"`, i.StyleName)
	}
	areamap := ""
	if i.Clickable() {
		areamap = fmt.Sprintf(`usemap="#%s_map"`, i.Name())
	}
	img := fmt.Sprintf(`<image name="%s" width="%d" height="%d" src="%s"%s%s%s>`, i.ID, i.Width, i.Height, i.URL, alt, style, areamap)
	if len(i.AreaMap.Areas) > 0 {
		img = fmt.Sprintf(`%s%s`, img, i.AreaMap)
	}
	return img
}

//Bindings returns nil
func (i *Image) Bindings() *Binding { return nil }

//Children returns an empty Elements
func (i *Image) Children() *Elements { return &Elements{slice: []*Element{}} }

//Class of image
func (i *Image) Class() string { return "image" }

//Style of image
func (i *Image) Style() string { return i.StyleName }

//Name of image
func (i *Image) Name() string { return i.ID }

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
