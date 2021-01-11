package handler

import (
	"unsafe"
)

func FlashPostHandler(body []byte) {
	type dataStruct struct {
		Item struct {
			ID     uint
			Number int
			UserID uint
		}
	}
	var reqData *dataStruct
	reqData = *(**dataStruct)(unsafe.Pointer(&body))
	//log.Println(reqData)
	flash(reqData.Item.UserID, reqData.Item.ID, reqData.Item.Number)
}
