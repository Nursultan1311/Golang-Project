package controllers

import (
	"Test2/initializers"
	"Test2/pkg/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("Secret_key"))

func Addtoorder(c *gin.Context) {
	var body struct {
		Bookid   int8
		Quantity int8
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse info",
		})
		return
	}
	var user models.User
	initializers.GetDB().Find(&user, "email=?", GetUserEmail(c))

	var order models.Order
	order.CreatedDate = time.Now().Format("Jan 2, 2006 03:04:05 PM")

	order.Quantity = body.Quantity

	order.BookId = body.Bookid

	order.UserId = int8(user.ID)

	var count int64
	er := initializers.GetDB().Debug().Table("orders").Where("user_id = ? AND book_id = ?", user.ID, body.Bookid).Count(&count)
	fmt.Println(count)
	if er.Error != nil {
		log.Println(er.Error)
	}

	if count == 0 {

		err := initializers.GetDB().Create(&order)

		if err.Error != nil {

			log.Println(err.Error)

		}

		var count int64

		initializers.GetDB().Where("user_id = ?", user.ID).Count(&count)

		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "Added in your Cart",
			"count": count,
		})
		return

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "This product already in your cart",
		})
		return
	}
}

func Viewmyorders(c *gin.Context) {

	var user models.User
	initializers.GetDB().Find(&user, "email=?", GetUserEmail(c))

	var orders []models.Order
	err := initializers.GetDB().Where("user_id=?", user.ID).Find(&orders)

	if err.Error != nil {
		log.Println(err.Error)
	}

	var ordercount int64
	initializers.GetDB().Table("orders").Where("user_id = ?", user.ID).Count(&ordercount)

	c.JSON(200, gin.H{
		"orderlist":   orders,
		"orderscount": ordercount,
	})
}
