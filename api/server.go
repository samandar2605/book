package api

import (
	"github.com/book/storage/repo"

	"github.com/gin-gonic/gin"
	_ "net/http"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/book/api/book"
	_ "github.com/book/api/docs" // for swagger
)

// @title           Swagger for book api
// @version         1.0
// @description     This is a book service api.
// @host      		localhost:8080

func NewServer(data repo.RepoBook) *gin.Engine {
	r := gin.Default()

	h := book.Handler{
		Storage: data,
	}

	r.GET("/book/:id", h.GetBook)
	r.GET("/book", h.GetBookAll)
	r.POST("/book", h.CreateBook)
	r.DELETE("/book/:id", h.DeleteBook)
	r.PUT("/book/:id", h.UpdateBook)

	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
