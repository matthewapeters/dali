package dali_test

import (
	"testing"

	"github.com/matthewapeters/dali"
)

func TestSelectElement(t *testing.T) {
	se := dali.NewSelectElement("testElement", "testElement", "do_testElement")
	se.AddOption("Label One", "1")
	se.AddOption("Label Two", "2")
	se.AddOption("Label Three", "3")

	expected := `<select id="testElement" onchange="do_testElement()"><option value="1">Label One</option><option value="2">Label Two</option><option value="3">Label Three</option></select>`
	if se.String() != expected {
		t.Errorf(`Expected %s but got %s`, expected, se)
	}

}
