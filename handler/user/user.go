package user

import (
	"net/http"
	"pku-class/market/data"
	eh "pku-class/market/error-handler"
	"pku-class/market/jwt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LoginPostHandler(c *gin.Context) {
	type dataStruct struct {
		User struct {
			Username string
			Password string
		}
	}
	var reqData dataStruct
	var user data.User
	var err error

	authorizeType := c.GetHeader("AuthorizeType")
	if authorizeType == "token" {
		user.Username = c.GetHeader("username")
	} else {
		err = c.BindJSON(&reqData)
		eh.ErrorHandler(err, "LoginPostHandler", "")
		user.Username = reqData.User.Username
		user.Password = reqData.User.Password
	}
	result := authorize(&user, authorizeType)
	if result {
		user.Token, err = jwt.GenerateToken(user.Username, user.ID, user.Role)
		eh.ErrorHandler(err, "generate token wrong", "")
		data.PackingToGinJson(c, user)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": false,
		})
	}
	return
}

func OrderGetHandler(c *gin.Context) {
	username := c.GetHeader("username")
	role, err := strconv.Atoi(c.GetHeader("role"))
	eh.ErrorHandler(err, "OrderGetHandler", "")
	orders := getOrders(username, role)
	data.PackingToGinJson(c, orders)
	return
}

func OrderDeletePostHandler(c *gin.Context) {
	type dataStruct struct {
		ID uint
	}
	var reqData dataStruct
	err := c.BindJSON(&reqData)
	eh.ErrorHandler(err, "OrderDeletePostHandler", "")
	result := deleteOrder(reqData.ID)
	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func ChargePostHandler(c *gin.Context) {
	type dataStruct struct {
		Amount int
	}
	var reqData dataStruct
	err := c.BindJSON(&reqData)
	eh.ErrorHandler(err, "ChargePostHandler", "")
	id, err := strconv.Atoi(c.GetHeader("id"))
	eh.ErrorHandler(err, "ChargePostHandler", "")
	result, balance := charge(uint(id), reqData.Amount)
	c.JSON(http.StatusOK, gin.H{
		"result":  result,
		"balance": balance,
	})
}
