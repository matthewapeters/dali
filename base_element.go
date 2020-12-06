package dali

import (
	"errors"
	"fmt"

	"github.com/zserge/lorca"
)

//Element is an interface for describing an HTML element
type Element interface {
	String() string
	Class() string
	Name() string
	Clickable() bool
	Styles() Styles
	Children() *Elements
	Bindings() *Binding
	Value() string
	SetUI(*lorca.UI)
	GetUI() *lorca.UI
}

//Base is the common structure that all Elements have
type Base struct {
	ID    string
	Style string
	UI    *lorca.UI
	Binding
	Element
}

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
		html += fmt.Sprintf(`%s`, *el)
	}
	return html
}

// AddElement appends an element to the slice of elements
func (els *Elements) AddElement(e Element) {
	newSlice := append(els.slice, &e)
	els.slice = newSlice
}
