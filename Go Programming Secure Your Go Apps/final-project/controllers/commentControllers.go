package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	photoID, _ := strconv.Atoi(c.Param("photoID"))
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	Photo := models.Photo{}

	err := db.Select("user_id").First(&Photo, uint(photoID)).Error

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Photo Not Found",
			"message": "Photo doesn't exist, failed to create comment",
		})
		return
	}

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.PhotoID = uint(photoID)

	err = db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Comment)
}

func GetAllComment(c *gin.Context) {
	db := database.GetDB()
	allComment := []models.Comment{}

	db.Find(&allComment)

	if len(allComment) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Not Found",
			"error_message": "No comment found",
		})
		return
	}

	c.JSON(http.StatusOK, allComment)
}

func GetAllCommentPhoto(c *gin.Context) {
	db := database.GetDB()
	photoID, _ := strconv.Atoi(c.Param("photoID"))
	allComment := []models.Comment{}

	db.Where("photo_id = ?", photoID).Find(&allComment)

	if len(allComment) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Not Found",
			"error_message": "No comment found",
		})
		return
	}

	c.JSON(http.StatusOK, allComment)
}

func GetComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	commentID, _ := strconv.Atoi(c.Param("commentID"))

	Comment.ID = uint(commentID)

	err := db.First(&Comment, "id = ?", commentID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentID, _ := strconv.Atoi(c.Param("commentID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentID)

	err := db.Model(&Comment).Where("id = ?", commentID).Updates(models.Comment{Message: Comment.Message}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	err = db.First(&Comment, "id = ?", commentID).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	commentID, _ := strconv.Atoi(c.Param("commentID"))

	err := db.Where("id = ?", commentID).Delete(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Comment Deleted",
	})
}
