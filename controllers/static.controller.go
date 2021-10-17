package controllers

import (
	"net/http"
	"yofio/avanzado/controllers/dto/output"

	"github.com/gin-gonic/gin"
	// "jesus.tn79/ficha_electronica/models"
)

func (con Controller) Static(c *gin.Context) {
	result, err := con.repo.Static()
	if err != nil {
		c.JSON(http.StatusInternalServerError, output.ResponseError{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
