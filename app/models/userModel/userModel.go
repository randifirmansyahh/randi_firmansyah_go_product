package userModel

import "randi_firmansyah/app/helper/modelHelper"

type User struct {
	Id       int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	No_Hp    string `json:"no_hp"`
	Image    string `json:"image"`
	modelHelper.DateAuditModel
}
