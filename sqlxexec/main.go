package main

import (
	"fmt"
	"log"
	"os"
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

type EmployeeReader interface {
	Get(id int) (*Employee, error)
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
	db, err := sqlx.Open("mysql",
		connectionString())
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
		BirthDate: time.Now(),
		HireDate:  time.Now(),
	}

	log.Println("Inserting Elvis")
	err = insertEmployee(db, e)
	if err != nil {
		log.Println(err)
	}
	log.Println("Inserted Elvis")
	// Prove it worked

	me := &MysqlEmployee{db: db}

	elvis, err := GetElvis(me, 500001)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(elvis)

}

type MysqlEmployee struct {
	db *sqlx.DB
}

func insertEmployee(db *sqlx.DB, e *Employee) error {
	// assignment
	var err error
	res, err := db.Exec("insert into employees(emp_no,first_name,last_name,birth_date,hire_date) values (?,?,?,?,?)", e.Number, e.FirstName, e.LastName, e.BirthDate, e.HireDate)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())
	return err
}

func (e *MysqlEmployee) Get(id int) (*Employee, error) {
	var emp Employee
	err := e.db.Get(&emp, "select emp_no, first_name, last_name from employees where emp_no=?", id)
	return &emp, err
}

func GetElvis(er EmployeeReader, id int) (*Employee, error) {
	return er.Get(id)
}
