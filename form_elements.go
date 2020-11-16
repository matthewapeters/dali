package dali

import "fmt"

//InputType identifies input element types
type InputType string

// InputEventType identifies the binding event type
type InputEventType string

const (
	//InputTypes

	ButtonInput        = InputType("button")
	CheckboxInput      = InputType("checkbox")
	ColorInput         = InputType("color")
	DateInput          = InputType("date")
	DatetimeLocalInput = InputType("datetime-local")
	EmailInput         = InputType("email")
	FileInput          = InputType("file")
	HiddenInput        = InputType("hidden")
	ImageInput         = InputType("image")
	MonthInput         = InputType("month")
	NumberInput        = InputType("number")
	PasswordInput      = InputType("password")
	RadioInput         = InputType("radio")
	RangeInput         = InputType("range")
	ResetInput         = InputType("reset")
	SearchInput        = InputType("search")
	SubmitInput        = InputType("submit")
	TelInput           = InputType("tel")
	TextInput          = InputType("text")
	TimeInput          = InputType("time")
	URLInput           = InputType("url")
	WeekInput          = InputType("week")

	//InputEventTypes

	OnBlur        = InputEventType("onblur")
	OnChange      = InputEventType("onchange")
	OnFocus       = InputEventType("onfocus")
	OnInput       = InputEventType("oninput")
	OnInvalid     = InputEventType("oninvalid")
	OnReset       = InputEventType("onreset")
	OnSearch      = InputEventType("onsearch")
	OnSelect      = InputEventType("onselect")
	OnSubmit      = InputEventType("onsubmit")
	OnContextmenu = InputEventType("oncontextmenu")
)

// InputElement is for inputting values
type InputElement struct {
	InputType
	Text string
	Base
	Binding
	InputEventType
}

func (tf *InputElement) String() string {
	style := ""
	defaultValue := ""
	boundEvent := ""
	if tf.Style != "" {
		style = fmt.Sprintf(` style="%s"`, tf.Style)
	}
	if tf.Text != "" {
		defaultValue = fmt.Sprintf(` value="%s"`, tf.Text)
	}
	if tf.Binding.FunctionName != "" && tf.InputEventType != "" {
		boundEvent = fmt.Sprintf(`%s="%s()"`, tf.InputEventType, tf.Bindings().FunctionName)
	}

	return fmt.Sprintf(`<input type="%s" id="%s"%s%s%s>`, tf.InputType, tf.Name(), defaultValue, style, boundEvent)
}

// NewInputElement creates an input element
func NewInputElement(name string, inputType InputType) *InputElement {
	return &InputElement{
		Base:           Base{ID: name},
		InputEventType: OnChange,
		InputType:      inputType,
	}
}

//Children returns the child elements
func (tf *InputElement) Children() *Elements { return &Elements{slice: []*Element{}} }

// Bindings returns the element Bindings
func (tf *InputElement) Bindings() *Binding { return &tf.Binding }

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
	Binding
	InputEventType
	Options *OptionSlice
}

//NewSelectElement creates a new select element
func NewSelectElement(name, functionName string) *SelectElement {
	os := OptionSlice([]*OptionElement{})
	return &SelectElement{
		Base:           Base{ID: name},
		Binding:        Binding{FunctionName: functionName},
		InputEventType: OnChange,
		Options:        &os,
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
	if se.Binding.FunctionName != "" {
		binding = fmt.Sprintf(` %s="%s()"`, se.InputEventType, se.Binding.FunctionName)
	}
	return fmt.Sprintf(`<select id="%s"%s>%s</select>`, se.Name(), binding, se.Options)
}

//Children returns an empty list
func (se *SelectElement) Children() *Elements { return &Elements{slice: []*Element{}} }

// Bindings returns the Binding
func (se *SelectElement) Bindings() *Binding { return &se.Binding }
