package dali

import (
	"errors"
	"fmt"
)

//Heading is a column title
type Heading struct {
	Span int
	Text string
	Base
}

//Children return nil for Headings
func (hd *Heading) Children() *Elements { return &Elements{} }

//Bindings returns the bindings for Headings
func (hd *Heading) Bindings() *map[EventType]*Binding { return hd.BoundEvents }

func (hd *Heading) String() string {
	style := ""
	binding := ""
	if hd.Style != "" {
		style = fmt.Sprintf(` style="%s"`, hd.Style)
	}
	if hd.BoundEvents != nil {
		for e, bnd := range *hd.BoundEvents {
			binding += fmt.Sprintf(` %s="%s()"`, e, bnd.FunctionName)
		}
	}
	return fmt.Sprintf(`<th id="%s"%s%s>%s</th>`, hd.Name(), style, binding, hd.Text)
}

//Headings is the collection of headings
type Headings []*Heading

func (hs Headings) String() string {
	html := ""
	for _, h := range hs {
		html += fmt.Sprintf("%s", h)
	}
	return html
}

//Cell is a cell
type Cell struct {
	Span int
	Base
	Elements *Elements
}

func (cell *Cell) String() string {
	style := ""
	if cell.Style != "" {
		style = fmt.Sprintf(` style="%s"`, cell.Style)
	}
	return fmt.Sprintf(`<td%s>%s</td>`, style, cell.Elements)
}

//Bindings returns nil for Cell
func (cell *Cell) Bindings() *map[EventType]*Binding { return nil }

// Children returns the Cell's child elements
func (cell *Cell) Children() *Elements { return cell.Elements }

//Cells is a slice of Cell elements
type Cells []*Cell

func (cells Cells) String() string {
	s := ""
	for _, cell := range cells {
		s += fmt.Sprintf(`%c%c%c%s`, 10, 9, 9, cell)
	}
	s += fmt.Sprintf(`%c`, 10)
	return s
}

//HeadingRow is the row of headings
type HeadingRow struct {
	Headings
	Base
}

func (hr *HeadingRow) String() string {
	if len(hr.Headings) == 0 {
		return ""
	}
	style := ""
	if hr.Style != "" {
		style = fmt.Sprintf(` style="%s"`, hr.Style)
	}
	return fmt.Sprintf(`%c%c<tr%s>%s%c</tr>`, 10, 9, style, hr.Headings, 9)
}

//Row is a row
type Row struct {
	Cells
	Base
}

func (row *Row) String() string {
	style := ""
	if row.Style != "" {
		style = fmt.Sprintf(` style="%s"`, row.Style)
	}
	return fmt.Sprintf(`%c%c<tr%s>%s%c</tr>`, 10, 9, style, row.Cells, 9)
}

//Rows is a slice of Row elements
type Rows []*Row

func (rows Rows) String() string {
	s := ""
	for _, row := range rows {
		s += fmt.Sprintf(`%s`, row)
	}
	return s
}

//THead is the table Header
type THead struct {
	HeadingRow *HeadingRow
	Base
}

func (th *THead) String() string {
	if th.HeadingRow == nil || len(th.HeadingRow.Headings) == 0 {
		return ""
	}
	style := ""
	if th.Style != "" {
		style = fmt.Sprintf(` style="%s"`, th.Style)
	}
	return fmt.Sprintf(`%c<thead%s>%s%c</thead>`, 10, style, th.HeadingRow, 10)
}

//TBody is the table body
type TBody struct {
	Rows
	Base
}

func (tb *TBody) String() string {
	style := ""
	if tb.Style != "" {
		style = fmt.Sprintf(` style="%s"`, tb.Style)
	}
	return fmt.Sprintf(`%c<tbody%s>%s%c</tbody>`, 10, style, tb.Rows, 10)
}

// Table is a table
type Table struct {
	ColumnCount int
	RowCount    int
	THead       *THead
	TBody       *TBody
	Base
	Elements *Elements
}

func (tab *Table) String() string {
	style := ""
	if tab.Style != "" {
		style = fmt.Sprintf(` style="%s"`, tab.Style)
	}
	return fmt.Sprintf(`<table id="%s"%s>%s%s%c</table>`, tab.Name(), style, tab.THead, tab.TBody, 10)
}

// NewTableElement creates a new Table element
func NewTableElement(name string, columns, rows int, headings []string) *Table {
	tableRows := []*Row{}
	headingRow := HeadingRow{Headings: Headings{}}

	tab := &Table{
		ColumnCount: columns,
		RowCount:    rows,
		Base: Base{
			ID: name,
		},
		THead:    &THead{HeadingRow: &headingRow},
		TBody:    &TBody{Rows: tableRows},
		Elements: &Elements{},
	}

	for i, h := range headings {
		h := Heading{
			Base: Base{
				ID: fmt.Sprintf(`heading_%d`, i),
			},
			Text: h,
		}
		headingRow.Headings = append(headingRow.Headings, &h)
		tab.Elements.AddElement(&h)
	}

	for rowNum := 0; rowNum < rows; rowNum++ {
		row := Row{}
		row.Cells = Cells{}
		for colNum := 0; colNum < columns; colNum++ {
			c := Cell{
				Base:     Base{ID: fmt.Sprintf(`%s_%d_%d`, name, rowNum, colNum)},
				Elements: &Elements{slice: []*Element{}},
			}
			row.Cells = append(row.Cells, &c)
			tab.Elements.AddElement(&c)
		}
		tableRows = append(tableRows, &row)
	}

	tab.TBody.Rows = tableRows

	return tab
}

//Bindings returns empty Bindings on table
func (tab *Table) Bindings() *map[EventType]*Binding { return nil }

//Children will return each of the table  Cells
func (tab *Table) Children() *Elements {
	return tab.Elements
}

// GetCell provides access to the cell at row, column
func (tab *Table) GetCell(column, row int) (*Cell, error) {
	if row > tab.RowCount-1 || column > tab.ColumnCount-1 {
		return nil, errors.New("dimensions out of range")
	}
	return tab.TBody.Rows[row].Cells[column], nil
}

// AddCellElement allows elements to be added to cells based on row and column
func (tab *Table) AddCellElement(column, row int, element *Element) error {
	c, err := tab.GetCell(row, column)
	if err != nil {
		return err
	}
	c.Elements.AddElement(*element)
	return nil
}

//SetCommonStyles will set all of the header and cells to the provided style
func (tab *Table) SetCommonStyles(style string) {
	for _, el := range tab.Children().slice {
		(*el).SetStyle(style)
	}
}
