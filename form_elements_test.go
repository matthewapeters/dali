package dali

import "testing"

func TestSelectElement(t *testing.T) {
	se := NewSelectElement("testElement", "do_testElement")
	se.AddOption("Label One", "1")
	se.AddOption("Label Two", "2")
	se.AddOption("Label Three", "3")

	expected := `<select id="testElement" onselect="do_testElement"><option value="1">Label One</option><option value="2">Label Two</option><option value="3">Label Three</option></select>`
	if se.String() != expected {
		t.Errorf(`Expected %s but got %s`, expected, se)
	}

}
