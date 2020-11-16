package dali

import (
	"fmt"
	"testing"
)

func TestDiv(t *testing.T) {

	P := NewDiv("TestDivOne")
	S := &Span{Text: "Test Text", Base: Base{ID: "TestSpanOne"}}
	P.Elements.AddElement(S)
	html := fmt.Sprintf("%s", P)
	expected := `<div id="TestDivOne"><span id="TestSpanOne">Test Text</span></div>`
	if html != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, html)
	}
}
