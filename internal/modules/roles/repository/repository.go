package repository

import (
	"go-minimal/internal/modules/roles/model"

	"github.com/jackc/pgx/v5"
)

type RoleRepositoryI interface {
	GetAll() ([]model.RoleResponse, error)
	Create(role model.RoleResponse) (model.RoleResponse, error)
	GetByRoleID(roleID int) (model.RoleResponse, error)
}

type RoleRepository struct {
	db *pgx.Conn
}

func NewRoleRepository(db *pgx.Conn) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}
