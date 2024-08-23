package main

import (
	"starfruit.top/user/driver"
	"starfruit.top/user/model"
)

func main() {
	driver.DB.Create(&model.User{})
}
