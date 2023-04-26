package postgres

import (
	"movie/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type actorRepo struct {
	db *pgxpool.Pool
}

func NewActorRepo(db *pgxpool.Pool) storage.ActorRepoI {
	return &actorRepo{
		db: db,
	}
}
