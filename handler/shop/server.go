package shop

import (
	"pku-class/market/data"
	db "pku-class/market/database"
	eh "pku-class/market/error-handler"

	"gorm.io/gorm"
)

func getItems(theme string) (items []data.Commodity) {
	err := db.DB.Where("type = ?", theme).Find(&items).Error
	if err != gorm.ErrRecordNotFound {
		eh.ErrorHandler(err, "getItems", "")
	}
	return
}

func consume(userID uint, commodityID uint, number int) (result bool, amount int, balance int) {
	result = false
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var user data.User
	user.ID = userID
	var commodity data.Commodity
	commodity.ID = commodityID
	err := tx.Where("id = ?", commodity.ID).First(&commodity).Error
	eh.ErrorHandler(err, "consume", "")
	amount = (int(commodity.Price) * number)
	err = tx.Where("id = ?", user.ID).First(&user).Error
	eh.ErrorHandler(err, "consume", "")
	if user.Balance < amount || commodity.Remain < number {
		return
	}
	err = tx.Model(&user).Update("balance", user.Balance-amount).Error
	eh.ErrorHandler(err, "consume", "")
	err = tx.Model(&commodity).Update("remain", commodity.Remain-number).Error
	eh.ErrorHandler(err, "consume", "")
	err = tx.Commit().Error
	eh.ErrorHandler(err, "consume", "")
	if err != nil {
		return false, 0, 0
	} else {
		return true, amount, user.Balance
	}
}
