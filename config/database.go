package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewConnection() (*gorm.DB, error) {
	db, err := gorm.Open(
		"mysql",
		"root:@tcp(db_1:3306)/detoplan?multiStatements=True&parseTime=True&loc=Asia%2FTokyo",
	)
	db.LogMode(true)
	return db, err
}
