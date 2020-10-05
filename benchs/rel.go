package benchs

import (
	"context"
	"fmt"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/postgres"
	"github.com/go-rel/rel/where"
)

var repo rel.Repository
var ctx = context.Background()

type RModel struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	Title   string `db:"title"`
	Fax     string `db:"fax"`
	Web     string `db:"web"`
	Age     int    `db:"age"`
	Right   bool   `db:"right"`
	Counter int64  `db:"counter"`
}

func (r RModel) Table() string {
	return "models"
}
func NewRModel() *RModel {
	m := new(RModel)
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
	st := NewSuite("rel")
	st.InitF = func() {
		st.AddBenchmark("Insert", 2000*ORM_MULTI, 0, RelInsert)
		st.AddBenchmark("BulkInsert 100 row", 500*ORM_MULTI, 0, RelInsertMulti)
		st.AddBenchmark("Update", 2000*ORM_MULTI, 0, RelUpdate)
		st.AddBenchmark("Read", 4000*ORM_MULTI, 0, RelRead)
		st.AddBenchmark("MultiRead limit 1000", 2000*ORM_MULTI, 1000, RelReadSlice)
		var err error
		adapter, _ := postgres.Open(ORM_SOURCE)
		repo = rel.New(adapter)
		repo.Instrumentation(func(ctx context.Context, op string, message string) func(err error) {
			return func(err error) {
				if err != nil {
					fmt.Printf("err: %v\n", err)
				}
			}
		})

		if err != nil {
			fmt.Printf("conn err: %v\n", err)
		}
	}
}

func RelInsert(b *B) {
	var m *RModel
	wrapExecute(b, func() {
		initDB()
		m = NewRModel()
	})
	for i := 0; i < b.N; i++ {
		m.Id = 0
		if err := repo.Insert(ctx, m); err != nil {
			fmt.Printf("insert err: %v\n", err)
			b.FailNow()
		}
	}
}

func RelInsertMulti(b *B) {
	var ms []RModel
	wrapExecute(b, func() {
		initDB()
	})
	for i := 0; i < b.N; i++ {
		ms = make([]RModel, 100)
		for i := 0; i < 100; i++ {
			ms[i] = *NewRModel()
		}
		if err := repo.InsertAll(ctx, &ms); err != nil {
			fmt.Printf("bulkinsert err: %v\n", err)
			b.FailNow()
		}
	}
}

func RelUpdate(b *B) {
	var m *RModel
	wrapExecute(b, func() {
		initDB()
		m = NewRModel()
		if err := repo.Insert(ctx, m); err != nil {
			fmt.Printf("insert before update err: %v\n", err)
			b.FailNow()
		}
	})
	for i := 0; i < b.N; i++ {
		if err := repo.Update(ctx, m); err != nil {
			fmt.Printf("update err: %v\n", err)
			b.FailNow()
		}
	}
}

func RelRead(b *B) {
	var m *RModel
	wrapExecute(b, func() {
		initDB()
		m = NewRModel()
		if err := repo.Insert(ctx, m); err != nil {
			fmt.Printf("insert before read err: %v\n", err)
			b.FailNow()
		}
	})

	for i := 0; i < b.N; i++ {
		if err := repo.Find(ctx, m); err != nil {
			fmt.Printf("read err: %v\n", err)
			b.FailNow()
		}
	}
}

func RelReadSlice(b *B) {
	var m *RModel
	wrapExecute(b, func() {
		initDB()
		m = NewRModel()
		for i := 0; i < b.L; i++ {
			m.Id = 0
			if err := repo.Insert(ctx, m); err != nil {
				fmt.Printf("insert before readslice err: %v\n", err)
				b.FailNow()
			}
		}
	})

	for i := 0; i < b.N; i++ {
		var models []RModel
		if err := repo.FindAll(ctx, &models, where.Gt("id", 1), rel.Limit(b.L)); err != nil {
			fmt.Printf("slice err: %v\n", err)
			b.FailNow()
		}
	}
}
