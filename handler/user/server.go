package user

import (
	"pku-class/market/data"
	db "pku-class/market/database"
	eh "pku-class/market/error-handler"

	"gorm.io/gorm"
)

func authorize(user *data.User, authorizeType string) (result bool) {
	result = false
	password := user.Password
	err := db.DB.Where("username = ?", user.Username).First(user).Error
	if authorizeType == "token" {
		result = true
	} else if err == nil && password == user.Password {
		result = true
	}
	if err != gorm.ErrRecordNotFound {
		eh.ErrorHandler(err, "authorize", "")
	}
	return
}

func getOrders(username string, role int) (orders []data.Order) {
	if role == 0 {
		err := db.DB.Find(&orders).Error
		if err != gorm.ErrRecordNotFound {
			eh.ErrorHandler(err, "getOrders", "")
		}
	} else {
		err := db.DB.Where("username = ?", username).Find(&orders).Error
		if err != gorm.ErrRecordNotFound {
			eh.ErrorHandler(err, "getOrders", "")
		}
	}
	return
}

func deleteOrder(id uint) (result bool) {
	result = false
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var order data.Order
	order.ID = id
	err := tx.Delete(&order).Error
	eh.ErrorHandler(err, "deleteOrder", "")
	err = tx.Commit().Error
	eh.ErrorHandler(err, "deleteOrder", "")
	if err != nil {
		return false
	} else {
		return true
	}
}

func charge(id uint, amount int) (result bool, balance int) {
	result = false
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var user data.User
	user.ID = id
	err := tx.Where("id = ?", user.ID).First(&user).Error
	eh.ErrorHandler(err, "charge", "")
	err = tx.Model(&user).Update("balance", user.Balance+amount).Error
	eh.ErrorHandler(err, "charge", "")
	err = tx.Commit().Error
	eh.ErrorHandler(err, "charge", "")
	if err != nil {
		return false, user.Balance
	} else {
		return true, user.Balance
	}
}
