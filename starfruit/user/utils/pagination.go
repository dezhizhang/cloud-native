package utils

import "gorm.io/gorm"

// Pagination 获取分页数
func Pagination(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
