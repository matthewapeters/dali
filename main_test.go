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
	b := NewButton("This Is A Button", "MyTestButton", "do_MyTestButton")
	expected := `<button id="MyTestButton" onclick="do_MyTestButton()" >This Is A Button</button>`
	if fmt.Sprintf("%s", b) != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, b)
	}
	b.Binding.BoundFunction = func() { fmt.Println("Yaba Daba Do!") }
	d := NewDiv("")
	d.Elements.AddElement(b)

	W := NewWindow(100, 100, "", "")
	W.Elements.AddElement(d)
	W.BindChildren(nil)
	if len(W.Bindings) != 1 {
		t.Errorf("Expected 1 binding, found %d", len(W.Bindings))
	}
}

func TestWindow(t *testing.T) {
	w := NewWindow(300, 300, "/home/matthewp/Downloads", "")
	p := NewDiv("TestDiv")
	p.Elements.AddElement(NewButton("This is Button One", "buttonOne", "do_buttonOne"))
	if len(p.Elements.slice) != 1 {
		t.Errorf("Expect there to be 1 element, got %d", len(p.Elements.slice))
	}
	expected := `<div id="TestDiv"><button id="buttonOne" onclick="do_buttonOne()" >This is Button One</button></div>`
	if fmt.Sprintf("%s", p) != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, p)
	}

	// Add the divs to the body
	body := NewBodyElement()
	body.Elements.AddElement(p)

	expected = `<body><div id="TestDiv"><button id="buttonOne" onclick="do_buttonOne()" >This is Button One</button></div></body>`
	if fmt.Sprintf("%s", body) != expected {
		t.Errorf(`expected "%s" but got "%s"`, expected, body)
	}

	w.Elements.AddElement(body)
	expected = `<html><body><div id="TestDiv"><button id="buttonOne" onclick="do_buttonOne()" >This is Button One</button></div></body></html>`

	if fmt.Sprintf("%s", w) != expected {
		t.Errorf("Expected %s but got %s", expected, w)
	}
	p.StyleName = "border:solid 1px #000000"
	expected = `<html><body><div id="TestDiv" style="border:solid 1px #000000"><button id="buttonOne" onclick="do_buttonOne()" >This is Button One</button></div></body></html>`
	if fmt.Sprintf("%s", w) != expected {
		t.Errorf("Expected %s but got %s", expected, w)
	}
	p2 := NewDiv("TestDivTwo")
	body.Elements.AddElement(p2)
	expected = `<html><body><div id="TestDiv" style="border:solid 1px #000000"><button id="buttonOne" onclick="do_buttonOne()" >This is Button One</button></div><div id="TestDivTwo"></div></body></html>`
	if fmt.Sprintf("%s", w) != expected {
		t.Errorf(`Expected "%s" but got "%s"`, expected, w)
	}
}
