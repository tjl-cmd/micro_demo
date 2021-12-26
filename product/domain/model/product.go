package model

type Product struct {
	ID                 int64          `json:"id" gorm:"primary_key;not_null;auto_increment"`
	ProductName        string         `json:"product_name"`
	ProductSku         string         `json:"product_sku" gorm:"unique_index;not_null"`
	ProductPrice       float64        `json:"product_price"`
	ProductDescription string         `json:"product_description"`
	ProductImage       []ProductImage `json:"product_image" gorm:"ForeignKey:ImageProductID"`
	ProductSize        []ProductSize  `json:"product_size" gorm:"ForeignKey:SizeProductID"`
	ProductSeo         ProductSeo     `json:"product_seo" gorm:"ForeignKey:SeoProductID"`
}
