package repository

import (
	"context"
	"database/sql"
	"errors"
	"gpt/internal/domain"
	"time"
)

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) domain.RoleRepository {
	return &roleRepository{db}
}

func (r *roleRepository) FindByName(ctx context.Context, name string) (*domain.Roles, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id,name,guard_name FROM roles WHERE name=?",
		name,
	)

	var roles domain.Roles
	err := row.Scan(&roles.ID, &roles.Name, &roles.GuardName)
	if err != nil {
		return nil, err
	}

	return &roles, nil
}

func (r *roleRepository) FindByID(ctx context.Context, id int64) (*domain.Roles, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id,name,guard_name FROM roles WHERE id=?",
		id,
	)

	var roles domain.Roles
	err := row.Scan(&roles.ID, &roles.Name, &roles.GuardName)
	if err != nil {
		return nil, err
	}

	return &roles, nil
}

func (r *roleRepository) GetAll(ctx context.Context) ([]*domain.Roles, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, guard_name,created_at,updated_at FROM roles order by id asc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*domain.Roles

	for rows.Next() {
		var role domain.Roles

		if err := rows.Scan(&role.ID, &role.Name, &role.GuardName, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}

		roles = append(roles, &role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) Create(ctx context.Context, roles *domain.Roles) error {
	query := "INSERT INTO roles(name,guard_name,created_at,updated_at) VALUES(?,?,?,?)"
	_, err := r.db.ExecContext(ctx, query,
		roles.Name, roles.GuardName, time.Now(), time.Now())
	return err
}

func (r *roleRepository) Update(ctx context.Context, roles *domain.Roles) error {
	query := `
		UPDATE roles
		SET name = ?, guard_name = ?
		WHERE id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		roles.Name,
		roles.GuardName,
		roles.ID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("role not found")
	}

	return nil
}

func (r *roleRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM roles WHERE id = ?`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("role not found")
	}

	return nil
}

func (r *roleRepository) AssignPermissionToRole(ctx context.Context, roleID int64, permissionID int64) error {
	query := "INSERT INTO role_has_permissions(permission_id, role_id) VALUES (?, ?)"
	_, err := r.db.ExecContext(ctx, query, permissionID, roleID)
	return err
}
