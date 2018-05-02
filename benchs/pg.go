package benchs

import (
	"fmt"
	"github.com/go-pg/pg"
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
		if err := pgdb.Insert(m); err != nil {
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
		ms = make([]*Model, 0, 100)
		for i := 0; i < 100; i++ {
			ms = append(ms, NewModel())
		}
		if err := pgdb.Insert(&ms); err != nil {
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
		if err := pgdb.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := pgdb.Update(m); err != nil {
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
		if err := pgdb.Insert(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := pgdb.Select(m); err != nil {
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
			if err := pgdb.Insert(m); err != nil {
				fmt.Println(err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := pgdb.Model(&models).Where("id > ?", 0).Limit(b.L).Select(); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
