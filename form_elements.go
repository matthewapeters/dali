package dali

import "fmt"

//InputType identifies input element types
type InputType string

// InputEventType identifies the binding event type
type InputEventType EventType

const (
	//InputTypes

	//ButtonInput        = InputType("button")
	ButtonInput = InputType("button")
	//CheckboxInput      = InputType("checkbox")
	CheckboxInput = InputType("checkbox")
	//ColorInput         = InputType("color")
	ColorInput = InputType("color")
	//DateInput          = InputType("date")
	DateInput = InputType("date")
	//DatetimeLocalInput = InputType("datetime-local")
	DatetimeLocalInput = InputType("datetime-local")
	//EmailInput         = InputType("email")
	EmailInput = InputType("email")
	//FileInput          = InputType("file")
	FileInput = InputType("file")
	//HiddenInput        = InputType("hidden")
	HiddenInput = InputType("hidden")
	//ImageInput         = InputType("image")
	ImageInput = InputType("image")
	//MonthInput         = InputType("month")
	MonthInput = InputType("month")
	//NumberInput        = InputType("number")
	NumberInput = InputType("number")
	//PasswordInput      = InputType("password")
	PasswordInput = InputType("password")
	//RadioInput         = InputType("radio")
	RadioInput = InputType("radio")
	//RangeInput         = InputType("range")
	RangeInput = InputType("range")
	//ResetInput         = InputType("reset")
	ResetInput = InputType("reset")
	//SearchInput        = InputType("search")
	SearchInput = InputType("search")
	//SubmitInput        = InputType("submit")
	SubmitInput = InputType("submit")
	//TelInput           = InputType("tel")
	TelInput = InputType("tel")
	//TextInput          = InputType("text")
	TextInput = InputType("text")
	//TimeInput          = InputType("time")
	TimeInput = InputType("time")
	//URLInput           = InputType("url")
	URLInput = InputType("url")
	//WeekInput          = InputType("week")
	WeekInput = InputType("week")
)

// InputElement is for inputting values
type InputElement struct {
	InputType
	Text string
	Base
	InputEventType
}

func (tf *InputElement) String() string {
	style := ""
	defaultValue := ""
	if tf.Style != "" {
		style = fmt.Sprintf(` style="%s"`, tf.Style)
	}
	if tf.Text != "" {
		defaultValue = fmt.Sprintf(` value="%s"`, tf.Text)
	}
	return fmt.Sprintf(`<input type="%s" id="%s"%s%s%s>`, tf.InputType, tf.ID()(), defaultValue, style, tf.BoundEvents)
}

// NewInputElement creates an input element
func NewInputElement(name string, inputType InputType) *InputElement {
	return &InputElement{
		Base: Base{ID: name,
			BoundEvents: &BoundEvents{
				EventType(OnChange): &Binding{},
			},
		},
		InputType: inputType,
	}
}

//Children returns the child elements
func (tf *InputElement) Children() *Elements { return &Elements{slice: []*Element{}} }

//OptionElement for use in SelectElement
type OptionElement struct {
	Text  string
	Value string
}

//OptionSlice is a slice of OptionElements
type OptionSlice []*OptionElement

func (oe *OptionElement) String() string {
	return fmt.Sprintf(`<option value="%s">%s</option>`, oe.Value, oe.Text)
}

func (oes *OptionSlice) String() string {
	s := ""
	for _, oe := range *oes {
		s += oe.String()
	}
	return s
}

//SelectElement provides a selection drop-down
type SelectElement struct {
	Base
	Options *OptionSlice
}

//NewSelectElement creates a new select element
func NewSelectElement(name, functionName string) *SelectElement {
	os := OptionSlice([]*OptionElement{})
	return &SelectElement{
		Base: Base{
			ID: name,
			BoundEvents: &BoundEvents{
				ChangeEvent: &Binding{FunctionName: functionName},
			},
		},
		Options: &os,
	}
}

//AddOption will add an option to the SelectElement
func (se *SelectElement) AddOption(label, value string) {
	o := *se.Options
	os := OptionSlice(append([]*OptionElement(o), &OptionElement{Text: label, Value: value}))
	se.Options = &os
}

func (se *SelectElement) String() string {
	binding := ""
	if se.BoundEvents != nil {
		for e, bnd := range *se.BoundEvents {
			binding += fmt.Sprintf(` %s="%s()"`, e, bnd.FunctionName)
		}
	}
	return fmt.Sprintf(`<select id="%s"%s>%s</select>`, se.ID()(), binding, se.Options)
}

//Children returns an empty list
func (se *SelectElement) Children() *Elements { return &Elements{slice: []*Element{}} }
