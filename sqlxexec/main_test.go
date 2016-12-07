package main

import (
	"fmt"
	"testing"
	"time"
)

func TestEmployee(t *testing.T) {
	var empls = []struct {
		id   int
		val  *Employee
		name string
		good bool
	}{
		{val: &Employee{FirstName: "Brian"},
			id:   1,
			name: "Brian",
			good: true,
		},
		{val: &Employee{FirstName: "Elvis"},
			id:   2,
			name: "Brian",
			good: false,
		},
	}

	for _, tt := range empls {
		if tt.val.FirstName != tt.name && tt.good {
			t.Errorf("case: %d -- expected %s, got %s", tt.id, tt.name, tt.val.FirstName)
		}
	}
}

/*

type EmployeeReader interface {
	Get(id int) (*Employee, error)
}

*/

type FakeEmployee struct{}

func (f *FakeEmployee) Get(id int) (*Employee, error) {

	e := &Employee{
		Number:    500001,
		FirstName: "Elvis",
		LastName:  "Presley",
		Gender:    "M",
		BirthDate: time.Now(),
		HireDate:  time.Now(),
	}
	return e, nil
}

func TestGetEmployee(t *testing.T) {
	fe := &FakeEmployee{}
	emp, err := GetElvis(fe, 500001)
	if err != nil {
		t.Error(err)
	}
	if emp.FirstName != "Elvis" {
		t.Errorf("Expected Elvis, got %s", emp.FirstName)
	}
}

func BenchmarkGetEmployee(b *testing.B) {

	fe := &FakeEmployee{}
	var emp *Employee
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		emp, _ = GetElvis(fe, 500001)
	}
	fmt.Println(emp)
}
