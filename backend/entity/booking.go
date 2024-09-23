// package entity

// import (
//     "time"

//     "gorm.io/gorm"
// )

// type Booking struct {
//     gorm.Model
//     BookingDate time.Time
//     TotalPrice  string

//     // 1 Booking สามารถมี 1 Payment
//     Payment *Payment `gorm:"foreignKey:BookingID"`
// }

package entity

import (
    "time"

    "gorm.io/gorm"
)

type Booking struct {
    gorm.Model
    BookingDate time.Time
    TotalPrice  float64 // เปลี่ยนเป็น float64

    // 1 Booking สามารถมี 1 Payment
    Payment *Payment `gorm:"foreignKey:BookingID"`
}

