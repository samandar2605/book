package api

import (
	"github.com/book/storage/repo"

	_ "net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/book/api/book"
	_ "github.com/book/api/docs" // for swagger
)

func NewServer(data repo.RepoBook) *gin.Engine {
	r := gin.Default()

	h := book.Handler{
		Storage: data,
	}

	r.GET("/books/:id", h.GetBook)
	r.GET("/books", h.GetBookAll)
	r.POST("/books", h.CreateBook)
	r.DELETE("/books/:id", h.DeleteBook)
	r.PUT("/books/:id", h.UpdateBook)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
