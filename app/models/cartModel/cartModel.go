package cartModel

import (
	"randi_firmansyah/app/helper/modelHelper"
	"randi_firmansyah/app/models/productModel"
	"randi_firmansyah/app/models/userModel"
)

type Cart struct {
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

type CartRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=50,alphanum"`
	ProductId int    `json:"product_id" validate:"required,min=3,max=50,numeric"`
	Qty       int    `json:"qty" validate:"required,min=1,numeric"`
}

type CartResponse struct {
	Id       int                  `json:"id"`
	Username string               `json:"username"`
	Product  productModel.Product `json:"product"`
	Qty      int                  `json:"qty"`
	Total    int                  `json:"total"`
	modelHelper.DateAuditModel
}
