package utils

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/malailiyati/beginnerBackend/internal/models"
)

// ---------- VALIDATOR ----------
var (
	emailRe   = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	lowerRe   = regexp.MustCompile(`[a-z]`)
	upperRe   = regexp.MustCompile(`[A-Z]`)
	specialRe = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

func ValidateEmail(s string) error {
	if s == "" {
		return errors.New("email tidak boleh kosong")
	}
	if !emailRe.MatchString(s) {
		return errors.New("format email tidak valid")
	}
	return nil
}

func ValidatePassword(p string) error {
	if p == "" {
		return errors.New("password tidak boleh kosong")
	}
	if len(p) < 8 {
		return errors.New("password minimal 8 karakter")
	}
	if !lowerRe.MatchString(p) {
		return errors.New("password harus mengandung huruf kecil")
	}
	if !upperRe.MatchString(p) {
		return errors.New("password harus mengandung huruf besar")
	}
	if !specialRe.MatchString(p) {
		return errors.New("password harus mengandung karakter spesial")
	}
	return nil
}

func ValidateLogin(b models.Login) error {
	if err := ValidateEmail(b.Email); err != nil {
		return err
	}
	return ValidatePassword(b.Password)
}

func ValidateRegister(b models.Register) error {
	if err := ValidateEmail(b.Email); err != nil {
		return err
	}
	return ValidatePassword(b.Password)
}

func ValidatePatch(b models.UpdateUser) error {
	if b.Email == nil && b.Password == nil {
		return errors.New("tidak ada field yang diupdate")
	}
	if b.Email != nil {
		if err := ValidateEmail(*b.Email); err != nil {
			return err
		}
	}
	if b.Password != nil {
		if err := ValidatePassword(*b.Password); err != nil {
			return err
		}
	}
	return nil
}

// ---------- RESPONSE HELPER ----------
func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func OKWithMessage(c *gin.Context, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{"success": true, "message": msg, "data": data})
}

func BadRequest(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": msg})
}

func Unauthorized(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": msg})
}

func NotFound(c *gin.Context, msg string) {
	c.JSON(http.StatusNotFound, gin.H{"success": false, "error": msg})
}

func Conflict(c *gin.Context, msg string) {
	c.JSON(http.StatusConflict, gin.H{"success": false, "error": msg})
}

func ServerError(c *gin.Context, msg string, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"error":   msg,
		"detail":  err.Error(),
	})
}
