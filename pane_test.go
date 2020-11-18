package dali

import (
	"fmt"
	"testing"
)

func TestPane(t *testing.T) {

	P := NewPane("TestPaneOne")
	S := Span{Text: "Test Text", ID: "TestSpanOne"}
	P.AddElement(S)
	html := fmt.Sprintf("%s", P)
	expected := `<div id="TestPaneOne"><span name="TestSpanOne">Test Text</span></div>`
	if html != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, html)
	}
}
