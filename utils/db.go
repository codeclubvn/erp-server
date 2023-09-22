package utils

import (
	"erp/api/request"
	"erp/infrastructure"
	"fmt"

	"github.com/jackc/pgx/v4"
	"gorm.io/gorm"
)

func ErrNoRows(err error) bool {
	return err == pgx.ErrNoRows
}

func MustHaveDb(db *infrastructure.Database) {
	if db == nil {
		panic("Database engine is null")
	}
}

func MustHaveGormDb(db *gorm.DB) {
	if db == nil {
		panic("Gorm Database engine is null")
	}
}

type QueryPaginationBuilder[E any] struct {
	db *gorm.DB
}

func QueryPagination[E any](query *gorm.DB, o request.PageOptions, data *[]*E) *QueryPaginationBuilder[E] {
	MustHaveGormDb(query)
	q := &QueryPaginationBuilder[E]{
		db: query,
	}
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	offset := (o.Page - 1) * o.Limit

	q.db = q.db.Debug().Offset(int(offset)).Limit(int(o.Limit)).Find(&data)

	fmt.Println(data)
	return q
}

func (q *QueryPaginationBuilder[E]) Count(total *int64) *QueryPaginationBuilder[E] {
	q.db.Count(total)
	return q
}

func (q *QueryPaginationBuilder[E]) Error() error {
	return q.db.Error
}
