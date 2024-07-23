package repository

import (
	"database/sql"
)

// Registry represents the specification of this pkg
type Registry interface {
}

// New returns an implementation instance which satisfying Registry
func New(pgConn *sql.DB) Registry {
	return impl{
		pgConn: pgConn,
	}
}

type impl struct {
	pgConn *sql.DB
}
