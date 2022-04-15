package productModel

type Product struct {
	Id    int    `gorm:"primaryKey;autoIncrement;" json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Qty   int    `json:"qty"`
	Image string `json:"image"`
}
