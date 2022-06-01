package db

import (
	"douyin/pkg/constants"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)
	if err != nil {
		panic(err)
	}

	// DB.SingularTable(true)

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	m := DB.Migrator()
	if m.HasTable(&User{}) && m.HasTable(&Video{}) {
		return
	}

	if !m.HasTable(&User{}) {
		if err = m.CreateTable(&User{}); err != nil {
			panic(err)
		}
	}
	if !m.HasTable(&Video{}) {
		if err = m.CreateTable(&Video{}); err != nil {
			panic(err)
		}
	}
}
