package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Employee represents the employee model in the database
type Employee struct {
	Number    int
	Birthdate time.Time
	FirstName string
	LastName  string
	Gender    string
	HireDate  time.Time
}

func main() {
	// hardcoded here - don't do this :)
	db, err := sql.Open("mysql",
		"docker:docker@tcp(127.0.0.1:3306)/employees?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	// always close the db when you're done.
	defer db.Close()
	err = noStructs(db)
	if err != nil {
		log.Fatal(err)
	}

	employees, err := withStructs(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, emp := range employees {
		log.Println("Struct:", emp.Number, emp.FirstName, emp.LastName)
	}

}

func noStructs(db *sql.DB) error {
	var (
		empno int
		fname string
		lname string
	)
	rows, err := db.Query("select emp_no, first_name, last_name  from employees limit 10")
	if err != nil {
		return err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&empno, &fname, &lname)
		if err != nil {
			return err
		}
		log.Println(empno, fname, lname)
	}
	err = rows.Err()
	return err
}

func withStructs(db *sql.DB) ([]*Employee, error) {
	// we'll return this
	var employees []*Employee
	var (
		emp_no int
		fname  string
		lname  string
	)
	rows, err := db.Query("select emp_no, first_name, last_name  from employees limit 10")
	if err != nil {
		return employees, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&emp_no, &fname, &lname)
		if err != nil {
			return employees, err
		}
		emp := &Employee{
			Number:    emp_no,
			FirstName: fname,
			LastName:  lname,
		}
		employees = append(employees, emp)
	}
	err = rows.Err()
	return employees, err
}
