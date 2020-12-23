package dali

// HeaderLevel enums header sizes
type HeaderLevel int

// BR a break tag
type BR struct {
	Base
}

// Children will return an empty Elements
func (br *BR) Children() *Elements {
	return &Elements{slice: []*Element{}}
}

//LineBreak generates a BR tag
func LineBreak() *BR {
	return &BR{
		Base: Base{ElementClass: "br"},
	}
}

func (br *BR) String() string {
	return "<br/>"
}
