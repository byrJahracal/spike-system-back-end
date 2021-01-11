package data

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PackingToGinJson(c *gin.Context, object interface{}) {
	switch object.(type) {
	case User:
		{
			var role string
			if object.(User).Role == 0 {
				role = "seller"
			} else if object.(User).Role == 1 {
				role = "consumer"
			} else {
				role = ""
			}
			c.JSON(http.StatusOK, gin.H{
				"username": object.(User).Username,
				"role":     role,
				"token":    object.(User).Token,
				"balance":  object.(User).Balance,
				"result":   true,
			})
		}
	case []User:
		{
			var users []gin.H
			for _, s := range object.([]User) {
				var role string
				if s.Role == 0 {
					role = "seller"
				} else if s.Role == 1 {
					role = "consumer"
				} else {
					role = ""
				}
				h := gin.H{
					"token":    s.Token,
					"username": s.Username,
					"role":     role,
					"balance":  s.Balance,
					"result":   true,
				}
				users = append(users, h)
			}
			c.JSON(http.StatusOK, gin.H{
				"total": len(users),
				"users": users,
			})
		}
	case Commodity:
		{
			c.JSON(http.StatusOK, gin.H{
				"commodity": gin.H{
					"id":      object.(Commodity).ID,
					"name":    object.(Commodity).Name,
					"remain":  object.(Commodity).Remain,
					"comment": object.(Commodity).Comment,
					"price":   object.(Commodity).Price,
					"type":    object.(Commodity).Type,
				},
			})
		}
	case []Commodity:
		{
			var commoditys []gin.H
			for _, s := range object.([]Commodity) {
				h := gin.H{
					"id":      s.ID,
					"name":    s.Name,
					"remain":  s.Remain,
					"comment": s.Comment,
					"price":   s.Price,
					"type":    s.Type,
				}
				commoditys = append(commoditys, h)
			}
			c.JSON(http.StatusOK, gin.H{
				"total": len(commoditys),
				"items": commoditys,
			})
		}
	case Order:
		{
			var state string
			if object.(Order).State == 0 {
				state = "秒杀成功"
			} else if object.(Order).State == 1 {
				state = "秒杀失败"
			} else {
				log.Fatalln("wrong user.role", object.(Order).State)
			}
			c.JSON(http.StatusOK, gin.H{
				"order": gin.H{
					"id":            object.(Order).ID,
					"username":      object.(Order).Username,
					"commodityName": object.(Order).CommodityName,
					"number":        object.(Order).Number,
					"amount":        object.(Order).Amount,
					"state":         state,
					"time":          object.(Order).CreatedAt,
				},
			})
		}
	case []Order:
		{
			var orders []gin.H
			for _, s := range object.([]Order) {
				var state string
				if s.State == 0 {
					state = "秒杀成功"
				} else if s.State == 1 {
					state = "秒杀失败"
				} else {
					log.Fatalln("wrong user.role", s.State)
				}
				h := gin.H{
					"id":            s.ID,
					"username":      s.Username,
					"commodityName": s.CommodityName,
					"number":        s.Number,
					"amount":        s.Amount,
					"state":         state,
					"time":          s.CreatedAt,
				}
				orders = append(orders, h)
			}
			c.JSON(http.StatusOK, gin.H{
				"total": len(orders),
				"items": orders,
			})
		}
	case Menu:
		{
			c.JSON(http.StatusOK, gin.H{
				"menu": gin.H{
					"icon":  object.(Menu).Icon,
					"title": object.(Menu).Title,
				},
			})
		}
	case []Menu:
		{
			var menus []gin.H
			for _, s := range object.([]Menu) {
				h := gin.H{
					"icon":  s.Icon,
					"title": s.Title,
				}
				menus = append(menus, h)
			}
			c.JSON(http.StatusOK, gin.H{
				"total": len(menus),
				"menu":  menus,
			})
		}
	}

}
