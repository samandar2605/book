package book

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/book/models"
	"github.com/book/storage/postgres"
	"github.com/book/storage/repo"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Storage repo.RepoBook
}

// // @Summary create
// // @Description this functions
// // @Tags Book
// // @Accept json
// // @Produce json
// // @Param 	book body models.Book true "Book"
// // @Success 200 {object} models.Book
// // @Failure 400  {object} models.ResponseError
// // @Router /books [post]
// func (h *Handler) CreateBook(ctx *gin.Context) {
// 	var b models.Book
// 	err := ctx.ShouldBindJSON(&b)
// 	fmt.Println(b)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	book, err := h.Storage.CreateBook(&b)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "failed to create book",
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, book)
// }

// @Summary 	Get_book_by_id
// @Description Get book by id
// @Tags 		Book
// @Accept 		json
// @Produce 	json
// @Param 		id path int true "ID"
// @Success 	200 {object} models.Book
// @Failure 	400 {object} models.ResponseError
// @Router 		/Book/{id} [get]

func (h *Handler) GetBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})

		return
	}

	book, err := h.Storage.GetBookById(id)
	fmt.Println(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get method",
		})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

// @Summary Get books
// @Description Get books
// @Tags Book
// @Accept json
// @Produce json
// @Param limit query int true "Limit"
// @Param page query int true "Page"
// @Param author query string false "Author"
// @Param title query string false "Title"
// @Success 200 "successfully"
// @Failure 500 {object} models.ResponseError
// @Router /books [get]

func (h *Handler) GetBookAll(ctx *gin.Context) {
	queryParams, err := validateGetbooksQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	resp, err := h.Storage.GetBookAll(queryParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func validateGetbooksQuery(ctx *gin.Context) (postgres.GetBooksQueryParam, error) {
	var (
		limit int64 = 10
		page  int64 = 1
		err   error
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			return postgres.GetBooksQueryParam{}, err
		}
	}

	if ctx.Query("page") != "" {
		page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			return postgres.GetBooksQueryParam{}, err
		}
	}

	return postgres.GetBooksQueryParam{
		Limit:  int32(limit),
		Page:   int32(page),
		Author: ctx.Query("author"),
		Title:  ctx.Query("title"),
	}, nil
}

// @Summary Update a book
// @Description Update a books
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param book body models.CreateBookRequest true "Book"
// @Success 200 {object} models.Book
// @Failure 500 {object} models.ResponseError
// @Router /books/{id} [put]

func (h *Handler) UpdateBook(ctx *gin.Context) {
	var b models.Book

	err := ctx.ShouldBindJSON(&b)
	fmt.Println(b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		fmt.Println("Error at update method'da")
		return
	}

	b.Id = id
	book, err := h.Storage.UpdateBook(b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create book",
		})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// @Summary Delete a book
// @Description Delete a book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ResponseError
// @Router /books/{id} [delete]

func (h *Handler) DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to convert",
		})
		return
	}

	err = h.Storage.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to Delete method",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "successful delete method",
	})
}
