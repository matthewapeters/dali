package dali

import "testing"

func TestSpan(t *testing.T) {

	s := NewSpanElement("", "span1", "This is Span 1")
	if s.ID() != "span1" {
		t.Errorf("expected span1 get %s", s.ID())
	}

}
