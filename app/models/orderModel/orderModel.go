package orderModel

import (
	"randi_firmansyah/app/helper/modelHelper"
	"randi_firmansyah/app/models/productModel"
	"randi_firmansyah/app/models/userModel"
)

type Order struct {
	Id          int                  `gorm:"primaryKey;autoIncrement;" json:"id"`
	User_Id     int                  `json:"user_id"`
	User        userModel.User       `json:"user" gorm:"foreignKey:User_Id;references:id"`
	Product_Id  int                  `json:"product_id"`
	Product     productModel.Product `json:"product" gorm:"foreignKey:Product_Id;references:id"`
	Qty         int                  `json:"qty"`
	Total       int                  `json:"total"`
	OrderStatus bool                 `json:"order_status"`
	modelHelper.DateAuditModel
}

type OrderResponse struct {
	Id          int  `json:"id"`
	User_Id     int  `json:"user_id"`
	Product_Id  int  `json:"product_id"`
	Qty         int  `json:"qty"`
	Total       int  `json:"total"`
	OrderStatus bool `json:"order_status"`
	modelHelper.DateAuditModel
}
