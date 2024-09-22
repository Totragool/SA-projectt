package controller

import (
    "net/http"
    "example.com/paymentSystem/entity"
    "example.com/paymentSystem/config"
    "github.com/gin-gonic/gin"
)

// POST /Member
func CreateMember(c *gin.Context) {
    var member entity.Member
    if err := c.ShouldBindJSON(&member); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // TODO: Hash the password before saving
    // member.Password = HashPassword(member.Password)

    if err := config.DB().Create(&member).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"data": member})
}

// GET /Member
func GetMember(c *gin.Context) {
    var members []entity.Member

    db := config.DB()
    results := db.Find(&members)
    if results.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, members)
}

// GET /Member/:id
func GetMemberID(c *gin.Context) {
    var member entity.Member
    id := c.Param("id")

    if err := config.DB().First(&member, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": member})
}

// DELETE /Member/:id
func DeleteMember(c *gin.Context) {
    id := c.Param("id")
    if tx := config.DB().Delete(&entity.Member{}, id); tx.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /Member/:id
func UpdateMember(c *gin.Context) {
    var member entity.Member
    id := c.Param("id")

    if err := config.DB().First(&member, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    if err := c.ShouldBindJSON(&member); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // TODO: Hash the password if it's being updated
    // member.Password = HashPassword(member.Password)

    if err := config.DB().Save(&member).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": member})
}
