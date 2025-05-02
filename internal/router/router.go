package router

import (
	bookHandler "learn/rest-api/internal/book/transport"
	userHandler "learn/rest-api/internal/user/transport"

	"github.com/gin-gonic/gin"
)

func NewRouter(bH *bookHandler.BookHandler, uH *userHandler.UserHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		// version
		v1 := api.Group("/v1")
		{
			// Entity
			users := v1.Group("/users")
			{
				users.GET("/:id", uH.GetUser)
				users.POST("", uH.PostUser)
			}

			books := v1.Group("/books")
			{
				books.GET("/:id", bH.GetBook)
				books.POST("", bH.PostBook)
			}

		}
	}

	return r
}
