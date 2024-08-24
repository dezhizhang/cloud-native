package global

import "gorm.io/gorm"

// SALT 密码的盐
const SALT = "xiaozhicloud"

// DB mysql连接对像
var DB *gorm.DB
