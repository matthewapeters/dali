package dali

import "fmt"

//Button type
type Button struct {
	ID              string
	ButtonText      string
	StyleExpression string
	Element
	Binding
}

func (b *Button) String() string {
	style := ""
	if b.StyleExpression != "" {
		style = fmt.Sprintf(` style="%s"`, b.StyleExpression)
	}
	return fmt.Sprintf(`<button id="%s" onclick="%s()" %s>%s</button>`, b.Name(), b.Name(), style, b.ButtonText)
}

// Children will return an empty Elements
func (b *Button) Children() *Elements { return &Elements{slice: []*Element{}} }

//Name returns the ID of the button
func (b *Button) Name() string {
	return b.ID
}

//Class set the Class of the button
func (b *Button) Class() string {
	return "button"
}

//Style set the style of the button
func (b *Button) Style() string {
	return b.StyleExpression
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
		ID:         name,
		ButtonText: label,
		Binding:    Binding{FunctionName: funcName},
	}
}
