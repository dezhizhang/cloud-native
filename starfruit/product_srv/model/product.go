package model

type Category struct {
	BaseModel
	ParentCategory   *Category
	ParentCategoryId int32  `json:"parent_category_id"`
	Name             string `gorm:"type:varchar(20);not null" json:"name"`
	Level            int32  `gorm:"type:int(11);not null;default:1" json:"level"`
	IsTab            bool   `gorm:"type:bool;not null;default:false" json:"isTab"`
}

// Brands 表
type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null" json:"name"`
	Logo string `gorm:"type:varchar(255);not null" json:"logo"`
}

type ProductCategoryBrand struct {
	BaseModel
	Category   Category
	CategoryId int32 `gorm:"type:int(11);index:idx_category_brand,unique" json:"category_id"`
	Brands     Brands
	BrandsId   int32 `gorm:"type:int;index:idx_category_brand,unique" json:"brands_id"`
}

type Banner struct {
	BaseModel
	Name  string `gorm:"type:varchar(100);not null" json:"name"`
	Url   string `gorm:"type:varchar(255);not null" json:"url"`
	Index int32  `gorm:"type:int;index:idx_banner,unique" json:"index"`
}
