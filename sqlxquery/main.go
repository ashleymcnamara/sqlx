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
	Birthdate time.Time `db:"birth_date"`
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
	employees, err := better(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, emp := range employees {
		log.Println("Better:", emp.Number, emp.FirstName, emp.LastName)
	}

	employees, err = best(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, emp := range employees {
		log.Println("Best:", emp.Number, emp.FirstName, emp.LastName)
	}
}

func better(db *sqlx.DB) ([]*Employee, error) {
	// we'll return this
	var employees []*Employee
	rows, err := db.Queryx("select emp_no, first_name, last_name  from employees limit 10")
	// check to see if there's an error FIRST
	if err != nil {
		return employees, err
	}
	// THEN defer rows.Close() to avoid panic
	defer rows.Close()
	for rows.Next() {
		var e Employee
		err := rows.StructScan(&e) // HOORAY, less scanning
		if err != nil {
			log.Println(err)
			continue
		}
		employees = append(employees, &e)
	}
	err = rows.Err()
	return employees, err
}

func best(db *sqlx.DB) ([]*Employee, error) {
	var employees []*Employee
	err := db.Select(&employees, "select emp_no, first_name, last_name  from employees limit 10")
	return employees, err

}
