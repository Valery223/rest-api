package router

import (
	userHandler "learn/rest-api/internal/book/transport"

	"github.com/gin-gonic/gin"
)

func NewRouter(uH *userHandler.BookHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// version
		v1 := api.Group("/v1")
		{
			// Entity
			users := v1.Group("/books")
			{
				users.GET("/:id", uH.GetBook)
			}

		}
	}

	return r
}
