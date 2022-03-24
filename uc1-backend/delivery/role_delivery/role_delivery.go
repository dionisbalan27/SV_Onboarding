package role_delivery

import (
	"backend-api/usecase/role_usecase"

	"github.com/gin-gonic/gin"
)

type RoleDelivery interface {
	GetAllRole(c *gin.Context)
}

type roleDelivery struct {
	roleUsecase role_usecase.RoleUsecase
}

func GetRoleDelivery(role role_usecase.RoleUsecase) RoleDelivery {
	return &roleDelivery{
		roleUsecase: role,
	}
}
