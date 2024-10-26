package controllers

import (
	"errors"
	"exchangeapp/backend/global"
	"exchangeapp/backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatExchangeRate(ctx *gin.Context) {
	var rate models.ExchangeRate

	if err := ctx.ShouldBindBodyWithJSON(&rate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rate.Date = time.Now()

	if err := global.Db.AutoMigrate(&rate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := global.Db.Create(&rate)

	if err := result.Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, rate)
}

func GetExchangeRates(ctx *gin.Context) {
	var rates []models.ExchangeRate

	result := global.Db.Find(&rates)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, rates)
}
