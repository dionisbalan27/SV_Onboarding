package role_repository

import (
	"backend-api/models/user/entity"
)

func (repo *roleRepository) GetAllRole() ([]entity.Role, error) {
	roles := []entity.Role{}
	err := repo.mysqlConn.Model(&entity.Role{}).Scan(&roles).Error

	if err != nil {
		return nil, err
	}

	return roles, nil

}
