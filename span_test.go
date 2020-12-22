package dali

import "testing"

func TestSpan(t *testing.T) {

	s := Span{Text: "This is Span 1", Base: Base{ID: "span1"}}

	if s.ID()() != "span1" {
		t.Errorf("expected span1 get %s", s.ID()())
	}

}
