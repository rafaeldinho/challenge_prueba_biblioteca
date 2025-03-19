package repository

import (
	ctx "context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/challenge_prueba_biblioteca/src/domain/model"
)

func (r *bookRepository) GetAll() ([]model.Book, error) {

	var results []model.Book

	rstSearch, err := r.collection.Find(ctx.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	for rstSearch.Next(ctx.Background()) {
		var book model.Book
		if err := rstSearch.Decode(&book); err != nil {
			return nil, err
		}
		results = append(results, book)
	}

	if err := rstSearch.Err(); err != nil {
		return nil, err
	}

	rstSearch.Close(ctx.Background())

	return results, nil

}
