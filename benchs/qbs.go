package benchs

import (
	"fmt"

	"github.com/coocood/qbs"
	"database/sql"
)

var qo *qbs.Qbs
type QModel struct {
	Id      int64 `qbs:"pk" orm:"auto" gorm:"primary_key" db:"id"`
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

func NewQModel() *QModel {
	m := new(QModel)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}


func initQDB() {

	sqls := []string{
		`DROP TABLE IF EXISTS q_model;`,
		`CREATE TABLE q_model (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT q_model_pkey PRIMARY KEY (id)
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
	st := NewSuite("qbs")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000 * ORM_MULTI, 0, QbsInsert)
		st.AddBenchmark("BulkInsert 100 row", 500 * ORM_MULTI, 0, QbsInsertMulti)
		st.AddBenchmark("Update", 2000 * ORM_MULTI, 0, QbsUpdate)
		st.AddBenchmark("Read", 4000 * ORM_MULTI, 0, QbsRead)
		st.AddBenchmark("MultiRead limit 1000", 2000 * ORM_MULTI, 1000, QbsReadSlice)

		qbs.Register("postgres", ORM_SOURCE, "q_model", qbs.NewPostgres())
		qbs.ChangePoolSize(ORM_MAX_IDLE)
		qbs.SetConnectionLimit(ORM_MAX_CONN, true)
		var err error
		qo, err = qbs.GetQbs()
		if err != nil {
			fmt.Printf("conn err: %v\n", err)
		}
	}
}

func QbsInsert(b *B) {
	var m *QModel
	wrapExecute(b, func() {
		initQDB()
		m = NewQModel()
	})
	defer qo.Close()
	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := qo.Save(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func QbsInsertMulti(b *B) {
	panic(fmt.Errorf("Don't support bulk insert, err driver: bad connection"))
	var ms []*QModel
	wrapExecute(b, func() {
		initQDB()

		ms = make([]*QModel, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewQModel())
		}
	})
	for i := 0; i < b.N; i++ {
		if err := qo.BulkInsert(ms); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func QbsUpdate(b *B) {
	var m *QModel
	wrapExecute(b, func() {
		initQDB()
		m = NewQModel()
		qo.Save(m)
	})

	for i := 0; i < b.N; i++ {
		if _, err := qo.Save(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func QbsRead(b *B) {
	var m *QModel
	wrapExecute(b, func() {
		initQDB()
		m = NewQModel()
		qo.Save(m)
	})

	for i := 0; i < b.N; i++ {
		if err := qo.Find(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func QbsReadSlice(b *B) {
	var m *QModel
	wrapExecute(b, func() {
		initQDB()
		m = NewQModel()
		for i := 0; i < b.L; i++ {
			m.Id = 0
			if _, err := qo.Save(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})
	for i := 0; i < b.N; i++ {
		var models []*QModel
		if err := qo.Where("id > ?", 0).Limit(b.L).FindAll(&models); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}