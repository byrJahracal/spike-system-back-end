package handler

import (
	"pku-class/market/data"
	db "pku-class/market/database"
	eh "pku-class/market/error-handler"
)

func flash(userID uint, commodityID uint, number int) {
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
	eh.ErrorHandler(err, "flash", "")
	amount := (int(commodity.Price) * number)
	err = tx.Where("id = ?", user.ID).First(&user).Error
	eh.ErrorHandler(err, "flash", "")
	if user.Balance < amount || commodity.Remain < number {
		return
	}
	err = tx.Model(&user).Update("balance", user.Balance-amount).Error
	eh.ErrorHandler(err, "flash", "")
	err = tx.Model(&commodity).Update("remain", commodity.Remain-number).Error
	eh.ErrorHandler(err, "flash", "")
	order := data.Order{
		Username:      user.Username,
		CommodityName: commodity.Name,
		Number:        number,
		Amount:        amount,
		State:         0,
	}
	err = tx.Create(&order).Error
	eh.ErrorHandler(err, "flash", "")
	err = tx.Commit().Error
	eh.ErrorHandler(err, "flash", "")
	return
}
