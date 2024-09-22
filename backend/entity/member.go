
package entity

import (
    "time"

    "gorm.io/gorm"
)

type Member struct {
    gorm.Model
    Password   string `gorm:"not null"`
    Email      string `gorm:"unique;not null"`
    FirstName  string `gorm:"not null"`
    LastName   string `gorm:"not null"`
    Birthday   time.Time
    Gender     string
    TotalPoint uint

    // 1 Member สามารถมีหลาย Payment
    Payments []Payment `gorm:"foreignKey:MemberID"`
}
