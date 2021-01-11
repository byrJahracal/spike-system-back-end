package manage

import (
	"log"
	"pku-class/market/data"
	db "pku-class/market/database"
	eh "pku-class/market/error-handler"
	"time"
)

func updateCommodity(commodity *data.Commodity) (result bool) {
	result = false
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := tx.Save(commodity).Error
	eh.ErrorHandler(err, "updateCommodity", "")
	err = tx.Commit().Error
	eh.ErrorHandler(err, "updateCommodity", "")
	log.Println(time.Now())
	log.Println("商品信息已被修改！")
	log.Print(commodity)
	log.Println("")
	if err != nil {
		return false
	}
	return true
}

func deleteCommodity(id uint) (result bool) {
	result = false
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	commodity := data.Commodity{ID: id}
	err := tx.Delete(&commodity).Error
	eh.ErrorHandler(err, "deleteCommodity", "")
	err = tx.Commit().Error
	eh.ErrorHandler(err, "deleteCommodity", "")
	if err != nil {
		return false
	}
	return true
}

func createCommodity(commodity *data.Commodity) (result bool) {
	result = false
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	err := tx.Create(commodity).Error
	eh.ErrorHandler(err, "createCommodity", "")
	err = tx.Commit().Error
	eh.ErrorHandler(err, "createCommodity", "")
	if err != nil {
		return false
	}
	return true
}
