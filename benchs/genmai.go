package benchs

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/naoina/genmai"
)

var genmaidb *genmai.DB

type GModel struct {
	Id      int `db:"pk"`
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

func (*GModel) TableName() string {
	return "models"
}

func NewGModel() *GModel {
	m := new(GModel)
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
	st := NewSuite("genmai")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, 0, GenmaiInsert)
		st.AddBenchmark("BulkInsert 100 row", 500*ORM_MULTI, 0, GenmaiInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, 0, GenmaiUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, 0, GenmaiRead)
		st.AddBenchmark("MultiRead limit 1000", 2000*ORM_MULTI, 1000, GenmaiReadSlice)

		db, err := genmai.New(&genmai.PostgresDialect{}, ORM_SOURCE)
		checkErr(err)
		genmaidb = db
	}
}

func GenmaiInsert(b *B) {
	var m *GModel
	wrapExecute(b, func() {
		initDB()
		m = NewGModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = i
		if _, err := genmaidb.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GenmaiInsertMulti(b *B) {
	var ms []*GModel
	wrapExecute(b, func() {
		initDB()
	})

	for i := 0; i < b.N; i++ {
		ms = make([]*GModel, 100)
		for i := 0; i < 100; i++ {
			ms[i] = NewGModel()
		}
		if _, err := genmaidb.Insert(&ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GenmaiUpdate(b *B) {
	var m *GModel
	wrapExecute(b, func() {
		initDB()
		m = NewGModel()
		if _, err := genmaidb.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	for i := 0; i < b.N; i++ {
		if _, err := genmaidb.Update(m); err != nil {
			fmt.Printf("update err: %v\n", err)
			b.FailNow()
		}
	}
}

func GenmaiRead(b *B) {
	var m *GModel
	wrapExecute(b, func() {
		initDB()
		m = NewGModel()
		if _, err := genmaidb.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	for i := 0; i < b.N; i++ {
		var results []GModel
		if err := genmaidb.Select(&results); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func GenmaiReadSlice(b *B) {
	wrapExecute(b, func() {
		initDB()
		for i := 0; i < b.N; i++ {
			var ms []*GModel
			ms = make([]*GModel, 0, b.L)
			for i := 0; i < b.L; i++ {
				ms = append(ms, NewGModel())
			}
			if _, err := genmaidb.Insert(&ms); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})
	for i := 0; i < b.N; i++ {
		var models []GModel
		if err := genmaidb.Select(&models, genmaidb.Where("id", ">", 0).Limit(b.L)); err != nil {
			fmt.Printf("sel %v\n", err)
			b.FailNow()
		}
	}
}
