package data

import (
	"gorm.io/gorm"
)

/*type ModelHide struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type ModelShow struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}*/

type User struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Token    string `json:"token" gorm:"-"`
	Username string `json:"username" gorm:"not null,unique"`
	Password string `json:"password,omitempty"`
	Role     int    `json:"role" gorm:"default:-1"`
	Balance  int    `json:"balance" gorm:"default:0"`
}

type Commodity struct {
	ID      uint   `json:"id" gorm:"primarykey"`
	Name    string `json:"name" gorm:"not null"`
	Remain  int    `json:"remain" gorm:"unsigned"`
	Comment string `json:"comment"`
	Price   int    `json:"price"gorm:"unsigned"`
	Type    string `json:"type"`
}

type Order struct {
	gorm.Model
	Username      string `json:"username" gorm:"not null"`
	CommodityName string `json:"commodityName" gorm:"not null"`
	Number        int    `json:"number" gorm:"not null"`
	Amount        int    `json:"amount" gorm:"not null"`
	State         int    `json:"state"gorm:"not null"`
}

type Menu struct {
	ID    uint   `json:"id" gorm:"primarykey"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
	Theme string `json:"theme"`
}
