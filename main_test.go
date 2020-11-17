package dali

import (
	"fmt"
	"testing"
)

func TestStyleSheet(t *testing.T) {
	s := StyleSheet{}
	if fmt.Sprintf("%s", s) != "" {
		t.Errorf("Expected empty string.  Got %s", s)
	}
	s = StyleSheet{URL: "MyStyleSheet.css"}
	expected := `<link rel="stylesheet" href="MyStyleSheet.css">`
	if fmt.Sprintf("%s", s) != expected {
		t.Errorf("Expected %s but got %s", expected, s)
	}
}

func TestButton(t *testing.T) {
	b := Button{
		ID:         "MyTestButton",
		ButtonText: "This Is A Button",
	}
	expected := `<button id="MyTestButton" onclick="do_MyTestButton()" >This Is A Button</button>`
	if fmt.Sprintf("%s", b) != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, b)
	}
}

func TestWindow(t *testing.T) {
	w := NewWindow(300, 300, "/home/matthewp/Downloads", "")
	p := NewPane("TestPane")
	b := Button{}
	b.ID = "buttonOne"
	b.ButtonText = "This is Button One"
	p.AddElement(b)
	w.AddPane(p)
	expected := `<html><body><div id="TestPane"><button id="buttonOne" onclick="do_buttonOne()" >This is Button One</button> </div></body></html>`

	if fmt.Sprintf("%s", w) != expected {
		t.Errorf("Expected %s but got %s", expected, w)
	}
	p.Style = "border:solid 1px #000000"
	w.AddPane(p)
	expected = `<html><body><div id="TestPane" style="border:solid 1px #000000"><button id="buttonOne" onclick="do_buttonOne()" >This is Button One</button> </div></body></html>`
	if fmt.Sprintf("%s", w) != expected {
		t.Errorf("Expected %s but got %s", expected, w)
	}
}
