package controllers

import (
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CreateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&SocialMedia)
	} else {
		ctx.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, SocialMedia)
}

func GetAllSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	allSocialMedia := []models.SocialMedia{}

	db.Find(&allSocialMedia)

	if len(allSocialMedia) == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Not Found",
			"error_message": "No social media found",
		})
		return
	}

	ctx.JSON(http.StatusOK, allSocialMedia)
}

func GetSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	socialMediaID, _ := strconv.Atoi(ctx.Param("socialMediaID"))

	SocialMedia.ID = uint(socialMediaID)

	err := db.First(&SocialMedia, "id = ?", socialMediaID).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, SocialMedia)
}

func UpdateSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	SocialMedia := models.SocialMedia{}

	socialMediaID, _ := strconv.Atoi(ctx.Param("socialMediaID"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&SocialMedia)
	} else {
		ctx.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaID)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaID).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, SocialMedia)
}

func DeleteSocialMedia(ctx *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	err := db.Where("id = ?", socialMediaId).Delete(&SocialMedia).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Delete Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Status":  "Success",
		"Message": "Social Media Deleted",
	})
}
