package role_usecase

import (
	"backend-api/models/user/dto"
	"backend-api/repository/role_repository"
)

type RoleUsecase interface {
	GetAllRole() dto.Response
}

type roleUsecase struct {
	roleRepo role_repository.RoleRepository
}

func GetRoleUsecase(role role_repository.RoleRepository) RoleUsecase {
	return &roleUsecase{
		roleRepo: role,
	}
}
