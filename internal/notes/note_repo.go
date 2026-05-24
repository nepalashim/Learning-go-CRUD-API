package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func (r *Repo) ListNotes(ctx context.Context) ([]Note, error) {

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{} //empty filter to get all notes
	// find returns a cursor that we can iterate over to get the results---> over matching documents in the collection.
	cursor, err := r.coll.Find(opCtx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to list notes: %w", err)
	}

	//cursor must be closed after the use
	// avoids memory leaks and ensures that database resources are released properly. By deferring the close operation, we guarantee that the cursor will be closed regardless of how the function exits, whether it returns successfully or encounters an error.
	defer cursor.Close(opCtx)

	var notes []Note

	if err := cursor.All(opCtx, &notes); err != nil {
		return nil, fmt.Errorf("failed to decode notes: %w", err)
	}

	return notes, nil

}

func (r *Repo) GetByID(ctx context.Context, id bson.ObjectID) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	var note Note

	err := r.coll.FindOne(opCtx, filter, options.FindOne()).Decode(&note)
	if err != nil {
		return Note{}, fmt.Errorf("failed to get note by ID: %w", err)
	}

	return note, nil
}

func (r *Repo) UpdateNote(ctx context.Context, id bson.ObjectID, update UpdateNoteRequest) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}

	updateFields := bson.M{
		"updatedAt": time.Now().UTC(),
	}

	if update.Title != nil {
		updateFields["title"] = *update.Title
	}
	if update.Content != nil {
		updateFields["content"] = *update.Content
	}
	if update.Pinned != nil {
		updateFields["pinned"] = *update.Pinned
	}

	updateDoc := bson.M{"$set": updateFields}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var updated Note
	err := r.coll.FindOneAndUpdate(opCtx, filter, updateDoc, opts).Decode(&updated)
	if err != nil {
		return Note{}, fmt.Errorf("failed to update note: %w", err)
	}
	return updated, nil
}
