package controller

import (
    "net/http"
    "example.com/paymentSystem/entity"
    "example.com/paymentSystem/config"
    "github.com/gin-gonic/gin"
)

// POST /Payment
func CreatePayment(c *gin.Context) {
    var payment entity.Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if payment.PaymentDate.IsZero() {
        c.JSON(http.StatusBadRequest, gin.H{"error": "PaymentDate is required"})
        return
    }

    // Optionally, you can add validations to check if MemberID, BookingID, and BenefitID exist
    var member entity.Member
    if err := config.DB().First(&member, payment.MemberID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Member not found"})
        return
    }

    var booking entity.Booking
    if err := config.DB().First(&booking, payment.BookingID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Booking not found"})
        return
    }

    var benefit entity.Benefits
    if err := config.DB().First(&benefit, payment.BenefitID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Benefit not found"})
        return
    }

    if err := config.DB().Create(&payment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"data": payment})
}

// GET /Payment
func GetPayment(c *gin.Context) {
    var payments []entity.Payment

    db := config.DB()
    results := db.Find(&payments)
    if results.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": results.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, payments)
}

// GET /Payment/:id
func GetPaymentID(c *gin.Context) {
    var payment entity.Payment
    id := c.Param("id")

    if err := config.DB().First(&payment, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": payment})
}

// DELETE /Payment/:id
func DeletePayment(c *gin.Context) {
    var payment entity.Payment
    id := c.Param("id")

    if err := config.DB().First(&payment, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err := config.DB().Delete(&payment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}

// PATCH /Payment/:id
func UpdatePayment(c *gin.Context) {
    var payment entity.Payment
    id := c.Param("id")

    if err := config.DB().First(&payment, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if payment.PaymentDate.IsZero() {
        c.JSON(http.StatusBadRequest, gin.H{"error": "PaymentDate is required"})
        return
    }

    if err := config.DB().Save(&payment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": payment})
}
