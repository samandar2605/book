package testing

import (
	"fmt"

	models "github.com/book/models"
	"github.com/book/storage/postgres"

	"github.com/bxcodec/faker/v4"
)

var dbManger *postgres.DBManager

func CreateBook() *models.Book {
	fmt.Println("FUnnga kirish >>>>>>>>>>>>>>>>>>>>")
	var book models.Book

	book.Title=faker.Name()
	book.Author=faker.Name()
	book.Price=12345
	book.Amount=1
	var err error
	book, err = dbManger.CreateBook(&book)
	// require.NoError(t, err)
	// require.NotEmpty(t, book)
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>-----------------<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	fmt.Println(book, err)
	return &book
}

// func deleteBook(id int, t *testing.T) {
// 	err := dbManger.DeleteBook(id)
// 	require.NoError(t, err)
// }

// func TestGetBookById(t *testing.T) {
// 	// b := CreateBook(t)
// 	book, err := dbManger.GetBookById(b.Id)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, book)

// 	deleteBook(book.Id, t)
// }

// func TestCreateBook(t *testing.T) {
// 	// b := CreateBook(t)
// 	fmt.Println(b)
// 	deleteBook(b.Id, t)
// }

// func TestUpdateBook(t *testing.T) {
// 	// b := CreateBook(t)

// 	b.Title = faker.Sentence()
// 	b.AuthorName = faker.Name()
// 	b.Price = 100000
// 	b.Amount = 5452
// 	b.CreatedAt = faker.Timestamp()
// 	Book, err := dbManger.UpdateBook(*b)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, Book)

// 	deleteBook(Book.Id, t)
// }

// func TestDeleteBook(t *testing.T) {
// 	// b := CreateBook(t)

// 	deleteBook(b.Id, t)
// }

// // func TestGetAll(t *testing.T) {
// 	// b := CreateBook(t)

// // 	Books, err := dbManger.GetBookAll()

// // 	require.NoError(t, err)
// // 	require.GreaterOrEqual(t, len(Books.Books), 1)

// // 	deleteBook(b.Id, t)
// // }
