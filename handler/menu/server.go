package menu

import (
	"pku-class/market/data"
	db "pku-class/market/database"
	eh "pku-class/market/error-handler"

	"gorm.io/gorm"
)

func getMenu(theme string) (list []data.Menu) {
	err := db.DB.Where("theme = ?", theme).Find(&list).Error
	if err != gorm.ErrRecordNotFound {
		eh.ErrorHandler(err, "getMenu", "")
	}
	return
}
