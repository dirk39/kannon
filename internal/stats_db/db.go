// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sq

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.countQueryStatsStmt, err = db.PrepareContext(ctx, countQueryStats); err != nil {
		return nil, fmt.Errorf("error preparing query CountQueryStats: %w", err)
	}
	if q.insertPreparedStmt, err = db.PrepareContext(ctx, insertPrepared); err != nil {
		return nil, fmt.Errorf("error preparing query InsertPrepared: %w", err)
	}
	if q.insertStatStmt, err = db.PrepareContext(ctx, insertStat); err != nil {
		return nil, fmt.Errorf("error preparing query InsertStat: %w", err)
	}
	if q.queryStatsStmt, err = db.PrepareContext(ctx, queryStats); err != nil {
		return nil, fmt.Errorf("error preparing query QueryStats: %w", err)
	}
	if q.queryStatsTimelineStmt, err = db.PrepareContext(ctx, queryStatsTimeline); err != nil {
		return nil, fmt.Errorf("error preparing query QueryStatsTimeline: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.countQueryStatsStmt != nil {
		if cerr := q.countQueryStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countQueryStatsStmt: %w", cerr)
		}
	}
	if q.insertPreparedStmt != nil {
		if cerr := q.insertPreparedStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertPreparedStmt: %w", cerr)
		}
	}
	if q.insertStatStmt != nil {
		if cerr := q.insertStatStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing insertStatStmt: %w", cerr)
		}
	}
	if q.queryStatsStmt != nil {
		if cerr := q.queryStatsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing queryStatsStmt: %w", cerr)
		}
	}
	if q.queryStatsTimelineStmt != nil {
		if cerr := q.queryStatsTimelineStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing queryStatsTimelineStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                     DBTX
	tx                     *sql.Tx
	countQueryStatsStmt    *sql.Stmt
	insertPreparedStmt     *sql.Stmt
	insertStatStmt         *sql.Stmt
	queryStatsStmt         *sql.Stmt
	queryStatsTimelineStmt *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                     tx,
		tx:                     tx,
		countQueryStatsStmt:    q.countQueryStatsStmt,
		insertPreparedStmt:     q.insertPreparedStmt,
		insertStatStmt:         q.insertStatStmt,
		queryStatsStmt:         q.queryStatsStmt,
		queryStatsTimelineStmt: q.queryStatsTimelineStmt,
	}
}
