package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

type MatrixNode struct {
	empty              bool
	Row, Column, Value int
	Next, Prev         *MatrixNode
}

func (n *MatrixNode) Insert(row, column, value int) bool {
	if n.empty {
		n.empty = false
		n.Row = row
		n.Column = column
		n.Value = value
		n.Next = &MatrixNode{empty: true}
		return true
	}
	return n.Next.Insert(row, column, value)
}

func (n *MatrixNode) Find(row, column int) *MatrixNode {
	if n.Column == column && n.Row == row {
		return n
	}
	if n.Next != nil {
		return n.Next.Find(row, column)
	}
	return nil
}

func (n *MatrixNode) Update(row, column, value int) bool {
	node := n.Find(row, column)
	if node != nil {
		node.Value = value
		return true
	}
	return false
}

// Define the Matrix type here.
type Matrix struct {
	origin string
	Root   *MatrixNode
}

func New(s string) (*Matrix, error) {
	m := &Matrix{Root: &MatrixNode{empty: true}, origin: s}
	matrix_rows := strings.Split(s, "\n")
	columns_per_row := len(strings.Split(matrix_rows[0], " "))
	// rows
	for r, row := range matrix_rows {
		matrix_cols := strings.Split(strings.TrimSpace(row), " ")
		if len(matrix_cols) != columns_per_row {
			return m, fmt.Errorf("uneven rows: %v != %v", columns_per_row, len(matrix_cols))
		}
		// columns
		for c, col := range matrix_cols {
			if len(col) > 0 {
				value, err := strconv.Atoi(col)
				if err != nil {
					return m, err
				}
				m.Root.Insert(r, c, value)
			}
		}
	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	cols := [][]int{}
	var col []int
	rows := m.Rows()
	for x := 0; x < len(rows[0]); x++ {
		col = []int{}
		for y := 0; y < len(rows); y++ {
			col = append(col, rows[y][x])
		}
		cols = append(cols, col)
	}
	return cols
}

func (m *Matrix) Rows() [][]int {
	result := [][]int{}
	row := []int{}
	current := m.Root
	if current.empty {
		return result
	}
	row = append(row, current.Value)
	i := current.Row
	for !current.Next.empty {
		current = current.Next
		if current.Row > i {
			i = current.Row
			result = append(result, row)
			row = []int{}
		}
		row = append(row, current.Value)
	}
	if len(row) > 0 {
		result = append(result, row)
	}
	return result
}

func (m *Matrix) Set(row, col, val int) bool {
	return m.Root.Update(row, col, val)
}
