package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"stock/models"
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

// AddGroupBarangByCode ...
func AddGroupBarangByCode(c *gin.Context) {

	var bodyReq models.BodyReqGroupBarang

	bodyRaw, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(bodyRaw, &bodyReq)

	// bodyReq := &models.BodyReqGroupBarang{}
	// groupBarang := c.Bind(bodyReq) // data is left unchanged because c.Request.Body has been used up.
	// fmt.Println(bodyReq)

	var groupBarang models.GroupBarang

	groupBarang.Code = bodyReq.Code
	groupBarang.Name = bodyReq.Name
	groupBarang.UpdatedBy = "admin"

	var gbService service.GroupBarangService

	res := gbService.Add(groupBarang)
	c.JSON(http.StatusOK, gin.H{"result": res})

}
