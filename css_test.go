package dali_test

import (
	"fmt"
	"testing"

	"github.com/matthewapeters/dali"
)

func TestCss(t *testing.T) {
	ss := dali.NewStyleSheet()
	ss.AddProperty(dali.Width, "100px")
	ss.AddProperty(dali.Height, "50px")
	expected := `height:50px;width:100px;`
	if fmt.Sprintf("%s", ss) != expected {
		t.Errorf(`expected:
		%s
		but got
		%s`, expected, fmt.Sprintf("%s", ss))
	}
	ss.AddProperty(dali.Display, "block")
	expected = `display:block;height:50px;width:100px;`
	if fmt.Sprintf("%s", ss) != expected {
		t.Errorf(`expected:
		%s
		but got
		%s`, expected, fmt.Sprintf("%s", ss))
	}
}
func TestStyleSheet(t *testing.T) {
	s := dali.NewStyleSheet()
	if fmt.Sprintf("%s", s) != "" {
		t.Errorf("Expected empty string.  Got %s", s)
	}
	s = dali.NewStyleSheet()
	s.URL = "MyStyleSheet.css"
	expected := `<link rel="stylesheet" href="MyStyleSheet.css">`
	if fmt.Sprintf("%s", s) != expected {
		t.Errorf("Expected %s but got %s", expected, s)
	}
}
