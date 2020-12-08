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
	bindings := ""
	if b.Style != "" {
		style = fmt.Sprintf(` style="%s"`, b.Style)
	}
	if b.BoundEvents != nil {
		for e, bnd := range *b.BoundEvents {
			bindings += fmt.Sprintf(` %s="%s()"`, e, bnd.FunctionName)
		}
	}
	return fmt.Sprintf(`<button id="%s" %s%s>%s</button>`, b.Name(), bindings, style, b.ButtonText)
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
			BoundEvents: &map[EventType]*Binding{
				ClickEvent: &Binding{FunctionName: funcName},
			},
		},
		ButtonText: label,
	}
}
