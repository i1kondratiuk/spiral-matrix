package spiral

import (
	"errors"
)

// Spirals square matrix values
func Spiral(matrix [][]int) ([]int, error) {
	if len(matrix) == 0 || len(matrix[0]) == 0 || len(matrix) != len(matrix[0]) {
		return nil, errors.New("square matrix only")
	}

	rowsNumber, colsNumber := len(matrix), len(matrix[0])

	arrayMatrix := newArrayMatrix(rowsNumber, colsNumber)

	rowVector := make([]int, rowsNumber*colsNumber)
	for i := range rowVector {
		nextRowIndex, nextColIndex := arrayMatrix.goSpiral()
		rowVector[i] = matrix[nextRowIndex][nextColIndex]
	}

	return rowVector, nil
}

type arrayMatrix struct {
	rowsNumber, colsNumber int
	rowsShift, colsShift   int

	currentRowIndex, currentColumnIndex int

	leftIndex, rightIndex int
	topIndex, downIndex   int
}

func newArrayMatrix(rowsNumber, colsNumber int) *arrayMatrix {
	return &arrayMatrix{
		rowsNumber: rowsNumber,
		colsNumber: colsNumber,
		leftIndex:  0,
		rightIndex: colsNumber - 1,
		topIndex:   0,
		downIndex:  rowsNumber - 1,
	}
}

func (m *arrayMatrix) goSpiral() (int, int) {
	m.currentRowIndex += m.rowsShift
	m.currentColumnIndex += m.colsShift

	switch {
	case m.currentRowIndex == 0 && m.currentColumnIndex == 0:
		m.rowsShift, m.colsShift = 0, 1
	case m.currentColumnIndex+m.colsShift > m.rightIndex:
		m.rowsShift, m.colsShift = 1, 0
		m.topIndex++
	case m.currentRowIndex+m.rowsShift > m.downIndex:
		m.rowsShift, m.colsShift = 0, -1
		m.rightIndex--
	case m.currentColumnIndex+m.colsShift < m.leftIndex:
		m.rowsShift, m.colsShift = -1, 0
		m.downIndex--
	case m.currentRowIndex+m.rowsShift < m.topIndex:
		m.rowsShift, m.colsShift = 0, 1
		m.leftIndex++
	}

	return m.currentRowIndex, m.currentColumnIndex
}
