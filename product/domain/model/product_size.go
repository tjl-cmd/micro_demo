package model

type ProductSize struct {
	ID            int64  `json:"id" gorm:"primary_key;not_null;auto_increment"`
	SizeName      string `json:"size_name"`
	SizeCode      string `json:"size_code" gorm:"unique_index;not_null"`
	SizeProductID int64  `json:"size_product_id"`
}
