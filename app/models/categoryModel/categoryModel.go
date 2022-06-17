package categoryModel

import "randi_firmansyah/app/helper/modelHelper"

type Category struct {
	Id   int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama string `json:"nama"`
	modelHelper.DateAuditModel
}

type CategoryReq struct {
	Nama string `json:"nama" validate:"required"`
}
