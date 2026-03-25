package repository

import (
	"context"
	"database/sql"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"
	"time"

	"github.com/google/uuid"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, email, public_id, status, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User

	for rows.Next() {
		var user domain.User

		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PublicId, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	// user.ID = uuid.New().String()
	query := "INSERT INTO users(name,email,password,public_id,status,created_at,updated_at) VALUES(?,?,?,?,?,?,?)"
	_, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.PublicId, user.Status, time.Now(), time.Now())
	return err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id,name,email,password,public_id,status,created_at,updated_at FROM users WHERE email=?",
		email,
	)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.PublicId, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id,name,email,password,public_id,status,created_at,updated_at FROM users WHERE id=?",
		id,
	)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.PublicId, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	query := "UPDATE users SET name=?, email=?, password=?, updated_at=? WHERE id=?"
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, time.Now(), user.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return response.ErrNotFound
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM users WHERE id=?"
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return response.ErrNotFound
	}
	return nil
}

func (r *userRepository) FindByPublicID(ctx context.Context, PublicId uuid.UUID) (*domain.User, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id,name,email,password,public_id,status,created_at,updated_at FROM users WHERE public_id=?",
		PublicId,
	)

	var user domain.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.PublicId, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetRolesAndPermissions(ctx context.Context, userID int64) ([]string, []string, error) {
	var roles []string
	var permissions []string

	roleQuery := `
		SELECT r.name FROM roles r
		JOIN model_has_roles mr ON r.id = mr.role_id
		WHERE mr.model_id = ?`

	rows, err := r.db.QueryContext(ctx, roleQuery, userID)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var roleName string
			if err := rows.Scan(&roleName); err == nil {
				roles = append(roles, roleName)
			}
		}
	}

	permQuery := `
		SELECT DISTINCT p.name FROM permissions p
		JOIN role_has_permissions rp ON p.id = rp.permission_id
		JOIN model_has_roles mr ON rp.role_id = mr.role_id
		WHERE mr.model_id = ?
		UNION
		SELECT p.name FROM permissions p
		JOIN model_has_permissions mp ON p.id = mp.permission_id
		WHERE mp.model_id = ?`

	permRows, err := r.db.QueryContext(ctx, permQuery, userID, userID)
	if err == nil {
		defer permRows.Close()
		for permRows.Next() {
			var permName string
			if err := permRows.Scan(&permName); err == nil {
				permissions = append(permissions, permName)
			}
		}
	}

	return roles, permissions, nil
}

func (r *userRepository) AssignRoleToUser(ctx context.Context, userID int64, roleID int64) error {
	query := "INSERT INTO model_has_roles(role_id, model_type, model_id) VALUES (?, 'User', ?)"
	_, err := r.db.ExecContext(ctx, query, roleID, userID)
	return err
}

func (r *userRepository) AssignPermissionToUser(ctx context.Context, userID int64, permissionID int64) error {
	query := "INSERT INTO model_has_permissions(permission_id, model_type, model_id) VALUES (?, 'User', ?)"
	_, err := r.db.ExecContext(ctx, query, permissionID, userID)
	return err
}
