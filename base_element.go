package dali

import (
	"errors"
	"fmt"
	"strings"

	"github.com/zserge/lorca"
)

// ElementID is the id of an element
type ElementID string

// ElementName is the name of an element
type ElementName string

// ElementClass is the class of an element
type ElementClass string

//BoundEvents is a mapping of events and the bound functions that trigger
type BoundEvents map[EventType]*Binding

//Element is an interface for describing an HTML element
type Element interface {
	String() string
	Class() string
	Name() string
	Clickable() bool
	Style() string
	SetStyle(string)
	Children() *Elements
	Bindings() *BoundEvents
	Value() string
	SetUI(*lorca.UI)
	GetUI() *lorca.UI
}

//Base is the common structure that all Elements have
type Base struct {
	ElementID    string
	ElementName  string
	ElementClass string
	ElementStyle *StyleSheet
	UI           *lorca.UI
	BoundEvents  *BoundEvents
	Elements     *Elements
	Element
}

//BindFunction allows you to bind a function to an event
func (b *Base) BindFunction(e EventType, functionName string, boundFunction func()) {
	bnd := Binding{FunctionName: functionName, BoundFunction: boundFunction}
	(*b.BoundEvents)[e] = &bnd
}

//Bindings returns the map of events to bound functions
func (b *Base) Bindings() *BoundEvents { return b.BoundEvents }

//SetText replaces the inner text of the element after the Window has been started
func (b *Base) SetText(s string) error {
	if b.GetUI() == nil {
		return errors.New("Window not started yet")
	}
	(*b.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").innerHTML="%s"`, b.ID(), s))
	return nil
}

func (b *Base) getName() string {
	if b.ElementName == "" {
		return ""
	}
	return fmt.Sprintf(` name="%s"`, b.ElementName)
}
func (b *Base) getID() string {
	if b.ElementID == "" {
		return ""
	}
	return fmt.Sprintf(` id="%s"`, b.ElementID)
}

func (b *Base) getStyle() string {
	if b.ElementStyle == nil {
		return ""
	}
	return fmt.Sprintf(` style="%s"`, b.ElementStyle)
}

func (b *Base) String() string {
	return fmt.Sprintf(
		`<%s%s%s%s>%s</%s>`,
		b.Class(), b.getName(), b.getID(), b.getStyle(), b.Elements, b.Class())
}

// Value returns the value of an item
func (b *Base) Value() string {
	return fmt.Sprintf("%s", (*b.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").value;`, b.ID())))
}

//Set assigns the value to the item
func (b *Base) Set(v string) {
	(*b.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").value="%s";`, b.ID(), v))
}

//Enable sets the base element disabled property to false
func (b *Base) Enable() {
	(*b.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").disabled=false`, b.ID()))
}

//Disable sets the base element disabledproperty to true
func (b *Base) Disable() {
	(*b.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").disabled=true`, b.ID()))
}

//ID return the ID() of the Base
func (b *Base) ID() string { return b.ElementID }

//Name returns the name of the element
func (b *Base) Name() string { return b.ElementName }

//Class returns the class of the element
func (b *Base) Class() string { return b.ElementClass }

//SetUI adds the UI to the Base
func (b *Base) SetUI(ui *lorca.UI) { b.UI = ui }

// GetUI returns the UI from the Base
func (b *Base) GetUI() *lorca.UI { return b.UI }

//Style returns the object style descriptors
func (b *Base) Style() string { return fmt.Sprintf("%s", b.ElementStyle) }

//SetStyle will set the style
func (b *Base) SetStyle(s string) {
	if b.ElementStyle == nil {
		b.ElementStyle = NewStyleSheet()
	}
	for _, kvp := range strings.Split(s, ";") {
		if strings.Index(kvp, ":") > 0 {
			k := strings.Split(kvp, ":")[0]
			v := strings.Split(kvp, ":")[1]
			b.ElementStyle.AddProperty(CSS(k), v)
		}
	}
}

// SetStyleProperty sets the indicated stylesheet property to the provided value
func (b *Base) SetStyleProperty(p CSS, value string) {
	if b.ElementStyle == nil {
		b.ElementStyle = NewStyleSheet()
	}
	b.ElementStyle.AddProperty(p, value)
}

//SetBoundFunction provides a clean way to set the bound function on an event
func (b *Base) SetBoundFunction(event EventType, f func()) {
	bnd := (*b.BoundEvents)[event]
	if bnd == nil {
		bnd := &Binding{
			BoundFunction: f,
			FunctionName:  fmt.Sprintf("%s_on_%s", b.ID(), event),
		}
		(*b.BoundEvents)[event] = bnd
	} else {
		bnd.BoundFunction = f
	}
}

//Elements is a slice of Elements
type Elements struct {
	slice []*Element
}

//Length provides the length of the Elements slice
func (els *Elements) Length() int {
	return len(els.slice)
}

//String for Elements
func (els *Elements) String() string {
	html := ""
	if els == nil {
		return html
	}
	for i, el := range els.slice {
		if el == nil {
			fmt.Printf("Element %d is nil\n", i)
			return ""
		}
		html += fmt.Sprintf(`%s`, *el)
	}
	return html
}

// AddElement appends an element to the slice of elements
func (els *Elements) AddElement(e Element) {
	newSlice := append(els.slice, &e)
	els.slice = newSlice
}

func (be *BoundEvents) String() string {
	if *be == nil {
		return ""
	}
	bindings := ""
	for e, b := range *be {
		bindings += fmt.Sprintf(` %s="%s()"`, e, b.FunctionName)

	}
	return bindings
}

//Children returns a pointer to the child elements
func (b *Base) Children() *Elements { return b.Elements }
