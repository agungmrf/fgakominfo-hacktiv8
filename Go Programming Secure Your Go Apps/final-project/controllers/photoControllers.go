package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Photo)
}

func GetAllPhoto(c *gin.Context) {
	db := database.GetDB()
	allPhoto := []models.Photo{}

	db.Find(&allPhoto)

	if len(allPhoto) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Not Found",
			"error_message": "No photo found",
		})
		return
	}

	c.JSON(http.StatusOK, allPhoto)
}

func GetPhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	photoID, _ := strconv.Atoi(c.Param("photoID"))

	Photo.ID = uint(photoID)

	err := db.First(&Photo, "id = ?", photoID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoID, _ := strconv.Atoi(c.Param("photoID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoID)

	err := db.Model(&Photo).Where("id = ?", photoID).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoURL: Photo.PhotoURL}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Photo)
}

func DeletePhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	photoID, _ := strconv.Atoi(c.Param("photoID"))

	err := db.Where("id = ?", photoID).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Photo Deleted",
	})
}
