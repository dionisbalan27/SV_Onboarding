package role_usecase

import (
	helpers "backend-api/helpers/helpers_user"
	"backend-api/models/user/dto"
)

func (role *roleUsecase) GetAllRole() dto.Response {
	roles, err := role.roleRepo.GetAllRole()

	if err != nil {
		return helpers.ResponseError("Data not found", err.Error())
	}

	return helpers.ResponseSuccess("ok", nil, roles)
}
