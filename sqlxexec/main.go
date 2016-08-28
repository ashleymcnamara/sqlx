package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Employee represents the employee model in the database
// 'db' struct tags tell sqlx how to map data
type Employee struct {
	Number    int       `db:"emp_no"`
	BirthDate time.Time `db:"birth_date"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Gender    string    `db:"gender"`
	HireDate  time.Time `db:"hire_date"`
}

func main() {
	// hardcoded here - don't do this :)
	db, err := sqlx.Open("mysql",
		"docker:docker@tcp(127.0.0.1:3306)/employees?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	// always close the db when you're done.
	defer db.Close()

	// max emp_no is 499999 in sample database
	e := &Employee{
		Number:    500001,
		FirstName: "Elvis",
		LastName:  "Presley",
		Gender:    "M",
		BirthDate: time.Parse("Jan 8, 1935"),
		HireDate:  time.Now(),
	}
	err = insertEmployee(e)

	// Prove it worked
	elvis, err := getElvis(db)
	if err != nil {
		log.Panic(err)
	}
	log.Println(elvis)

}

func insertEmployee(db *sqlx.DB, e *Employee) error {
	// Assignment -- fill in the blanks
	tx := db.MustBegin()
	// hint - use tx.NamedExec()

	//your insert here

	err := tx.Commit()
	return err
}

func getElvis(db *sqlx.DB) (*Employee, error) {
	var king Employee
	err := db.Get(&king, "select emp_no, first_name, last_name from employees where emp_no=?", 500001)
	return &king, err

}
