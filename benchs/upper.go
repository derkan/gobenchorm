package benchs

import (
	"fmt"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
)

var sess db.Session

type VModel struct {
	Id      int64  `db:"id,omitempty"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

func NewVModel() *VModel {
	m := new(VModel)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "https://github.com/derkan/gobenchorm"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

func init() {
	st := NewSuite("upper")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, 0, UpperIOInsert)
		st.AddBenchmark("BulkInsert 100 row", 500*ORM_MULTI, 0, UpperIOInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, 0, UpperIOUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, 0, UpperIORead)
		st.AddBenchmark("MultiRead limit 1000", 2000*ORM_MULTI, 1000, UpperIOReadSlice)
		settings, err := postgresql.ParseURL(ORM_SOURCE)
		checkErr(err)
		db, err := postgresql.Open(settings)
		checkErr(err)
		sess = db
	}
}

func UpperIOInsert(b *B) {
	var m *VModel
	wrapExecute(b, func() {
		initDB()
		m = NewVModel()
	})
	col := sess.Collection("models")
	var err error
	for i := 0; i < b.N; i++ {
		m.Id = 0
		err = col.InsertReturning(m)
		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func UpperIOInsertMulti(b *B) {
	wrapExecute(b, func() {
		initDB()
	})
	var err error
	for i := 0; i < b.N; i++ {
		v := NewVModel()
		batch := sess.SQL().InsertInto("models").Columns("name", "title", "fax", "web", "age", "right", "counter").Batch(100)
		go func() {
			for i := 0; i < 100; i++ {
				batch.Values(v.Name, v.Title, v.Fax, v.Web, v.Age, v.Right, v.Counter)
			}
			batch.Done()
		}()
		if batch.Wait() != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func UpperIOUpdate(b *B) {
	var m *VModel
	col := sess.Collection("models")
	wrapExecute(b, func() {
		initDB()
		m = NewVModel()
		if err := col.InsertReturning(m); err != nil {
			b.FailNow()
		}
	})
	for i := 0; i < b.N; i++ {
		if err := col.UpdateReturning(m); err != nil {
			fmt.Printf("update err: %v\n", err)
			b.FailNow()
		}
	}
}

func UpperIORead(b *B) {
	var m *VModel
	col := sess.Collection("models")
	wrapExecute(b, func() {
		initDB()
		m = NewVModel()
		if err := col.InsertReturning(m); err != nil {
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := sess.SQL().SelectFrom("models").Where("id = ?", m.Id).One(m); err != nil {
			b.FailNow()
		}
	}
}

func UpperIOReadSlice(b *B) {
	var m *VModel
	wrapExecute(b, func() {
		initDB()
		m = NewVModel()
		col := sess.Collection("models")
		for i := 0; i < b.L; i++ {
			m.Id = 0
			if _, err := col.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})
	for i := 0; i < b.N; i++ {
		var models []VModel
		if err := sess.SQL().SelectFrom("models").Where("id > ?", 0).Limit(b.L).All(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
