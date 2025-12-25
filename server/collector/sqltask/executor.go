package sqltask

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SQLExecutor struct {
	pool        *sql.DB
	DBType      string
	DSN         string
	ExecTimeout int
}

func newSQLExecutor(dbType, dsn string, execTimeout int) *SQLExecutor {
	return &SQLExecutor{
		DBType:      dbType,
		DSN:         dsn,
		ExecTimeout: execTimeout,
	}
}

func (t *SQLExecutor) exec(querySql string) ([]byte, error) {

	var err error
	t.pool, err = sql.Open(t.DBType, t.DSN)
	if err != nil {
		return nil, err
	}
	defer t.pool.Close()

	t.pool.SetConnMaxLifetime(0)
	t.pool.SetMaxIdleConns(1)
	t.pool.SetMaxOpenConns(1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(t.ExecTimeout)*time.Millisecond)
	defer cancel()

	rows, err := t.pool.QueryContext(ctx, querySql)
	if err != nil {
		return nil, err
	}

	return t.encodeJson(rows)
}

func (t *SQLExecutor) encodeJson(rows *sql.Rows) ([]byte, error) {

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := []map[string]interface{}{}
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		m := make(map[string]interface{})

		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		result = append(result, m)
	}

	return json.Marshal(result)
}
