package role_repository

import (
	"backend-api/models/user/entity"

	"gorm.io/gorm"
)

type RoleRepository interface {
	GetAllRole() ([]entity.Role, error)
}

type roleRepository struct {
	mysqlConn *gorm.DB
}

func GetRoleRepository(mysqlConnection *gorm.DB) RoleRepository {
	return &roleRepository{
		mysqlConn: mysqlConnection,
	}
}
