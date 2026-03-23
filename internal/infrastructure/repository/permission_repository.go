package repository

import (
	"context"
	"database/sql"
	"gpt/internal/domain"
	"time"
)

type PermissionRepository struct {
	db *sql.DB
}

func NewPermissionRepository(db *sql.DB) domain.PermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) GetAll(ctx context.Context) ([]domain.Permission, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, guard_name FROM permissions order by id asc")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []domain.Permission
	for rows.Next() {
		var p domain.Permission
		if err := rows.Scan(&p.ID, &p.Name, &p.GuardName); err != nil {
			return nil, err
		}
		permissions = append(permissions, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *PermissionRepository) GetByID(ctx context.Context, id string) (*domain.Permission, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, guard_name FROM permissions WHERE id = ?", id)

	var p domain.Permission
	if err := row.Scan(&p.ID, &p.Name, &p.GuardName); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PermissionRepository) GetByName(ctx context.Context, name string) (*domain.Permission, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, guard_name FROM permissions WHERE name = ?", name)

	var p domain.Permission
	if err := row.Scan(&p.ID, &p.Name, &p.GuardName); err != nil {
		return nil, err
	}

	return &p, nil
}

func (r *PermissionRepository) Create(ctx context.Context, permission *domain.Permission) error {
	// permission.ID = uuid.New().String()
	_, err := r.db.ExecContext(ctx, "INSERT INTO permissions (name, guard_name, created_at, updated_at) VALUES (?, ?, ?, ?)", permission.Name, permission.GuardName, time.Now(), time.Now())
	return err
}

func (r *PermissionRepository) Update(ctx context.Context, permission *domain.Permission) error {
	_, err := r.db.ExecContext(ctx, "UPDATE permissions SET name = ?, guard_name = ?, updated_at = ? WHERE id = ?", permission.Name, permission.GuardName, time.Now(), permission.ID)
	return err
}

func (r *PermissionRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM permissions WHERE id = ?", id)
	return err
}
