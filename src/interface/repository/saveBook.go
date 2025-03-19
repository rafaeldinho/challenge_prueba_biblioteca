package repository

import (
	ctx "context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/challenge_prueba_biblioteca/src/domain/model"
	"github.com/challenge_prueba_biblioteca/src/shared"
)

func (r *bookRepository) Save(book *model.Book) (error, bool) {
	if _, err := r.collection.InsertOne(ctx.Background(), book); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return fmt.Errorf(shared.BookAlreadyExists, book.ID), true
		}
		return err, false
	}
	return nil, false
}
