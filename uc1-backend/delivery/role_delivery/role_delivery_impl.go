package role_delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (res *roleDelivery) GetAllRole(c *gin.Context) {
	response := res.roleUsecase.GetAllRole()

	if response.Status != "ok" {
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)
}
