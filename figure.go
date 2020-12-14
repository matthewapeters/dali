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
	Base
	Span
	CaptionType
}

func (fc *FigureCaption) String() string {
	return fmt.Sprintf(`<figcaption id="%s"%s>%s</figcaption>`, fc.Name(), fc.BoundEvents, fc.Span)
}

// Figure is a potentially captioned element, commonly holding images and other content
type Figure struct {
	Base
	Elements      *Elements
	FigureCaption *FigureCaption
}

//NewFigure returns a new Figure element
func NewFigure(name, caption string, captPlacement CaptionType) *Figure {
	return &Figure{
		Base: Base{
			ID: name,
		},
		Elements: &Elements{},
		FigureCaption: &FigureCaption{
			CaptionType: captPlacement,
			Base: Base{
				ID: fmt.Sprintf(`%s_caption`, name),
			},
			Span: Span{
				Base: Base{ID: fmt.Sprintf(`%s_captionText`, name)},
				Text: caption,
			},
		},
	}
}

//Children Returns the child elements for the figure
func (fig *Figure) Children() *Elements { return fig.Elements }

func (fig *Figure) String() string {
	switch fig.FigureCaption.CaptionType {
	case BeforeCaption:
		return fmt.Sprintf(`<figure id="%s">%s%s</figure>`, fig.Name(), fig.FigureCaption, fig.Elements)
	case AfterCaption:
		return fmt.Sprintf(`<figure id="%s">%s%s</figure>`, fig.Name(), fig.Elements, fig.FigureCaption)
	}
	return fmt.Sprintf(`<figure id="%s">%s</figure>`, fig.Name(), fig.Elements)
}
