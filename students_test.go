package coverage

import (
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
