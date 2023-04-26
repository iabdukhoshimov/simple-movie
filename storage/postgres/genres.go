package postgres

import (
	"movie/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type genreRepo struct {
	db *pgxpool.Pool
}

func NewGenreRepo(db *pgxpool.Pool) storage.GenresRepoI {
	return &genreRepo{
		db: db,
	}
}
