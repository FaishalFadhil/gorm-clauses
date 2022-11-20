package controllers

import (
	"assignment-api/database"
	"assignment-api/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	currentTime := time.Now()

	// Create
	orderinput := models.Order{CustomerName: input.CustomerName, Items: input.Items, OrderedAt: currentTime}
	db.Create(&orderinput)

	c.JSON(http.StatusOK, gin.H{"data": orderinput})
}

func UpdateOrder(c *gin.Context) {
	var db = database.GetDB()

	var order models.Order
	err := db.First(&order, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(input)
	db.Session(&gorm.Session{FullSaveAssociations: true, AllowGlobalUpdate: true}).Model(&order).Updates(input)
	// db.Model(&order).Updates(&input)
	// db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&order)

	c.JSON(http.StatusOK, gin.H{"data": input})
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

func GetAllItems(c *gin.Context) {
	var db = database.GetDB()

	var items []models.Item
	err := db.Find(&items).Error

	if err != nil {
		fmt.Println("Error getting Item datas :", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetOneItem(c *gin.Context) {
	var db = database.GetDB()

	var item models.Item
	// err := db.Table("Car").Where("Id = ?", c.Param("id")).First(&car).Error
	err := db.First(&item, "ItemId = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": item})
}

func CreateItem(c *gin.Context) {
	var db = database.GetDB()
	// Validate input
	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	ItemInput := models.Item{ItemCode: input.ItemCode, Description: input.Description, Quantity: input.Quantity, OrderId: input.OrderId}
	db.Create(&ItemInput)

	c.JSON(http.StatusOK, gin.H{"data": ItemInput})
}

func UpdateItem(c *gin.Context) {
	var db = database.GetDB()

	var Item models.Item
	err := db.First(&Item, "ItemId = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&Item).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": Item})
}

func DeleteItem(c *gin.Context) {
	var db = database.GetDB()
	// Get model if exist
	var ItemDelete models.Item
	err := db.First(&ItemDelete, "ItemId = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&ItemDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
