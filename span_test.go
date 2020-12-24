package dali_test

import (
	"testing"

	"github.com/matthewapeters/dali"
)

func TestSpan(t *testing.T) {

	s := dali.NewSpanElement("", "span1", "This is Span 1")
	if s.ID() != "span1" {
		t.Errorf("expected span1 get %s", s.ID())
	}

}
