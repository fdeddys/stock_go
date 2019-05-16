package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindBarangByID ...
func FindBarangByID(c *gin.Context) {

	idBarang := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"result": "barang " + idBarang + "pong"})

}
