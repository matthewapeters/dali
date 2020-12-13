package dali

import "fmt"

//Button type
type Button struct {
	ButtonText string
	Style      string
	Base
}

func (b *Button) String() string {
	style := ""
	if b.Style != "" {
		style = fmt.Sprintf(` style="%s"`, b.Style)
	}
	return fmt.Sprintf(`<button id="%s"%s%s>%s</button>`, b.Name(), b.BoundEvents, style, b.ButtonText)
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

//NewButton creates a new button
func NewButton(label, name, funcName string) *Button {

	return &Button{
		Base: Base{ID: name,
			BoundEvents: &BoundEvents{
				ClickEvent: &Binding{FunctionName: funcName},
			},
		},
		ButtonText: label,
	}
}
