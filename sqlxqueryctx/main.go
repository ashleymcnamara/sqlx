package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

// countKey is the context key for the mysql record limit
const countKey = 0

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

	ctx := context.WithValue(context.Background(), countKey, 100)

	employees, err := getEmployees(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	for _, emp := range employees {
		log.Println("Best:", emp.Number, emp.FirstName, emp.LastName)
	}
}

func getEmployees(ctx context.Context, db *sqlx.DB) ([]*Employee, error) {
	count, ok := ctx.Value(countKey).(int)
	if !ok {
		count = 20
	}
	var employees []*Employee
	err := db.Select(&employees, "select emp_no, first_name, last_name  from employees limit ?", count)
	return employees, err

}
