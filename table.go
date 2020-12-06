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
	Elements *Elements
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
func NewTableElement(name string, columns, rows int) *Table {
	tableRows := []*Row{}

	tab := &Table{
		ColumnCount: columns,
		RowCount:    rows,
		Base: Base{
			ID: name,
		},
		Rows: tableRows,
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

	tab.Rows = tableRows

	return tab
}

//Bindings returns empty Bindings on table
func (tab *Table) Bindings() *Binding { return nil }

//Children will return each of the table  Cells
func (tab *Table) Children() *Elements {
	return tab.Elements
}

// GetCell provides access to the cell at row, column
func (tab *Table) GetCell(column, row int) (*Cell, error) {
	if row > tab.RowCount-1 || column > tab.ColumnCount-1 {
		return nil, errors.New("dimensions out of range")
	}
	return tab.Rows[row].Cells[column], nil
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
