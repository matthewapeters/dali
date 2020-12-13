package dali

import (
	"errors"
	"fmt"

	"github.com/zserge/lorca"
)

//BoundEvents is a mapping of events and the bound functions that trigger
type BoundEvents map[EventType]*Binding

//Element is an interface for describing an HTML element
type Element interface {
	String() string
	Class() string
	Name() string
	Clickable() bool
	Styles() string
	SetStyle(string)
	Children() *Elements
	Bindings() *BoundEvents
	Value() string
	SetUI(*lorca.UI)
	GetUI() *lorca.UI
}

//Base is the common structure that all Elements have
type Base struct {
	ID          string
	Style       string
	UI          *lorca.UI
	BoundEvents *BoundEvents
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
	(*b.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").innerHTML="%s"`, b.Name(), s))
	return nil
}

// Value returns the value of an item
func (b *Base) Value() string {
	return fmt.Sprintf("%s", (*b.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").value;`, b.Name())))
}

//Set assigns the value to the item
func (b *Base) Set(v string) {
	(*b.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").value="%s";`, b.Name(), v))
}

//Enable sets the base element disabled property to false
func (b *Base) Enable() {
	(*b.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").disabled=false`, b.Name()))
}

//Disable sets the base element disabledproperty to true
func (b *Base) Disable() {
	(*b.GetUI()).Eval(fmt.Sprintf(`document.getElementById("%s").disabled=true`, b.Name()))
}

//Name return the ID of the Base
func (b *Base) Name() string { return b.ID }

//SetUI adds the UI to the Base
func (b *Base) SetUI(ui *lorca.UI) { b.UI = ui }

// GetUI returns the UI from the Base
func (b *Base) GetUI() *lorca.UI { return b.UI }

//Styles returns the object style descriptors
func (b *Base) Styles() string { return b.Style }

//SetStyle will set the style
func (b *Base) SetStyle(s string) { b.Style = s }

//SetBoundFunction provides a clean way to set the bound function on an event
func (b *Base) SetBoundFunction(event EventType, f func()) {
	bnd := (*b.BoundEvents)[event]
	if bnd == nil {
		bnd := &Binding{
			BoundFunction: f,
			FunctionName:  fmt.Sprintf("%s_on_%s", b.Name(), event),
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

//String for Elements
func (els *Elements) String() string {
	if els == nil {
		return ""
	}
	html := ""
	for _, el := range els.slice {
		//fmt.Printf(`%d%b%v%c%s%b`, i, 9, el, 9, *el, 10)
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
	bindings := ""
	for e, b := range *be {
		bindings += fmt.Sprintf(` %s="%s()"`, e, b.FunctionName)

	}
	return bindings
}
