package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"stock/service"
)

// FindGroupBarangByID ...
func FindGroupBarangByID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 16)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"result": "failed convert to integer"})
	}
	var gbService service.GroupBarangService
	groupBarang := gbService.FindByID(int32(id))
	c.JSON(http.StatusOK, gin.H{"result": groupBarang})
}

// FindGroupBarangByCode ...
func FindGroupBarangByCode(c *gin.Context) {

	code := c.Param("code")
	var gbService service.GroupBarangService

	groupBarang := gbService.FindByCode(code)
	c.JSON(http.StatusOK, gin.H{"result": groupBarang})

}
