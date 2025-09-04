package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/malailiyati/beginnerBackend/internal/models"
	"github.com/malailiyati/beginnerBackend/internal/repositories"
	"github.com/malailiyati/beginnerBackend/internal/utils"
)

func LoginHandler(c *gin.Context, db *pgxpool.Pool) {
	var body models.Login
	if err := c.ShouldBind(&body); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := utils.ValidateLogin(body); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	u, err := repositories.GetUserByEmail(c.Request.Context(), db, body.Email)
	if err != nil {
		utils.Unauthorized(c, "email tidak ditemukan")
		return
	}
	if u.Password != body.Password {
		utils.Unauthorized(c, "password salah")
		return
	}

	utils.OK(c, gin.H{"id": u.ID, "email": u.Email, "role": u.Role})
}

func RegisterHandler(c *gin.Context, db *pgxpool.Pool) {
	var body models.Register
	if err := c.ShouldBind(&body); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := utils.ValidateRegister(body); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if ok, _ := repositories.EmailExists(c.Request.Context(), db, body.Email); ok {
		utils.Unauthorized(c, "email sudah terdaftar")
		return
	}

	u, err := repositories.CreateUser(c.Request.Context(), db, body.Email, body.Password)
	if err != nil {
		utils.ServerError(c, "gagal membuat user", err)
		return
	}
	utils.OK(c, u)
}

func PatchUserHandler(c *gin.Context, db *pgxpool.Pool) {
	currentEmail := c.Param("email")
	if _, err := repositories.GetUserByEmail(c.Request.Context(), db, currentEmail); err != nil {
		utils.NotFound(c, "user tidak ditemukan")
		return
	}

	var body models.UpdateUser
	if err := c.ShouldBind(&body); err != nil {
		utils.BadRequest(c, "invalid request body: "+err.Error())
		return
	}
	if err := utils.ValidatePatch(body); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if body.Email != nil && *body.Email != currentEmail {
		if ok, _ := repositories.EmailExists(c.Request.Context(), db, *body.Email); ok {
			utils.Conflict(c, "email baru sudah terdaftar")
			return
		}
	}

	u, err := repositories.PatchUserByEmail(c.Request.Context(), db, currentEmail, body)
	if err != nil {
		utils.ServerError(c, "gagal update user", err)
		return
	}
	utils.OKWithMessage(c, "user updated (partial)", u)
}

func GetAllUsersHandler(c *gin.Context, db *pgxpool.Pool) {
	users, err := repositories.GetAllUsers(c.Request.Context(), db)
	if err != nil {
		utils.ServerError(c, "gagal mengambil data users", err)
		return
	}
	utils.OK(c, users)
}
