package dali

import "fmt"

//CaptionType provides constants for describing how a figure's caption is to be positioned
type CaptionType string

const (
	//NoCaption will prevent the caption from displaying
	NoCaption = CaptionType("none")
	//BeforeCaption will make the caption the first child element
	BeforeCaption = CaptionType("before")
	//AfterCaption will make the caption the last child element
	AfterCaption = CaptionType("after")
)

// FigureCaption is a FigCaption element
type FigureCaption struct {
	Span
	CaptionType
}

func (fc *FigureCaption) String() string {
	return fmt.Sprintf(`<figcaption id="%s"%s>%s</figcaption>`, fc.ID(), fc.BoundEvents, &fc.Span)
}

// Figure is a potentially captioned element, commonly holding images and other content
type Figure struct {
	Base
	FigureCaption *FigureCaption
}

//NewFigure returns a new Figure element
func NewFigure(name, id, caption string, captPlacement CaptionType) *Figure {
	var te Element
	te = *Text(caption)
	return &Figure{
		Base: Base{
			ElementID:   id,
			ElementName: name,
			Elements:    &Elements{},
		},
		FigureCaption: &FigureCaption{
			CaptionType: captPlacement,
			Span: Span{
				Base: Base{
					ElementName: fmt.Sprintf(`%s_captionText`, name),
					ElementID:   fmt.Sprintf(`%s_captionText`, id),
					Elements:    &Elements{slice: []*Element{&te}},
				},
			},
		},
	}
}

//Children Returns the child elements for the figure
func (fig *Figure) Children() *Elements { return fig.Elements }

func (fig *Figure) String() string {
	switch fig.FigureCaption.CaptionType {
	case BeforeCaption:
		return fmt.Sprintf(`<figure id="%s">%s%s</figure>`, fig.ID(), fig.FigureCaption, fig.Elements)
	case AfterCaption:
		return fmt.Sprintf(`<figure id="%s">%s%s</figure>`, fig.ID(), fig.Elements, fig.FigureCaption)
	}
	return fmt.Sprintf(`<figure id="%s">%s</figure>`, fig.ID(), fig.Elements)
}
