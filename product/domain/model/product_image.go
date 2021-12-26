package model

type ProductImage struct {
	ID             int64  `json:"id" gorm:"primary_key;not_null;auto_increment"`
	ImageName      string `json:"image_name"`
	ImageCode      string `json:"image_code" gorm:"unique_index;not_null"`
	ImageUrl       string `json:"image_url"`
	ImageProductID int64  `json:"image_product_id"`
}
