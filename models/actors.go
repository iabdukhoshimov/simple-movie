package models

import "time"

type Actors struct {
	ID        int64      `json:"id"`
	FirstName string     `json:"first_name"`
	ImageURL  string     `json:"image_url"`
	CreatedAT *time.Time `json:"created_at"`
	UpdatedAT *time.Time `json:"updated_at"`
}

type CreateActor struct {
	FirstName string `json:"first_name"`
	ImageURL  string `json:"image_url"`
}

type UpdateActor struct {
	FirstName string `json:"first_name"`
	ImageURL  string `json:"image_url"`
}

type MovieActor struct {
	ID        int64      `json:"id"`
	MovieID   int64      `json:"movie_id"`
	ActorID   int64      `json:"actor_id"`
	CreatedAT *time.Time `json:"created_at"`
	UpdatedAT *time.Time `json:"updated_at"`
}
