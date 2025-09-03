package main

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var users = map[string]string{
	"user1@gmail.com": "Password!",
	"user2@gmail.com": "Password@",
}

func main() {
	router := gin.Default()

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, Response{
			Message: "Rute salah",
			Status:  "Rute tidak ditemukan",
		})
	})
	// login
	router.POST("/login", func(ctx *gin.Context) {
		body := Login{}
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}
		if err := ValidateLogin(body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		//cek email suda pernah di register atau belum
		if pass, ok := users[body.Email]; !ok {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "email tidak ditemukan",
			})
			return
		} else if pass != body.Password {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "password salah",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"body":    body,
		})
	})

	// register
	router.POST("/register", func(ctx *gin.Context) {
		body := Register{}
		if err := ctx.ShouldBind(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"success": false,
			})
			return
		}
		if err := ValidateRegister(body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		// cek email sudah terdaftar belum
		if _, exis := users[body.Email]; exis {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "email sudah terdaftar",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"body":    body,
		})
	})
	router.Run(":8080")
}

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// validasi login
func ValidateLogin(login Login) error {
	re, err := regexp.Compile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	if err != nil {
		return err
	}
	if login.Email == "" {
		return errors.New("email tidak boleh kosong")
	}
	if isMatched := re.Match([]byte(login.Email)); !isMatched {
		return errors.New("format email tidak valid")
	}
	// validasi password
	if login.Password == "" {
		return errors.New("password tidak boleh kosong")
	}
	if len(login.Password) < 8 {
		return errors.New("password minimal 8 karakter")
	}
	reLower := regexp.MustCompile("[a-z]")
	reUpper := regexp.MustCompile("[A-Z]")
	reSpecial := regexp.MustCompile("[^a-zA-Z0-9]")
	if !reLower.MatchString(login.Password) {
		return errors.New("password harus mengandung huruf kecil")
	}
	if !reUpper.MatchString(login.Password) {
		return errors.New("password harus mengandung huruf besar")
	}
	if !reSpecial.MatchString(login.Password) {
		return errors.New("password harus mengandung karakter spesial")
	}
	return nil
}

// validasi register
func ValidateRegister(register Register) error {
	re, err := regexp.Compile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	if err != nil {
		return err
	}
	if register.Email == "" {
		return errors.New("email tidak boleh kosong")
	}
	if isMatched := re.Match([]byte(register.Email)); !isMatched {
		return errors.New("format email tidak valid")
	}
	// validasi password
	if register.Password == "" {
		return errors.New("password tidak boleh kosong")
	}
	if len(register.Password) < 8 {
		return errors.New("password minimal 8 karakter")
	}
	reLower := regexp.MustCompile("[a-z]")
	reUpper := regexp.MustCompile("[A-Z]")
	reSpecial := regexp.MustCompile("[^a-zA-Z0-9]")
	if !reLower.MatchString(register.Password) {
		return errors.New("password harus mengandung huruf kecil")
	}
	if !reUpper.MatchString(register.Password) {
		return errors.New("password harus mengandung huruf besar")
	}
	if !reSpecial.MatchString(register.Password) {
		return errors.New("password harus mengandung karakter spesial")
	}
	return nil
}
