package main

import (
    "fmt"
    "net/http"
    "example.com/paymentSystem/config"
    "example.com/paymentSystem/controller"
    "github.com/gin-gonic/gin"
)

const PORT = "8020"

func main() {
    // เชื่อมต่อกับฐานข้อมูล
    config.ConnectionDB()

    // ตั้งค่าฐานข้อมูลและสร้างข้อมูลตัวอย่าง
    config.SetupDatabase()

    r := gin.Default()

    r.Use(CORSMiddleware())

    // Benefits Routes
    r.POST("/Benefits", controller.CreateBenefits)
    r.GET("/Benefits/:id", controller.GetBenefitsID)
    r.GET("/Benefits", controller.GetBenefits)
    r.DELETE("/Benefits/:id", controller.DeleteBenefits)
    r.PATCH("/Benefits/:id", controller.UpdateBenefits)

    // Booking Routes
    r.POST("/bookings", controller.CreateBooking)
    r.GET("/bookings/:id", controller.GetBookingID)
    r.GET("/bookings", controller.GetBooking)
    r.DELETE("/bookings/:id", controller.DeleteBooking)
    r.PATCH("/bookings/:id", controller.UpdateBooking)

    // Member Routes
    r.POST("/Member", controller.CreateMember)
    r.GET("/Member/:id", controller.GetMemberID)
    r.GET("/Member", controller.GetMember)
    r.DELETE("/Member/:id", controller.DeleteMember)
    r.PATCH("/Member/:id", controller.UpdateMember)

    // Payment Routes
    r.POST("/Payment", controller.CreatePayment)
    r.GET("/Payment/:id", controller.GetPaymentID)
    r.GET("/Payment", controller.GetPayment)
    r.DELETE("/Payment/:id", controller.DeletePayment)
    r.PATCH("/Payment/:id", controller.UpdatePayment)

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, fmt.Sprintf("API RUNNING... PORT: %s", PORT))
    })

    // เริ่มต้นเซิร์ฟเวอร์
    r.Run("localhost:" + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
