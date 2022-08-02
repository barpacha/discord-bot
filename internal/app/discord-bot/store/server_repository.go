package store

import (
	"context"
	"database/sql"
	"discord-bot/internal/app/discord-bot/model"
	"github.com/doug-martin/goqu/v9"
)

type ServerRepository struct {
	store *Store
}

func (s *ServerRepository) Create(ctx context.Context, id uint) error {
	dialect := goqu.Dialect(dialect)
	query, _, _ := dialect.From(serverTable).Insert().Rows(
		goqu.Record{
			"id": id,
		},
	).ToSQL()

	if _, err := s.store.db.ExecContext(
		ctx,
		query,
	); err != nil {
		return err
	}
	return nil
}

func (s *ServerRepository) Delete(ctx context.Context, id uint) error {
	dialect := goqu.Dialect(dialect)
	query, _, _ := dialect.From(serverTable).Delete().Where(goqu.Ex{"id": id}).ToSQL()

	if _, err := s.store.db.ExecContext(
		ctx,
		query,
	); err != nil {
		return err
	}
	return nil
}

func (s *ServerRepository) List(ctx context.Context) ([]model.Server, error) {
	dialect := goqu.Dialect(dialect)
	query, _, _ := dialect.From(serverTable).Select(
		"id",
		"active",
	).ToSQL()

	rows, err := s.store.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var servers []model.Server

	for rows.Next() {
		var server model.Server
		err = rows.Scan(
			&server.Id,
			&server.Active,
		)
		if err != nil {
			return nil, err
		}
	}
	return servers, nil
}

func (s *ServerRepository) Get(ctx context.Context, id uint) (model.Server, error) {
	dialect := goqu.Dialect(dialect)
	query, _, _ := dialect.From(serverTable).Insert().Rows(
		goqu.Record{
			"id": id,
		},
	).ToSQL()

	var server model.Server

	if err := s.store.db.QueryRowContext(
		ctx,
		query,
	).Scan(
		&server.Id,
		&server.Active,
	); err != nil {
		return model.Server{}, err
	}
	return server, nil
}
