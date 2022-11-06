package db

import (
	"context"
	"time"
)

const addMigration = `-- name: AddMigration :exec
INSERT INTO migration (name) VALUES (?)
`

func (q *Queries) AddMigration(ctx context.Context, name string) error {
	_, err := q.db.ExecContext(ctx, addMigration, name)
	return err
}

const getMigrationByName = `-- name: GetMigrationByName :one
SELECT id, executed_at, name
FROM migration
WHERE name = ?
`

func (q *Queries) GetMigrationByName(ctx context.Context, name string) (Migration, error) {
	row := q.db.QueryRowContext(ctx, getMigrationByName, name)
	var i Migration
	err := row.Scan(&i.ID, &i.ExecutedAt, &i.Name)
	return i, err
}

type Migration struct {
	ID         int32
	ExecutedAt time.Time
	Name       string
}
