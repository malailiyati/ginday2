package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/malailiyati/beginnerBackend/internal/models"
	"github.com/malailiyati/beginnerBackend/internal/repositories"
	"github.com/malailiyati/beginnerBackend/internal/utils"
	"github.com/malailiyati/beginnerBackend/internal/validators"
)

// POST /login
func LoginHandler(c *gin.Context) {
	var body models.Login
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := validators.ValidateLogin(body); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	u, err := repositories.GetUserByEmail(c.Request.Context(), body.Email)
	if err != nil {
		utils.Unauthorized(c, "email tidak ditemukan")
		return
	}
	// DEMO: plain text; production: hash + bcrypt compare
	if u.Password != body.Password {
		utils.Unauthorized(c, "password salah")
		return
	}

	utils.OK(c, gin.H{"id": u.ID, "email": u.Email, "role": u.Role})
}

// POST /register
func RegisterHandler(c *gin.Context) {
	var body models.Register
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := validators.ValidateRegister(body); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if ok, _ := repositories.EmailExists(c.Request.Context(), body.Email); ok {
		utils.Unauthorized(c, "email sudah terdaftar")
		return
	}

	u, err := repositories.CreateUser(c.Request.Context(), body.Email, body.Password)
	if err != nil {
		utils.ServerError(c, "gagal membuat user", err)
		return
	}
	utils.OK(c, gin.H{"id": u.ID, "email": u.Email, "role": u.Role})
}

// PATCH /users/:email  (partial update by current email)
func PatchUserHandler(c *gin.Context) {
	currentEmail := c.Param("email")

	if _, err := repositories.GetUserByEmail(c.Request.Context(), currentEmail); err != nil {
		utils.NotFound(c, "user tidak ditemukan")
		return
	}

	var body models.UpdateUser
	if err := c.ShouldBindJSON(&body); err != nil {
		utils.BadRequest(c, "invalid request body: "+err.Error())
		return
	}
	if err := validators.ValidatePatch(body); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	// optional: jika ubah email, cek konflik
	if body.Email != nil && *body.Email != currentEmail {
		if ok, _ := repositories.EmailExists(c.Request.Context(), *body.Email); ok {
			utils.Conflict(c, "email baru sudah terdaftar")
			return
		}
	}

	u, err := repositories.PatchUserByEmail(c.Request.Context(), currentEmail, body)
	if err != nil {
		utils.ServerError(c, "gagal update user", err)
		return
	}
	utils.OKWithMessage(c, "user updated (partial)", gin.H{"id": u.ID, "email": u.Email, "role": u.Role})
}
