package benchs

import (
	"fmt"

	"database/sql"

	"github.com/go-xorm/xorm"
)

var xo *xorm.Session

// initDB2 recreates tables before executing any benchmark.
func initDB2() {
	sqls := []string{
		`DROP TABLE IF EXISTS xorm_model;`,
		`CREATE TABLE xorm_model (
		id integer NOT NULL,
		name text NOT NULL,
		title text NOT NULL,
		fax text NOT NULL,
		web text NOT NULL,
		age integer NOT NULL,
		"right" boolean NOT NULL,
		counter bigint NOT NULL
		) WITH (OIDS=FALSE);`,
	}

	DB, err := sql.Open("postgres", ORM_SOURCE)
	checkErr(err)
	defer DB.Close()

	err = DB.Ping()
	checkErr(err)

	for _, sql := range sqls {
		_, err = DB.Exec(sql)
		checkErr(err)
	}
}

type XormModel struct {
	Id      int
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

func NewXormModel() *XormModel {
	m := new(XormModel)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

func init() {
	st := NewSuite("xorm")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, 0, XormInsert)
		st.AddBenchmark("BulkInsert 100 row", 500*ORM_MULTI, 0, XormInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, 0, XormUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, 0, XormRead)
		st.AddBenchmark("MultiRead limit 1000", 2000*ORM_MULTI, 1000, XormReadSlice)

		engine, _ := xorm.NewEngine("postgres", ORM_SOURCE)

		engine.SetMaxIdleConns(ORM_MAX_IDLE)
		engine.SetMaxOpenConns(ORM_MAX_CONN)

		xo = engine.NewSession()
		xo.NoCache()
	}
}

func XormInsert(b *B) {
	var m *XormModel
	wrapExecute(b, func() {
		initDB2()
		m = NewXormModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := xo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormInsertMulti(b *B) {
	var ms []XormModel
	wrapExecute(b, func() {
		initDB2()
	})
	for i := 0; i < b.N; i++ {
		ms = make([]XormModel, 100)
		for i := 0; i < 100; i++ {
			ms[i] = *NewXormModel()
		}
		if _, err := xo.InsertMulti(&ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormUpdate(b *B) {
	var m *XormModel
	wrapExecute(b, func() {
		initDB2()
		m = NewXormModel()
		if _, err := xo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := xo.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormRead(b *B) {
	var m *XormModel
	wrapExecute(b, func() {
		initDB2()
		m = NewXormModel()
		if _, err := xo.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := xo.NoCache().Get(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func XormReadSlice(b *B) {
	var m *XormModel
	wrapExecute(b, func() {
		initDB2()
		m = NewXormModel()
		for i := 0; i < b.L; i++ {
			m.Id = 0
			if _, err := xo.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []XormModel
		if err := xo.Table("models").Where("id > ?", 0).NoCache().Limit(b.L).Find(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}

}
