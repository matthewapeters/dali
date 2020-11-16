package dali

import "fmt"

//Button type
type Button struct {
	ButtonText string
	Style      string
	Base
	Binding
}

func (b *Button) String() string {
	style := ""
	if b.Style != "" {
		style = fmt.Sprintf(` style="%s"`, b.Style)
	}
	return fmt.Sprintf(`<button id="%s" onclick="%s()" %s>%s</button>`, b.Name(), b.Binding.FunctionName, style, b.ButtonText)
}

// Children will return an empty Elements
func (b *Button) Children() *Elements { return &Elements{slice: []*Element{}} }

//Class set the Class of the button
func (b *Button) Class() string {
	return "button"
}

// Clickable returns true for buttons
func (b *Button) Clickable() bool {
	return true
}

//OnClick returns the name of the onclick function
func (b *Button) OnClick() string {
	return fmt.Sprintf("on_%s()", b.Name())
}

// BindFunction will bind the button to this go function
func (b *Button) BindFunction(f func()) {
	b.Binding.BoundFunction = f
}

//Bindings returns the bindings for the button
func (b *Button) Bindings() *Binding { return &b.Binding }

//NewButton creates a new button
func NewButton(label, name, funcName string) *Button {
	return &Button{
		Base:       Base{ID: name},
		ButtonText: label,
		Binding:    Binding{FunctionName: funcName},
	}
}
