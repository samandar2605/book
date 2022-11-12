package repo

import (
	models "github.com/book/models"
	"github.com/book/storage/postgres"
)

type RepoBook interface {
	CreateBook(book *models.Book) (models.Book, error)
	GetBookById(id int) (models.Book, error)
	GetBookAll(params postgres.GetBooksQueryParam) (postgres.GetBooksResult, error)
	UpdateBook(book models.Book) (models.Book, error)
	DeleteBook(id int) error
}
