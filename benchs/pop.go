package benchs

import (
	"fmt"

	"github.com/gobuffalo/pop"
	"database/sql"
)

var popdb *pop.Connection

type PModel struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Rightx   bool   `db:"rightx"`  //escaping problem
	Counter int64  `db:"counter"`
}

func NewPModel() *PModel {
	m := new(PModel)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Rightx = true
	m.Counter = 1000

	return m
}

// initDBPop recreates tables before executing any benchmark.
func initDBPop() {

	sqls := []string{
		`DROP TABLE IF EXISTS p_models;`,
		`CREATE TABLE p_models (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			rightx boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT p_models_pkey PRIMARY KEY (id)
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

func PopConnect(name string) (*pop.Connection, error) {
	deet := &pop.ConnectionDetails{
		URL: "postgres://bench:pass@localhost:5432/benchdb?sslmode=disable",
		Pool:4,
	}
	if c, err := pop.NewConnection(deet); err != nil {
		return nil, err
	} else {
		pop.Connections[name] = c
		return pop.Connections[name], nil
	}
}

func init() {
	st := NewSuite("pop")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000 * ORM_MULTI, 0, PopInsert)
		st.AddBenchmark("BulkInsert 100 row", 500 * ORM_MULTI, 0, PopInsertMulti)
		st.AddBenchmark("Update", 2000 * ORM_MULTI, 0, PopUpdate)
		st.AddBenchmark("Read", 4000 * ORM_MULTI, 0, PopRead)
		st.AddBenchmark("MultiRead limit 1000", 2000 * ORM_MULTI, 1000, PopReadSlice)
		var err error
		popdb, err = PopConnect("bechdb")
		if err != nil {
			fmt.Printf("Can not connect to db err: %v\n", err)
		}
		err = popdb.Open()
		if err != nil {
			fmt.Printf("Can not connect to db err: %v\n", err)
		}
		//pop.Debug = true
	}
}

func PopInsert(b *B) {
	var m *PModel
	wrapExecute(b, func() {
		initDBPop()
		m = NewPModel()
	})
	for i := 0; i < b.N; i++ {
		m.Id = 0
		if err := popdb.Create(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PopInsertMulti(b *B) {
	panic(fmt.Errorf("Problematic bulk insert, too slow"))
	var ms []PModel
	wrapExecute(b, func() {
		initDBPop()
	})

	for i := 0; i < b.N; i++ {
		ms = make([]PModel, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, *NewPModel())
		}
		if err := popdb.Create(&ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PopUpdate(b *B) {
	var m *PModel
	wrapExecute(b, func() {
		initDBPop()
		m = NewPModel()
		if err := popdb.Create(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	for i := 0; i < b.N; i++ {
		if err := popdb.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PopRead(b *B) {
	var m *PModel
	wrapExecute(b, func() {
		initDBPop()
		m = NewPModel()
		if err := popdb.Create(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := popdb.First(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PopReadSlice(b *B) {
	var m *PModel
	wrapExecute(b, func() {
		initDBPop()
		m = NewPModel()
		for i := 0; i < b.L; i++ {
			m.Id = 0
			if err := popdb.Create(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*PModel
		if err := popdb.Where("id > ?", 0).Limit(b.L).All(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
