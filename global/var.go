package global

import "gorm.io/gorm"

var (
	DB           *gorm.DB
	DBWithoutLog *gorm.DB
)
