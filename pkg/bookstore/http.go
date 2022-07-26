package bookstore

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookStore struct {
	service Service
}

func NewController(service Service) *BookStore {
	return &BookStore{
		service: service,
	}
}

func (bs BookStore) RegisterRoutes(r *gin.Engine) {
	r.GET("/getbook", bs.CreateBookStore)
}

func (bs BookStore) CreateBookStore(c *gin.Context) {
	getBookResponse, err := bs.service.getBook(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, getBookResponse)
}
