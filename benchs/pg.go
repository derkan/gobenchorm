package benchs

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

var pgdb *pg.DB

func init() {
	st := NewSuite("pg")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, 0, PgInsert)
		st.AddBenchmark("BulkInsert 100 row", 500*ORM_MULTI, 0, PgInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, 0, PgUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, 0, PgRead)
		st.AddBenchmark("MultiRead limit 1000", 2000*ORM_MULTI, 1000, PgReadSlice)

		pgdb = pg.Connect(&pg.Options{
			Addr:     "localhost:5432",
			User:     "bench",
			Password: "pass",
			Database: "benchdb",
		})
	}
}

func PgInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0
		if _, err := pgdb.Model(m).Insert(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgInsertMulti(b *B) {
	var ms []*Model
	wrapExecute(b, func() {
		initDB()
	})

	for i := 0; i < b.N; i++ {
		ms = make([]*Model, 100)
		for i := 0; i < 100; i++ {
			ms[i] = NewModel()
		}
		if _, err := pgdb.Model(&ms).Insert(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		m.Id = 1
		if _, err := pgdb.Model(m).Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if _, err := pgdb.Model(m).Where("id = ?", 1).Update(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if _, err := pgdb.Model(m).Insert(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})
	m = NewModel()
	for i := 0; i < b.N; i++ {
		if err := pgdb.Model(m).Select(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func PgReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < b.L; i++ {
			m.Id = 0
			if _, err := pgdb.Model(m).Insert(); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []Model
		if err := pgdb.Model(&models).Where("id > ?", 0).Limit(b.L).Select(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
