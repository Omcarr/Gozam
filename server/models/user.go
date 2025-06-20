package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // provides standard fields
	ID         int    `json:"ID" gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `json:"email" gorm:"unique"`
	Password   string `json:"password"`
}

//will hash the passwords in future
// func (u *User) HashPassword() error {
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)
// 	return nil
// }
