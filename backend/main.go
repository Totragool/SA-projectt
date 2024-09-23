// package main

// import (
//     "fmt"
//     "net/http"
//     "example.com/paymentSystem/config"
//     "example.com/paymentSystem/controller"
//     "github.com/gin-gonic/gin"
// )

// const PORT = "8020"

// func main() {
//     // เชื่อมต่อกับฐานข้อมูล
//     config.ConnectionDB()

//     // ตั้งค่าฐานข้อมูลและสร้างข้อมูลตัวอย่าง
//     config.SetupDatabase()

//     r := gin.Default()

//     r.Use(CORSMiddleware())

//     // Benefits Routes
//     r.POST("/Benefits", controller.CreateBenefits)
//     r.GET("/Benefits/:id", controller.GetBenefitsID)
//     r.GET("/Benefits", controller.GetBenefits)
//     r.DELETE("/Benefits/:id", controller.DeleteBenefits)
//     r.PATCH("/Benefits/:id", controller.UpdateBenefits)

//     // Booking Routes
//     r.POST("/bookings", controller.CreateBooking)
//     r.GET("/bookings/:id", controller.GetBookingID)
//     r.GET("/bookings", controller.GetBooking)
//     r.DELETE("/bookings/:id", controller.DeleteBooking)
//     r.PATCH("/bookings/:id", controller.UpdateBooking)

//     // Member Routes
//     r.POST("/Member", controller.CreateMember)
//     r.GET("/Member/:id", controller.GetMemberID)
//     r.GET("/Member", controller.GetMember)
//     r.DELETE("/Member/:id", controller.DeleteMember)
//     r.PATCH("/Member/:id", controller.UpdateMember)

//     // Payment Routes
//     r.POST("/Payment", controller.CreatePayment)
//     r.GET("/Payment/:id", controller.GetPaymentID)
//     r.GET("/Payment", controller.GetPayment)
//     r.DELETE("/Payment/:id", controller.DeletePayment)
//     r.PATCH("/Payment/:id", controller.UpdatePayment)
//     r.POST("/Payment/Mock", controller.MockPayment)

//     r.GET("/", func(c *gin.Context) {
//         c.String(http.StatusOK, fmt.Sprintf("API RUNNING... PORT: %s", PORT))
//     })

//     // เริ่มต้นเซิร์ฟเวอร์
//     r.Run("localhost:" + PORT)
// }

// func CORSMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
        
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }

//         c.Next()
//     }
// }

// package main

// import (
//     "fmt"
//     "log"
//     "net/http"
//     "os"

//     "example.com/paymentSystem/config"
//     "example.com/paymentSystem/controller"
//     "github.com/gin-gonic/gin"
// )

// func main() {
    

//     // เชื่อมต่อกับฐานข้อมูล
//     config.ConnectionDB()

//     // ตั้งค่าฐานข้อมูลและสร้างข้อมูลตัวอย่าง
//     config.SetupDatabase()

//     r := gin.Default()

//     r.Use(CORSMiddleware())

//     // Group Routes
//     benefitsRoutes := r.Group("/Benefits")
//     {
//         benefitsRoutes.POST("/", controller.CreateBenefits)
//         benefitsRoutes.GET("/:id", controller.GetBenefitsID)
//         benefitsRoutes.GET("/", controller.GetBenefits)
//         benefitsRoutes.DELETE("/:id", controller.DeleteBenefits)
//         benefitsRoutes.PATCH("/:id", controller.UpdateBenefits)
//     }

//     bookingsRoutes := r.Group("/bookings")
//     {
//         bookingsRoutes.POST("/", controller.CreateBooking)
//         bookingsRoutes.GET("/:id", controller.GetBookingID)
//         bookingsRoutes.GET("/", controller.GetBooking)
//         bookingsRoutes.DELETE("/:id", controller.DeleteBooking)
//         bookingsRoutes.PATCH("/:id", controller.UpdateBooking)
//     }

//     membersRoutes := r.Group("/Member")
//     {
//         membersRoutes.POST("/", controller.CreateMember)
//         membersRoutes.GET("/:id", controller.GetMemberID)
//         membersRoutes.GET("/", controller.GetMember)
//         membersRoutes.DELETE("/:id", controller.DeleteMember)
//         membersRoutes.PATCH("/:id", controller.UpdateMember)
//     }

//     paymentsRoutes := r.Group("/Payment")
//     {
//         paymentsRoutes.POST("/", controller.CreatePayment)
//         paymentsRoutes.GET("/:id", controller.GetPaymentID)
//         paymentsRoutes.GET("/", controller.GetPayment)
//         paymentsRoutes.DELETE("/:id", controller.DeletePayment)
//         paymentsRoutes.PATCH("/:id", controller.UpdatePayment)
//         paymentsRoutes.POST("/Mock", controller.CreatePayment) // ใช้ CreatePayment แทน MockPayment
//     }

//     r.GET("/", func(c *gin.Context) {
//         port := getPort()
//         c.String(http.StatusOK, fmt.Sprintf("API RUNNING... PORT: %s", port))
//     })

//     // เริ่มต้นเซิร์ฟเวอร์
//     port := getPort()
//     if err := r.Run(":" + port); err != nil {
//         log.Fatalf("Failed to run server: %v", err)
//     }
// }

// func CORSMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         allowedOrigins := []string{
//             "http://localhost:3000", // เพิ่ม Origin ที่อนุญาต
//             "https://localhost:5173",
//         }

//         origin := c.GetHeader("Origin")
//         allowed := false
//         for _, o := range allowedOrigins {
//             if o == origin {
//                 allowed = true
//                 break
//             }
//         }

//         if allowed {
//             c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
//         }
//         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }

//         c.Next()
//     }
// }

// func getPort() string {
//     port := os.Getenv("PORT")
//     if port == "" {
//         port = "8020" // ค่าเริ่มต้น
//     }
//     return port
// }

// func loadEnv() {
//     // โหลด environment variables จากไฟล์ .env หากมี
//     // ใช้ package เช่น github.com/joho/godotenv
//     // ถ้าต้องการใช้งาน ให้ทำการติดตั้งและเพิ่มโค้ดด้านล่าง
//     /*
//        err := godotenv.Load()
//        if err != nil {
//            log.Println("No .env file found")
//        }
//     */
// }


package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "example.com/paymentSystem/config"
    "example.com/paymentSystem/controller"
)

func main() {

    // เชื่อมต่อกับฐานข้อมูล
    config.ConnectionDB()

    // ตั้งค่าฐานข้อมูลและสร้างข้อมูลตัวอย่าง
    config.SetupDatabase()

    // สร้าง instance ของ Gin
    r := gin.Default()

    // ใช้ middleware CORS
    r.Use(CORSMiddleware())

    // Group Routes
    benefitsRoutes := r.Group("/Benefits")
    {
        benefitsRoutes.POST("/", controller.CreateBenefits)
        benefitsRoutes.GET("/:id", controller.GetBenefitsID)
        benefitsRoutes.GET("/", controller.GetBenefits)
        benefitsRoutes.DELETE("/:id", controller.DeleteBenefits)
        benefitsRoutes.PATCH("/:id", controller.UpdateBenefits)
    }
    
    bookingsRoutes := r.Group("/bookings")
    {
        bookingsRoutes.POST("/", controller.CreateBooking)
        bookingsRoutes.GET("/:id", controller.GetBookingID)
        bookingsRoutes.GET("/", controller.GetBooking)
        bookingsRoutes.DELETE("/:id", controller.DeleteBooking)
        bookingsRoutes.PATCH("/:id", controller.UpdateBooking)
    }

    membersRoutes := r.Group("/Member")
    {
        membersRoutes.POST("/", controller.CreateMember)
        membersRoutes.GET("/:id", controller.GetMemberID)
        membersRoutes.GET("/:id", controller.GetBookingID)
        membersRoutes.GET("/", controller.GetMember)
        membersRoutes.DELETE("/:id", controller.DeleteMember)
        membersRoutes.PATCH("/:id", controller.UpdateMember)
    }

    paymentsRoutes := r.Group("/Payment")
    {
        paymentsRoutes.POST("/", controller.CreatePayment)
        paymentsRoutes.GET("/:id", controller.GetPaymentID)
        paymentsRoutes.GET("/", controller.GetPayment)
        paymentsRoutes.DELETE("/:id", controller.DeletePayment)
        paymentsRoutes.PATCH("/:id", controller.UpdatePayment)
        paymentsRoutes.POST("/Mock", controller.CreatePayment) // ใช้ CreatePayment แทน MockPayment
    }

    // Route ทดสอบ API ว่าทำงานอยู่
    r.GET("/", func(c *gin.Context) {
        port := getPort()
        c.String(http.StatusOK, fmt.Sprintf("API RUNNING... PORT: %s", port))
    })

    // เริ่มต้นเซิร์ฟเวอร์
    port := getPort()
    if err := r.Run(":" + port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        allowedOrigins := []string{
            "http://localhost:3000", // เพิ่ม Origin ที่อนุญาต
            "https://localhost:5173", // Frontend URL ที่จะเชื่อมต่อ
        }

        origin := c.GetHeader("Origin")
        allowed := false
        for _, o := range allowedOrigins {
            if o == origin {
                allowed = true
                break
            }
        }

        if allowed {
            c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
        }
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



// ฟังก์ชันหาพอร์ตจาก ENV หรือใช้พอร์ต 8020 โดยค่าเริ่มต้น
func getPort() string {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8020" // ค่าเริ่มต้น
    }
    return port
}
