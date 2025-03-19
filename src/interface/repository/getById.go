package repository

import (
	ctx "context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/challenge_prueba_biblioteca/src/domain/model"
)

func (r *bookRepository) GetById(ID int) (*model.Book, error) {

	var book model.Book
	err := r.collection.FindOne(ctx.TODO(), bson.M{"_id": ID}).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
