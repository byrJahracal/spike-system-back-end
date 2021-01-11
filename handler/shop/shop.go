package shop

import (
	"net/http"
	"pku-class/market/data"
	eh "pku-class/market/error-handler"
	"pku-class/market/rabbitmq"
	"unsafe"

	"strconv"

	"github.com/gin-gonic/gin"
)

func PostHandler(c *gin.Context) {
	type dataStruct struct {
		Theme string
	}
	var reqData dataStruct
	err := c.BindJSON(&reqData)
	eh.ErrorHandler(err, "PostHandler", "")
	items := getItems(reqData.Theme)
	data.PackingToGinJson(c, items)
	return
}

func ConsumePostHandler(c *gin.Context) {
	type dataStruct struct {
		Item struct {
			ID     uint
			Number int
		}
	}
	var reqData dataStruct
	err := c.BindJSON(&reqData)
	eh.ErrorHandler(err, "ConsumePostHandler", "")
	userID, err := strconv.Atoi(c.GetHeader("id"))
	eh.ErrorHandler(err, "ConsumePostHandler", "")
	result, amount, balance := consume(uint(userID), reqData.Item.ID, reqData.Item.Number)
	c.JSON(http.StatusOK, gin.H{
		"result":  result,
		"amount":  amount,
		"balance": balance,
	})
	return
}

func FlashPostHandler(c *gin.Context) {
	type SliceMock struct {
		addr uintptr
		len  int
		cap  int
	}
	type dataStruct struct {
		Item struct {
			ID     uint
			Number int
			UserID uint
		}
	}
	var reqData dataStruct
	err := c.BindJSON(&reqData)
	eh.ErrorHandler(err, "ConsumePostHandler", "")
	userID, err := strconv.Atoi(c.GetHeader("id"))
	eh.ErrorHandler(err, "ConsumePostHandler", "")
	reqData.Item.UserID = uint(userID)

	len := unsafe.Sizeof(reqData)

	reqBytes := SliceMock{
		addr: uintptr(unsafe.Pointer(&reqData)),
		len:  int(len),
		cap:  int(len),
	}
	data := *(*[]byte)(unsafe.Pointer(&reqBytes))
	rabbitmq.Publish(data)
	c.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}
