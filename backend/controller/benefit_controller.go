package controller

import (
    "net/http"
    "example.com/paymentSystem/entity"
    "example.com/paymentSystem/config"
    "github.com/gin-gonic/gin"
)

// POST /Benefits
func CreateBenefits(c *gin.Context) {
    var benefit entity.Benefits
    if err := c.ShouldBindJSON(&benefit); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB().Create(&benefit).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"data": benefit})
}

// GET /Benefits
func GetBenefits(c *gin.Context) {
    var benefits []entity.Benefits

    db := config.DB()
    results := db.Find(&benefits)
    if results.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, benefits)
}

// GET /Benefits/:id
func GetBenefitsID(c *gin.Context) {
    var benefit entity.Benefits
    id := c.Param("id")

    if err := config.DB().First(&benefit, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": benefit})
}

// DELETE /Benefits/:id
func DeleteBenefits(c *gin.Context) {
    id := c.Param("id")
    if tx := config.DB().Delete(&entity.Benefits{}, id); tx.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Benefits not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Benefits/:id
func UpdateBenefits(c *gin.Context) {
    var benefit entity.Benefits
    id := c.Param("id")

    if err := config.DB().First(&benefit, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    if err := c.ShouldBindJSON(&benefit); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB().Save(&benefit).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": benefit})
}
