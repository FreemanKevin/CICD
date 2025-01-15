package handlers

import (
	"demo/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get all items
// @Description Get a list of all items
// @Tags items
// @Accept json
// @Produce json
// @Success 200 {array} models.Item
// @Router /items [get]
func GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, models.Items)
}

// @Summary Get an item
// @Description Get an item by ID
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} models.Item
// @Router /items/{id} [get]
func GetItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, item := range models.Items {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

// @Summary Create an item
// @Description Create a new item
// @Tags items
// @Accept json
// @Produce json
// @Param item body models.Item true "Item object"
// @Success 201 {object} models.Item
// @Router /items [post]
func CreateItem(c *gin.Context) {
	var item models.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.ID = len(models.Items) + 1
	models.Items = append(models.Items, item)
	c.JSON(http.StatusCreated, item)
}

// @Summary Update an item
// @Description Update an existing item
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param item body models.Item true "Item object"
// @Success 200 {object} models.Item
// @Router /items/{id} [put]
func UpdateItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, item := range models.Items {
		if item.ID == id {
			newItem.ID = id
			models.Items[i] = newItem
			c.JSON(http.StatusOK, newItem)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}

// @Summary Delete an item
// @Description Delete an existing item
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} map[string]string
// @Router /items/{id} [delete]
func DeleteItem(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, item := range models.Items {
		if item.ID == id {
			models.Items = append(models.Items[:i], models.Items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Item not found"})
}
