package model

type ProductSeo struct {
	ID             int64  `json:"id" gorm:"primary_key;not_null;auto_increment"`
	SeoTitle       string `json:"seo_title"`
	SeoKeywords    string `json:"seo_keywords"`
	SeoDescription string `json:"seo_description"`
	SeoProductID   int64  `json:"seo_product_id"`
}
