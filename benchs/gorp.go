package benchs

import (
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v2"
)

var dbmap *gorp.DbMap

type GOModel struct {
	Id      int    `db:"id,primarykey,autoincrement"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

func NewGOModel() *GOModel {
	m := new(GOModel)
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
	st := NewSuite("gorp")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, 0, GorpInsert)
		st.AddBenchmark("BulkInsert 100 row", 500*ORM_MULTI, 0, GorpInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, 0, GorpUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, 0, GorpRead)
		st.AddBenchmark("MultiRead limit 1000", 2000*ORM_MULTI, 1000, GorpReadSlice)

		db, err := sql.Open("postgres", ORM_SOURCE)
		checkErr(err)
		d := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
		dbmap = d
		dbmap.AddTableWithName(GOModel{}, "models").SetKeys(true, "id")
	}
}

func GorpInsert(b *B) {
	var m *GOModel
	wrapExecute(b, func() {
		initDB()
		m = NewGOModel()
	})
	for i := 0; i < b.N; i++ {
		m.Id = 0
		if err := dbmap.Insert(m); err!=nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GorpInsertMulti(b *B) {
	panic(fmt.Errorf("Problematic bulk insert, too slow"))
	var ms [] interface{}
	wrapExecute(b, func() {
		initDB()
	})

	for i := 0; i < b.N; i++ {
		ms = make([]interface{}, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewGOModel())
		}
		if err := dbmap.Insert(ms...); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GorpUpdate(b *B) {
	var m *GOModel
	wrapExecute(b, func() {
		initDB()
		m = NewGOModel()
		if err := dbmap.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := dbmap.Update(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GorpRead(b *B) {
	var m *GOModel
	wrapExecute(b, func() {
		initDB()
		m = NewGOModel()
		if err := dbmap.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := dbmap.SelectOne(m, "SELECT * FROM models"); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GorpReadSlice(b *B) {
	var m *GOModel
	wrapExecute(b, func() {
		initDB()
		m = NewGOModel()
		for i := 0; i < b.L; i++ {
			m.Id = 0
			if err := dbmap.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*GOModel
		if _, err := dbmap.Select(&models, "SELECT * FROM models WHERE id>:id", map[string]interface{}{
			"id":0,
		}); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
