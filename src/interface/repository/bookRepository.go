package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/challenge_prueba_biblioteca/src/domain/model"
	"github.com/challenge_prueba_biblioteca/src/shared"
)

type bookRepository struct {
	collection *mongo.Collection
}

type BookRepository interface {
	Save(*model.Book) (error, bool)
	GetAll() ([]model.Book, error)
	GetById(int) (*model.Book, error)
	FetchBooks(*model.BookQuery) (*model.CurrencyResponse, error)
}

func NewBookRepository(client *mongo.Client) *bookRepository {
	return &bookRepository{collection: client.Database(shared.DatabaseCollection).Collection(shared.MutantCollection)}
}
