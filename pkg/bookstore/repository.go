package bookstore

import (
	"context"

	"github.com/mussabaheen/orm/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
)

type repository struct {
	db db.MongoClient
}

type Repository interface {
	getBook(ctx context.Context) (getBookResponse, error)
}

func NewRepository(db db.MongoClient) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) getBook(ctx context.Context) (getBookResponse, error) {
	_, err := r.db.InsertOne(ctx, createBookRequest{
		ID:     "123",
		Name:   "Mastering GO!!",
		Rating: 5,
	})
	if err != nil {
		return getBookResponse{}, err
	}
	var book getBookResponse
	err = r.db.FindOne(ctx, bson.M{"id": "123"}, &book)
	if err != nil {
		return getBookResponse{}, err
	}
	return book, nil
}
