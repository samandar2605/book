package postgres

import (
	"database/sql"
	"fmt"
	"log"

	models "github.com/book/models"
)

type GetBooksQueryParam struct {
	Author string
	Title  string
	Page   int32
	Limit  int32
}

type GetBooksResult struct {
	Books []models.Book  `json:"blogs"`
	Count int            `json:"count"`
}

func (b *DBManager) CreateBook(book *models.Book) (models.Book, error) {
	var kitob models.Book

	tx, err := b.db.Begin()
	if err != nil {
		fmt.Printf("Error at createing transaction %v",err)
		return models.Book{}, err
	}

	query := `
		INSERT INTO books(
			title,
			author_name,
			price,
			amount
		) values ($1,$2,$3,$4)
		RETURNING id,title,author_name,price,amount,created_at
	`
	result := tx.QueryRow(
		query,
		book.Title,
		book.AuthorName,
		book.Price,
		book.Amount,
	)
	if err = result.Scan(
		&kitob.Id,
		&kitob.Title,
		&kitob.AuthorName,
		&kitob.Price,
		&kitob.Amount,
		&kitob.CreatedAt,
	); err != nil {
		tx.Rollback()
		log.Fatalf("error scan'da %v",err.Error())
		return models.Book{}, err
	}
	tx.Commit()
	return kitob, nil
}

func (b DBManager) GetBookById(id int) (models.Book, error) {
	var kitob models.Book
	query := `
		SELECT 
		id,
		title,
		author_name,
		price,
		amount,
		created_at
	FROM books
	WHERE id=$1
	`

	result := b.db.QueryRow(
		query,
		id,
	)
	if err := result.Scan(
		&kitob.Id,
		&kitob.Title,
		&kitob.AuthorName,
		&kitob.Price,
		&kitob.Amount,
		&kitob.CreatedAt,
	); err != nil {
		return models.Book{}, err
	}

	return kitob, nil
}

func (b DBManager) GetBookAll(params GetBooksQueryParam) (GetBooksResult, error) {
	result := GetBooksResult{
		Books: make([]models.Book, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := " WHERE true "
	if params.Author != "" {
		filter += " AND author ilike '%" + params.Author + "%' "
	}

	if params.Title != "" {
		filter += " AND title ilike '%" + params.Title + "%' "
	}

	query := `
		SELECT 
		id,
		title,
		author_name,
		price,
		amount,
		created_at
	FROM books
	` + filter + `
	ORDER BY created_at desc
	` + limit

	rows, err := b.db.Query(query)
	if err != nil {
		return GetBooksResult{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var kitob models.Book
		if err := rows.Scan(
			&kitob.Id,
			&kitob.Title,
			&kitob.AuthorName,
			&kitob.Price,
			&kitob.Amount,
			&kitob.CreatedAt,
		); err != nil {
			return GetBooksResult{}, err
		}
		result.Books = append(result.Books, kitob)

	}

	return result, nil
}

func (b DBManager) UpdateBook(book models.Book) (models.Book, error) {
	var kitob models.Book
	tx, err := b.db.Begin()
	if err != nil {
		return models.Book{}, err
	}
	query := `
		update books set 
			title=$1,
			author_name=$2,
			price=$3,
			amount=$4
		where id=$5
		RETURNING id,title,author_name,price,amount,created_at
	`
	result := tx.QueryRow(
		query,
		book.Title,
		book.AuthorName,
		book.Price,
		book.Amount,
		book.Id,
	)

	if err = result.Scan(
		&kitob.Id,
		&kitob.Title,
		&kitob.AuthorName,
		&kitob.Price,
		&kitob.Amount,
		&kitob.CreatedAt,
	); err != nil {
		tx.Rollback()
		return models.Book{}, err
	}

	tx.Commit()
	return kitob, nil
}

func (b DBManager) DeleteBook(id int) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}
	res, err := tx.Exec("delete from books where id=$1", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}
	tx.Commit()
	return nil
}
