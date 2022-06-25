package coverage

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPeopleEmptyLen(t *testing.T) {
	p := People{}

	if p.Len() != 0 {
		t.Errorf("zero length expected")
	}
}

func TestPeopleSwap(t *testing.T) {
	persons := []Person{
		{
			firstName: "test2",
			lastName: "test2",
			birthDay: time.Now(),
		},
		{
			firstName: "test1",
			lastName: "test1",
			birthDay: time.Now(),
		},
	}
	
	p := People{persons[0], persons[1]}

	p.Swap(0, 1)

	if p[0].firstName != persons[1].firstName || p[1].firstName != persons[0].firstName {
		t.Errorf(
			"item0: first name '%s', expected '%s'; item1: first name '%s', expected '%s'",
			p[0].firstName, persons[1].firstName,
			p[1].firstName, persons[0].firstName,
		)
	}
}

func TestPeopleSortDate(t *testing.T) {
	persons := []Person{
		{
			firstName: "test1",
			lastName: "test1",
			birthDay: time.Now(),
		},
		{
			firstName: "test1",
			lastName: "test1",
			birthDay: time.Now().AddDate(0, 0, 1),
		},
	}
	
	p := People{persons[0], persons[1]}

	sort.Sort(p)

	if p[0].birthDay.Equal(persons[0].birthDay) {
		t.Errorf(
			"first item: birthday '%s', expected '%s'",
			p[0].birthDay, persons[1].birthDay,
		)
	}
}

func TestPeopleSortFirstname(t *testing.T) {
	persons := []Person{
		{
			firstName: "test2",
			lastName: "test1",
			birthDay: time.Now(),
		},
		{
			firstName: "test1",
			lastName: "test1",
			birthDay: time.Now(),
		},
	}
	
	p := People{persons[0], persons[1]}

	sort.Sort(p)

	if p[0].firstName == persons[0].firstName {
		t.Errorf(
			"first item: first name '%s', expected '%s'",
			p[0].firstName, persons[1].firstName,
		)
	}
}

func TestPeopleSortLastname(t *testing.T) {
	persons := []Person{
		{
			firstName: "test1",
			lastName: "test2",
			birthDay: time.Now(),
		},
		{
			firstName: "test1",
			lastName: "test1",
			birthDay: time.Now(),
		},
	}
	
	p := People{persons[0], persons[1]}

	sort.Sort(p)

	if p[0].lastName == persons[0].lastName {
		t.Errorf(
			"last item: last name '%s', expected '%s'",
			p[0].lastName, persons[1].lastName,
		)
	}
}

func makeMatrix(s string, t *testing.T) (*Matrix, error) {
	m, err := New(s)

	if m == nil {
		t.Errorf("matrix is nil")
		return nil, errors.New("matrix is nil")
	}

	if err != nil {
		t.Errorf("matrix create failed: %s", err)
		return nil, err
	}

	return m, nil
}

func TestMatrixSingleCell(t *testing.T) {
	m, err := makeMatrix("1", t)

	if err != nil {
		return
	}
	
	if m != nil && m.cols != 1 && m.rows != 1 {
		t.Errorf("matrix size %d:%d, expected 1:1", m.cols, m.rows)
	}
}

func TestMatrixRowLengthMismatch(t *testing.T) {
	m, err := New("1\n2 3")

	if m != nil && err == nil {
		t.Errorf("row length mismatch error expected")
	}
}

func TestMatrixNotADigit(t *testing.T) {
	m, err := New("Z")

	if m != nil && err == nil {
		t.Errorf("parse error expected")
	}
}

func checkIfEquals(a, b [][]int) error {
	if len(a) != len(b) {
		return fmt.Errorf("rows count mismatch: len(A)=%d, len(B)=%d", len(a), len(b))
	}

	for indexRow, rowA := range a {
		rowB := b[indexRow]
		
		if len(rowA) != len(rowB) {
			return fmt.Errorf(
				"column count mismatch: len(A[%d])=%d, len(B[%d])=%d",
				indexRow, len(rowA),
				indexRow, len(rowB),
			)
		}

		for indexCell, cellA := range rowA {
			cellB := rowB[indexCell]
		
			if cellA != cellB {
				return fmt.Errorf(
					"cell value mismatch: A[%d:%d]=%d, B[%d:%d]=%d",
					indexRow, indexCell, cellA,
					indexRow, indexCell, cellB,
				)
			}
		}
	}

	return nil
}

const testMatrixData = "1 2\n3 4"

func TestMatrixRows(t *testing.T) {
	m, err := makeMatrix(testMatrixData, t)

	if err != nil {
		return
	}

	expect := [][]int{
		{1, 2},
		{3, 4},
	}

	rows := m.Rows()

	if rows == nil {
		t.Errorf("call Rows() failed")
		return
	}

	if err = checkIfEquals(expect, rows); err != nil {
		t.Errorf("result of Rows() is not expected: %s", err)
	}
}

func TestMatrixCols(t *testing.T) {
	m, err := makeMatrix(testMatrixData, t)

	if err != nil {
		return
	}

	expect := [][]int{
		{1, 3},
		{2, 4},
	}

	cols := m.Cols()

	if cols == nil {
		t.Errorf("call Cols() failed")
		return
	}

	if err = checkIfEquals(expect, cols); err != nil {
		t.Errorf("result of Cols() is not expected: %s", err)
	}
}

func TestMatrixSetOk(t *testing.T) {
	m, err := makeMatrix(testMatrixData, t)

	if err != nil {
		return
	}

	ok := m.Set(1, 1, 100)

	if ok != true {
		t.Errorf("set value failed")
		return
	}
}

func TestMatrixSetOutOfRange(t *testing.T) {
	m, err := makeMatrix(testMatrixData, t)

	if err != nil {
		return
	}

	ok := m.Set(0, 10, 10)

	if ok == true {
		t.Errorf("set value that is out of range, error expected")
	}
}
