package utils

import (
	"erp/api/request"
	"erp/infrastructure"

	"github.com/jackc/pgx/v4"
)

func ErrNoRows(err error) bool {
	return err == pgx.ErrNoRows
}

func MustHaveDb(db *infrastructure.Database) {
	if db == nil {
		panic("Database engine is null")
	}
}

type QueryPaginationBuilder[E any] struct {
	db *infrastructure.Database
}

func QueryPagination[E any](db *infrastructure.Database, o request.PageOptions, data *[]E) *QueryPaginationBuilder[E] {
	MustHaveDb(db)
	copyDB := &infrastructure.Database{}
	*copyDB = *db
	q := &QueryPaginationBuilder[E]{
		db: copyDB,
	}
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	offset := (o.Page - 1) * o.Limit

	q.db.DB = q.db.Offset(offset).Limit(o.Limit).Find(data)
	return q
}

func (q *QueryPaginationBuilder[E]) Count(total *int64) *QueryPaginationBuilder[E] {
	q.db.Count(total)
	return q
}

func (q *QueryPaginationBuilder[E]) Error() error {
	return q.db.Error
}
