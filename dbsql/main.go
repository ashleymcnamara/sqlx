package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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

var (
	mysqluser = os.Getenv("MYSQL_ENV_MYSQL_USER")
	mysqlpw   = os.Getenv("MYSQL_ENV_MYSQL_PASSWORD")
	mysqlhost = os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	mysqlport = os.Getenv("MYSQL_PORT_3306_TCP_PORT")
	mysqldb   = os.Getenv("MYSQL_ENV_MYSQL_DATABASE")
)

func connectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mysqluser, mysqlpw, mysqlhost, mysqlport, mysqldb)
}

func main() {
	db, err := sql.Open("mysql",
		connectionString())
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
