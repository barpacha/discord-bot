package store

import "database/sql"

const (
	dialect     = "postgres"
	serverTable = "server"
)

type Store struct {
	db     *sql.DB
	server *ServerRepository
}

func New(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Server() *ServerRepository {
	if s.server == nil {
		s.server = &ServerRepository{store: s}
	}
	return s.server
}
