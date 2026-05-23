package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

//Repository is data access layer for notes

//Here we are going to keep all our mongodb queries and operations related to notes

type Repo struct {
	coll *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		coll: db.Collection("notes"),
	}
}

func (r *Repo) CreateNote(ctx context.Context, note Note) (Note, error) {

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	_, err := r.coll.InsertOne(opCtx, note)
	if err != nil {
		return Note{}, fmt.Errorf("failed to create note: %w", err)
	}
	return note, nil
}
