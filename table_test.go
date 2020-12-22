package dali

import (
	"fmt"
	"math"
	"testing"
)

func TestTable(t *testing.T) {
	tab := NewTableElement("testTable", "testTable", 1, 1, []string{})
	expected := `<table name="testTable" id="testTable">
<tbody>
	<tr>
		<td></td>
	</tr>
</tbody>
</table>`
	html := fmt.Sprintf("%s", tab)
	if html != expected {
		t.Errorf(`Expected:
%s  
but got: 
%s`, expected, html)
		for i := 0; float64(i) < math.Min(float64(len(expected)), float64(len(html))); i++ {
			fmt.Printf("'%s'%d\t'%s'%d\n", string(expected[i]), expected[i], string(html[i]), html[i])
		}
	}

	tab = NewTableElement("testTable", "testTable", 2, 5, []string{})
	expected = `<table name="testTable" id="testTable">
<tbody>
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
</tbody>
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

	tab = NewTableElement("cellTest", "cellTest", 2, 2, []string{})
	var b, b2 Element
	b = NewButton("0,0", "button1", "button1", "doButton1")
	b2 = NewButton("1,1", "button2", "button2", "doButton2")

	tab.AddCellElement(0, 0, &b)
	tab.AddCellElement(1, 1, &b2)

	expected = `<table name="cellTest" id="cellTest">
<tbody>
	<tr>
		<td><button id="button1" onclick="doButton1()">0,0</button></td>
		<td></td>
	</tr>
	<tr>
		<td></td>
		<td><button id="button2" onclick="doButton2()">1,1</button></td>
	</tr>
</tbody>
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
	expected = `<td><button id="button2" onclick="doButton2()">1,1</button></td>`
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
