package manage

import (
	"net/http"
	"pku-class/market/data"
	eh "pku-class/market/error-handler"

	"github.com/gin-gonic/gin"
)

func UpdatePostHandler(c *gin.Context) {
	type dataStruct struct {
		Item data.Commodity
	}
	var reqData dataStruct
	err := c.BindJSON(&reqData)
	eh.ErrorHandler(err, "UpdatePostHandler", "")
	result := updateCommodity(&reqData.Item)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func DeletePostHandler(c *gin.Context) {
	type dataStruct struct {
		ID uint
	}
	var reqData dataStruct
	err := c.BindJSON(&reqData)
	eh.ErrorHandler(err, "DeletePostHandler", "")
	result := deleteCommodity(reqData.ID)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func AddPostHandler(c *gin.Context) {
	type dataStruct struct {
		Item data.Commodity
	}
	var reqData dataStruct
	err := c.BindJSON(&reqData)
	eh.ErrorHandler(err, "AddPostHandler", "")
	result := createCommodity(&reqData.Item)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
