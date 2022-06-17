package productModel

import (
	"randi_firmansyah/app/helper/modelHelper"
	"randi_firmansyah/app/models/categoryModel"
)

type Product struct {
	Id          int                    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Category_Id int                    `json:"category_id" validate:"required,numeric"`
	Category    categoryModel.Category `json:"category" gorm:"foreignKey:Category_Id;references:id"`
	Nama        string                 `json:"nama" validate:"required"`
	Harga       int                    `json:"harga" validate:"required,numeric"`
	Qty         int                    `json:"qty" validate:"required,numeric"`
	Image       string                 `json:"image" validate:"required"`
	modelHelper.DateAuditModel
}

type ProductResponse struct {
	Id       int                    `json:"id"`
	Category categoryModel.Category `json:"category"`
	Nama     string                 `json:"nama"`
	Harga    int                    `json:"harga"`
	Qty      int                    `json:"qty"`
	Image    string                 `json:"image"`
	modelHelper.DateAuditModel
}
