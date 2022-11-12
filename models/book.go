package models

type Book struct {
	Id         int     `json:"id" db:"id"`
	Title      string  `json:"title" db:"title"`
	AuthorName string  `json:"author_name" db:"author_name"`
	Price      float64 `json:"price" db:"price"`
	Amount     int     `json:"amount" db:"amount"`
	CreatedAt  string  `json:"created_at" db:"created_at"`
}

type ResponseError struct {
	Message string `json:"message"`
}

type ResponseOK struct {
	Message string `json:"message"`
}

type CreateBookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Price      float64 `json:"price"`
	Amount      int `json:"amount"`
}
