package dali

// BR a break tag
type BR struct {
	Base
}

// Children will return an empty Elements
func (br *BR) Children() *Elements {
	return &Elements{slice: []*Element{}}
}

//NewBreak generates a BR tag
func NewBreak() *BR {
	return &BR{
		Base: Base{ElementClass: "br"},
	}
}

func (br *BR) String() string {
	return "<br/>"
}
