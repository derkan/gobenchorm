package benchs

import (
	"database/sql"
	"fmt"

	"github.com/eaigner/hood"
)

var hd *hood.Hood

type HdModel struct {
	Id      hood.Id `db:"id" sql:"pk"`
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

func NewHdModel() *HdModel {
	m := new(HdModel)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

// initHdDB recreates tables before executing any benchmark.
func initHdDB() {
	sqls := []string{
		`DROP TABLE IF EXISTS hd_model;`,
		`CREATE TABLE hd_model (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT hd_model_pkey PRIMARY KEY (id)
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

func init() {
	st := NewSuite("hood")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, 0, HdInsert)
		st.AddBenchmark("BulkInsert 100 row", 500*ORM_MULTI, 0, HdInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, 0, HdUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, 0, HdRead)
		st.AddBenchmark("MultiRead limit 1000", 2000*ORM_MULTI, 1000, HdReadSlice)
		db, err := sql.Open("postgres", ORM_SOURCE)
		if err != nil {
			fmt.Printf("conn err: %v\n", err)
		}
		hd = hood.New(db, hood.NewPostgres())
	}
}

func HdInsert(b *B) {
	var m *HdModel
	wrapExecute(b, func() {
		initHdDB()
		m = NewHdModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := hd.Save(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func HdInsertMulti(b *B) {
	panic(fmt.Errorf("Problematic bulk insert, too slow"))
	var ms []HdModel
	wrapExecute(b, func() {
		initHdDB()
	})

	for i := 0; i < b.N; i++ {
		ms = make([]HdModel, 100)
		for i := 0; i < 100; i++ {
			ms[i] = *NewHdModel()
		}
		if _, err := hd.SaveAll(&ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func HdUpdate(b *B) {
	var m *HdModel
	wrapExecute(b, func() {
		initHdDB()
		m = NewHdModel()
		if _, err := hd.Save(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	for i := 0; i < b.N; i++ {
		if _, err := hd.Save(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func HdRead(b *B) {
	var m *HdModel
	wrapExecute(b, func() {
		initHdDB()
		m = NewHdModel()
		if _, err := hd.Save(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	//github.com/eaigner/hood/base.go:50 should be patched to `fieldValue.SetString(string(driverValue.Elem().String()))`
	for i := 0; i < b.N; i++ {
		var models []HdModel
		if err := hd.Where("id", "=", m.Id).Find(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}

}

func HdReadSlice(b *B) {
	var m *HdModel
	wrapExecute(b, func() {
		initHdDB()
		m = NewHdModel()
		for i := 0; i < b.L; i++ {
			m.Id = 0
			if _, err := hd.Save(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []HdModel
		if err := hd.Where("id", ">", 0).Limit(b.L).Find(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}

}
