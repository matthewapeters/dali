package dali

import (
	"fmt"
	"testing"
)

func TestTable(t *testing.T) {
	tab := NewTableElement("testTable", 1, 1)
	expected := `<table id="testTable">
	<tr>
		<td></td>
	</tr>
</table>`
	html := fmt.Sprintf("%s", tab)
	if html != expected {
		t.Errorf(`Expected:
%s  
but got: 
%s`, expected, html)
	}

	tab = NewTableElement("testTable", 5, 2)
	expected = `<table id="testTable">
	<tr>
		<td></td>
		<td></td>
	</tr>
	<tr>
		<td></td>
		<td></td>
	</tr>
	<tr>
		<td></td>
		<td></td>
	</tr>
	<tr>
		<td></td>
		<td></td>
	</tr>
	<tr>
		<td></td>
		<td></td>
	</tr>
</table>`
	html = fmt.Sprintf("%s", tab)
	if html != expected {
		t.Errorf(`Expected:
%s  
but got: 
%s`, expected, html)
		if len(expected) != len(html) {
			t.Errorf("expected length: %d got length: %d", len(expected), len(html))
		}
	}

	tab = NewTableElement("cellTest", 2, 2)
	var b, b2 Element
	b = NewButton("0,0", "button1", "doButton1")
	b2 = NewButton("1,1", "button2", "doButton2")

	tab.AddElement(0, 0, &b)
	tab.AddElement(1, 1, &b2)

	expected = `<table id="cellTest">
	<tr>
		<td><button id="button1" onclick="doButton1()" >0,0</button></td>
		<td></td>
	</tr>
	<tr>
		<td></td>
		<td><button id="button2" onclick="doButton2()" >1,1</button></td>
	</tr>
</table>`
	html = fmt.Sprintf("%s", tab)
	if html != expected {
		t.Errorf(`Expected:
%s  
but got: 
%s`, expected, html)
	}

	c, err := tab.GetCell(1, 1)
	if err != nil {
		t.Error(err)
	}
	expected = `<td><button id="button2" onclick="doButton2()" >1,1</button></td>`
	html = fmt.Sprintf("%s", c)
	if html != expected {
		t.Errorf(`Expected:
%s  
but got: 
%s`, expected, html)
	}
	c, err = tab.GetCell(3, 3)
	if err == nil {
		t.Errorf("expected dimensions out of range")
	}
}
