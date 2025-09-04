package repositories

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/malailiyati/beginnerBackend/internal/models"
)

// ambil user by email
func GetUserByEmail(ctx context.Context, db *pgxpool.Pool, email string) (*models.User, error) {
	const q = `SELECT id, email, password, role, created_at, updated_at 
	           FROM users WHERE email = $1`
	var u models.User
	if err := db.QueryRow(ctx, q, email).
		Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

// cek apakah email sudah ada
func EmailExists(ctx context.Context, db *pgxpool.Pool, email string) (bool, error) {
	const q = `SELECT 1 FROM users WHERE email = $1`
	var dummy int
	if err := db.QueryRow(ctx, q, email).Scan(&dummy); err != nil {
		// kalau tidak ketemu, Scan error → berarti email belum ada
		return false, nil
	}
	return true, nil
}

// insert user baru
func CreateUser(ctx context.Context, db *pgxpool.Pool, email, password string) (*models.User, error) {
	const q = `
		INSERT INTO users (email, password, role, created_at, updated_at)
		VALUES ($1, $2, 'user', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id, email, password, role, created_at, updated_at
	`
	var u models.User
	if err := db.QueryRow(ctx, q, email, password).
		Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

// update sebagian (PATCH) by email
func PatchUserByEmail(ctx context.Context, db *pgxpool.Pool, currentEmail string, body models.UpdateUser) (*models.User, error) {
	const q = `
		UPDATE users SET
			email      = COALESCE($1, email),
			password   = COALESCE($2, password),
			updated_at = CURRENT_TIMESTAMP
		WHERE email = $3
		RETURNING id, email, password, role, created_at, updated_at
	`
	var u models.User
	if err := db.QueryRow(ctx, q, body.Email, body.Password, currentEmail).
		Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}

// ambil semua user (opsional, buat debug)
func GetAllUsers(ctx context.Context, db *pgxpool.Pool) ([]models.User, error) {
	const q = `SELECT id, email, password, role, created_at, updated_at FROM users ORDER BY id`
	rows, err := db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Email, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
