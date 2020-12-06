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
}
