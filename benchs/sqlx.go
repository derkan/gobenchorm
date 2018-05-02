package benchs

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var sqlxdb *sqlx.DB

func init() {
	st := NewSuite("sqlx")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000 * ORM_MULTI, 0, SqlxInsert)
		st.AddBenchmark("BulkInsert 100 row", 500 * ORM_MULTI, 0, SqlxInsertMulti)
		st.AddBenchmark("Update", 2000 * ORM_MULTI, 0, SqlxUpdate)
		st.AddBenchmark("Read", 4000 * ORM_MULTI, 0, SqlxRead)
		st.AddBenchmark("MultiRead limit 1000", 2000 * ORM_MULTI, 1000, SqlxReadSlice)

		db, err := sqlx.Connect("postgres", ORM_SOURCE)
		checkErr(err)
		sqlxdb = db
	}
}

func SqlxInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})
	var err error
	for i := 0; i < b.N; i++ {
		if err = sqlxdb.QueryRowx(`INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
			m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter).StructScan(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func SqlxInsertMulti(b *B) {
	panic(fmt.Errorf("benchmark not implemeted yet - https://github.com/jmoiron/sqlx/issues/134"))
}

func SqlxUpdate(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		if err := sqlxdb.QueryRowx(`INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
			m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter).StructScan(m); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		sqlxdb.MustExec(`UPDATE models SET name = $1, title = $2, fax = $3, web = $4, age = $5, "right" = $6, counter = $7 WHERE id = $8`,
			m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter, m.Id)
	}
}

func SqlxRead(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		sqlxdb.MustExec(`INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES ($1, $2, $3, $4, $5, $6, $7)`, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
	})
	for i := 0; i < b.N; i++ {
		m := []Model{}
		if err := sqlxdb.Select(&m, "SELECT * FROM models"); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

func SqlxReadSlice(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
		for i := 0; i < b.L; i++ {
			sqlxdb.MustExec(`INSERT INTO models (name, title, fax, web, age, "right", counter) VALUES ($1, $2, $3, $4, $5, $6, $7)`, m.Name, m.Title, m.Fax, m.Web, m.Age, m.Right, m.Counter)
		}
	})

	for i := 0; i < b.N; i++ {
		var models []*Model
		if err := sqlxdb.Select(&models, "SELECT * FROM models WHERE id > $1 LIMIT $2", 0, b.L); err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}
