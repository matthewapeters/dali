package dali

import (
	"errors"
	"fmt"
)

//Cell is a cell
type Cell struct {
	Span int
	Base
	Elements *Elements
}

func (cell *Cell) String() string {
	style := ""
	if cell.Style != "" {
		style = fmt.Sprintf(` style="%s"`, cell.Styles())
	}
	return fmt.Sprintf(`<td%s>%s</td>`, style, cell.Elements)
}

//Cells is a slice of Cell elements
type Cells []*Cell

func (cells Cells) String() string {
	s := ""
	for _, cell := range cells {
		s += fmt.Sprintf(`		%s
`, cell)
	}

	return s
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
	return fmt.Sprintf(`<tr%s>
%s	</tr>`, style, row.Cells)
}

//Rows is a slice of Row elements
type Rows []*Row

func (rows Rows) String() string {
	s := ""
	for _, row := range rows {
		s += fmt.Sprintf(`	%s
`, row)
	}
	return s
}

// Table is a table
type Table struct {
	ColumnCount int
	RowCount    int
	Rows
	Base
}

func (tab *Table) String() string {
	style := ""
	if tab.Style != "" {
		style = fmt.Sprintf(` style="%s"`, tab.Style)
	}
	return fmt.Sprintf(`<table id="%s"%s>
%s</table>`, tab.Name(), style, tab.Rows)
}

// NewTableElement creates a new Table element
func NewTableElement(name string, rows, columns int) *Table {
	tableRows := []*Row{}

	for rowNum := 0; rowNum < rows; rowNum++ {
		row := Row{}
		row.Cells = Cells{}
		for colNum := 0; colNum < columns; colNum++ {
			row.Cells = append(row.Cells, &Cell{})
		}
		tableRows = append(tableRows, &row)
	}

	return &Table{
		ColumnCount: columns,
		RowCount:    rows,
		Base: Base{
			ID: name,
		},
		Rows: tableRows,
	}
}

// GetCell provides access to the cell at row, column
func (tab *Table) GetCell(row, column int) (*Cell, error) {
	if row > tab.RowCount || column > tab.ColumnCount {
		return nil, errors.New("dimensions out of range")
	}
	return tab.Rows[row].Cells[column], nil
}

// AddElement allows elements to be added to cells based on row and column
func (tab *Table) AddElement(row, column int, element *Element) error {
	c, err := tab.GetCell(row, column)
	if err != nil {
		return err
	}
	c.Elements.AddElement(*element)
	return nil
}
