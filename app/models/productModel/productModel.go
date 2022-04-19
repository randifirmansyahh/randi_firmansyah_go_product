package productModel

import (
	"randi_firmansyah/app/helper/modelHelper"
	"randi_firmansyah/app/models/categoryModel"
)

type Product struct {
	Id          int                    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Category_Id int                    `json:"category_id"`
	Category    categoryModel.Category `json:"category" gorm:"foreignKey:Category_Id;references:id"`
	Nama        string                 `json:"nama"`
	Harga       int                    `json:"harga"`
	Qty         int                    `json:"qty"`
	Image       string                 `json:"image"`
	modelHelper.DateAuditModel
}

type ProductResponse struct {
	Id          int    `json:"id"`
	Category_Id int    `json:"category_id"`
	Nama        string `json:"nama"`
	Harga       int    `json:"harga"`
	Qty         int    `json:"qty"`
	Image       string `json:"image"`
	modelHelper.DateAuditModel
}
