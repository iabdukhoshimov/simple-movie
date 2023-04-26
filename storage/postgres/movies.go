package postgres

import (
	"movie/storage"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/i.abdukhoshimov/golang_bootcamp/golang-psql/storage"
)

type movieRepo struct {
	db *pgxpool.Pool
}

func NewMovieRepo(db *pgxpool.Pool) storage.MovieRepoI {
	return &movieRepo{
		db: db,
	}
}
