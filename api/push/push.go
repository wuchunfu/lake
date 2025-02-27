package push

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/merico-dev/lake/api/shared"
	"github.com/merico-dev/lake/services"
)

/*
	POST /push/:tableName
	[
		{
			"id": 1,
			"sha": "osidjfoawehfwh08"
		}
	]
*/

func Post(c *gin.Context) {
	var err error
	tableName := c.Param("tableName")
	var rowsToInsert []map[string]interface{}
	err = c.ShouldBindJSON(&rowsToInsert)
	if err != nil {
		shared.ApiOutputError(c, err, http.StatusBadRequest)
		return
	}
	rowsAffected, err := services.InsertRow(tableName, rowsToInsert)
	if err != nil {
		shared.ApiOutputError(c, err, http.StatusBadRequest)
		return
	}
	shared.ApiOutputSuccess(c, gin.H{"rowsAffected": rowsAffected}, http.StatusOK)
}
