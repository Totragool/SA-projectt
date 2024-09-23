// package entity

// import "gorm.io/gorm"

// type Benefits struct {
//     gorm.Model
//     BenefitsName  string `gorm:"unique;not null"`
//     FlyingFrom    string
//     GoingTo       string
//     PointRequired uint
//     Quantity      uint
//     Code          string `gorm:"unique;not null"`
//     Trip          string
//     Type          string

//     // 1 Benefits สามารถมีหลาย Payment
//     Payments []Payment `gorm:"foreignKey:BenefitID"`
// }

package entity

import "gorm.io/gorm"

type Benefits struct {
    gorm.Model
    BenefitsName  string `gorm:"unique;not null"`
    FlyingFrom    string
    GoingTo       string
    PointRequired uint
    Quantity      uint
    Code          string `gorm:"unique;not null"`
    Trip          string
    Type          string

    // 1 Benefits สามารถมีหลาย Payment
    Payments []Payment `gorm:"foreignKey:BenefitID"`
}
