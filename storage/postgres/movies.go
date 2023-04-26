package postgres

import (
	"context"
	"movie/storage"

	"movie/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type movieRepo struct {
	db *pgxpool.Pool
}

func NewMovieRepo(db *pgxpool.Pool) storage.MovieRepoI {
	return &movieRepo{
		db: db,
	}
}

func (m *movieRepo) Create(ctx context.Context, req models.CreateMovie) (string, error) {
	query := `insert into movies
			  (id, title, image_url)
			  values ($1, $2, $3)`

	_, err := m.db.Exec(ctx, query,
		"",
		req.Title,
		req.ImageURL)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (m *movieRepo) Update(id string, req models.UpdateMovie) error {
	return nil
}

func (m *movieRepo) Delete(id string) error {
	return nil
}

func (m *movieRepo) Get(id string) (models.Movie, error) {
	return models.Movie{}, nil
}

func (m *movieRepo) GetAll() ([]models.Movie, error) {
	return []models.Movie{}, nil
}
