package model

import "time"

type BaseModel struct {
	Id        int64     `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"DEFAULT:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `gorm:"DEFAULT:NULL"`
	IsDeleted int       `gorm:"DEFAULT:0"`
}
