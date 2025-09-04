// internal/utils/response.go
package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

func OKWithMessage(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": msg,
		"data":    data,
	})
}

func Created(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    data,
	})
}

func BadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"success": false,
		"error":   msg,
	})
}

func Unauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"success": false,
		"error":   msg,
	})
}

func NotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"error":   msg,
	})
}

func Conflict(c *gin.Context, msg string) {
	c.JSON(http.StatusConflict, gin.H{
		"success": false,
		"error":   msg,
	})
}

// In production, sebaiknya JANGAN mengirim detail error mentah ke client.
// Untuk dev, kamu bisa expose "detail". Di prod, hapus field "detail".
func ServerError(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error":   msg,
		"detail":  err.Error(), // <- untuk dev. Sembunyikan di production.
	})
}
