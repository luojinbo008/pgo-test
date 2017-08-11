package mysql

import (
	"github.com/penguinn/pgo/app"
	"github.com/jinzhu/gorm"
	"github.com/penguinn/pgo-test/http/dao/mysql"
)

type DBNameModel int

func NewDBNameDao() *mysql.DBNameDao{
	userModel := new(DBNameModel)
	db := app.UseModel("mysql", userModel.ConnName(), true).(*gorm.DB)
	return &mysql.DBNameDao{DB:db}
}

func (p *DBNameModel) ConnName() string {
	return "default"
}
