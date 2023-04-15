package controllers

import (
	"fmt"
	"go-jwt/database"
	"go-jwt/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(product)

	product.UserID = uint(userData["id"].(float64))

	err = db.Create(&product).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := models.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	product.UserID = uint(userData["id"].(float64))

	err = db.Model(&product).Where("id=?", productID).Updates(models.Product{Title: product.Title, Description: product.Description}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
