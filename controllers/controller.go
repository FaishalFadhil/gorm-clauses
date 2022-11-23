package controllers

import (
	"clauses/database"
	"clauses/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetAllOrders(c *gin.Context) {
	var db = database.GetDB()

	var orders []models.Order
	err := db.Preload("Items").Find(&orders).Error

	if err != nil {
		fmt.Println("Error getting order datas :", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func GetOneOrder(c *gin.Context) {
	var db = database.GetDB()

	var orderOne models.Order
	// err := db.Table("Order").Where("Id = ?", c.Param("id")).First(&Order).Error
	err := db.Preload("Items").First(&orderOne, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": orderOne})
}

func CreateOrder(c *gin.Context) {
	var db = database.GetDB()
	// Validate input
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	orderinput := models.Order{CustomerName: input.CustomerName, Items: input.Items}
	db.Create(&orderinput)

	c.JSON(http.StatusOK, gin.H{"data": orderinput})
}

func UpdateOrder(c *gin.Context) {
	db := database.GetDB()
	var orders models.Order

	err := db.Preload("Items").First(&orders, "id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// fmt.Println(&input)

	u64, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		fmt.Println(err)
	}
	wd := uint(u64)

	// fmt.Println(orders)

	orderinput := models.Order{Id: wd, CustomerName: input.CustomerName, Items: input.Items}

	for _, v := range input.Items {
		if v.Id == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Id not found",
				"message": "Id not found on items list",
			})
			return
		}
	}
	// Input Update Order
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},                       // key colume
		DoUpdates: clause.AssignmentColumns([]string{"customer_name"}), // column needed to be updated
	}).Create(&orderinput)

	// Input Update Item
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},                                              // key colume
		DoUpdates: clause.AssignmentColumns([]string{"item_code", "description", "quantity"}), // column needed to be updated
	}).Create(&orderinput.Items)

	c.JSON(http.StatusOK, gin.H{
		"data": orderinput,
		"code": http.StatusOK,
	})
}

func UpdateDoNothingOrder(c *gin.Context) {
	db := database.GetDB()
	var orders models.Order

	err := db.Preload("Items").First(&orders, "id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// fmt.Println(&input)
	db.Clauses(clause.OnConflict{DoNothing: true}).Create(&input)

	c.JSON(http.StatusOK, gin.H{
		"data": input,
		"code": http.StatusOK,
	})
}

func UpdateNoConflictOrder(c *gin.Context) {
	db := database.GetDB()
	var orders models.Order

	err := db.Preload("Items").First(&orders, "id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "request not found",
			"message": err.Error(),
		})
		return
	}

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// fmt.Println(&input)
	db.Clauses(clause.Insert{}).Create(&input)

	c.JSON(http.StatusOK, gin.H{
		"data": input,
		"code": http.StatusOK,
	})
}

func DeleteOrder(c *gin.Context) {
	var db = database.GetDB()
	// Get model if exist
	var orderDelete models.Order
	err := db.First(&orderDelete, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&orderDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
