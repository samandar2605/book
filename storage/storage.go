package storage

import (
	"github.com/book/storage/postgres"
	"github.com/book/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Book() repo.RepoBook
}

type StoragePg struct {
	Db       *sqlx.DB
	BookRepo repo.RepoBook
}

// NewStoragePg
func NewStoragePg(db *sqlx.DB) *StoragePg {
	return &StoragePg{
		Db:       db,
		BookRepo: postgres.NewBookRepo(db),
	}
}

func (s StoragePg) Book() repo.RepoBook {
	return s.BookRepo
}
