package menu

import (
	"pku-class/market/data"
	eh "pku-class/market/error-handler"

	"github.com/gin-gonic/gin"
)

func PostHandler(c *gin.Context) {
	type dataStruct struct {
		Theme string
	}
	var reqData dataStruct
	err := c.BindJSON(&reqData)
	eh.ErrorHandler(err, "ChargePostHandler", "")
	list := getMenu(reqData.Theme)
	data.PackingToGinJson(c, list)
	return
}
