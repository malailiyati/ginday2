package routers

import (
	"github.com/malailiyati/beginnerBackend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.POST("/register", handlers.RegisterHandler)
	r.POST("/login", handlers.LoginHandler)
	r.PATCH("/users/:email", handlers.PatchUserHandler) // partial update by email
}
