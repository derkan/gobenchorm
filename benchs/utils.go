package benchs

import (
	"database/sql"
	"fmt"
	"os"
)

type Model struct {
	Id      int `qbs:"pk" orm:"auto" gorm:"primary_key" db:"id"`
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

// NewModel initializes a new model struct for inserts
func NewModel() *Model {
	m := new(Model)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

var (
	ORM_MULTI    int
	ORM_MAX_IDLE int
	ORM_MAX_CONN int
	ORM_SOURCE   string
)

// checkErr prints and exists on error
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

//wrapExecute sets timer for a benchmark step
func wrapExecute(b *B, cbk func()) {
	b.StopTimer()
	defer b.StartTimer()
	cbk()
}

// initDB recreates tables before executing any benchmark.
func initDB() {
	sqls := []string{
		`DROP TABLE IF EXISTS models;`,
		`CREATE TABLE models (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT models_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
		`DROP TABLE IF EXISTS model;`,
		`CREATE TABLE model (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT model_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
	}

	DB, err := sql.Open("postgres", ORM_SOURCE)
	checkErr(err)
	defer DB.Close()

	err = DB.Ping()
	checkErr(err)

	for _, stmt := range sqls {
		_, err = DB.Exec(stmt)
		checkErr(err)
	}
}
