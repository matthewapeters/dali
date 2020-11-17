package dali

import "fmt"

//Button type
type Button struct {
	ID              string
	ClassName       string
	ButtonText      string
	StyleExpression string
	Element
}

func (b Button) String() string {
	style := ""
	class := ""
	if b.StyleExpression != "" {
		style = fmt.Sprintf(` style="%s"`, b.StyleExpression)
	}
	if b.Class() != "" {
		class = fmt.Sprintf(` class="%s"`, b.Class())
	}
	return fmt.Sprintf(`<button id="%s" onclick="do_%s()" %s%s>%s</button>`, b.Name(), b.Name(), class, style, b.ButtonText)
}

//Name returns the ID of the button
func (b Button) Name() string {
	return b.ID
}

//Class set the Class of the button
func (b Button) Class() string {
	return b.ClassName
}

//Style set the style of the button
func (b Button) Style() string {
	return b.StyleExpression
}

// Clickable returns true for buttons
func (b Button) Clickable() bool {
	return true
}

//OnClick returns the name of the onclick function
func (b Button) OnClick() string {
	return fmt.Sprintf("on_%s()", b.Name())
}
