package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/malailiyati/beginnerBackend/internal/handlers"
)

func InitRouter(db *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/register", func(c *gin.Context) { handlers.RegisterHandler(c, db) })
	r.POST("/login", func(c *gin.Context) { handlers.LoginHandler(c, db) })
	r.PATCH("/users/:email", func(c *gin.Context) { handlers.PatchUserHandler(c, db) })
	r.GET("/users", func(c *gin.Context) { handlers.GetAllUsersHandler(c, db) })

	return r
}
