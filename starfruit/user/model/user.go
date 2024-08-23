package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        int64          `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	IsDeleted bool           `gorm:"column:is_deleted" json:"is_deleted"`
}

type User struct {
	BaseModel
	Username string     `gorm:"type:varchar(20);not null" json:"username"`                       // 名称
	Mobile   string     `gorm:"index:idx_mobile;unique,type:varchar(11);not null" json:"mobile"` // 手机号
	Password string     `gorm:"type:varchar(120);not null" json:"password"`                      // 密码
	BirthDay *time.Time `gorm:"type:datetime" json:"birthday"`                                   // 生日
	Gender   string     `gorm:"type:varchar(3);not null default:'male'" json:"gender"`           // 性别
	Role     int        `gorm:"type:int;default:1" json:"role"`                                  // 角色
}
