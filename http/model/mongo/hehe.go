package mongo

import (
	"gopkg.in/mgo.v2"
	"github.com/penguinn/pgo/app"
	"github.com/penguinn/pgo-test/http/dao/mongo"
)

type DBNameModel int

func NewDBNameDao() *mongo.DBNameDao {
	dbNameModel := new(DBNameModel)
	db := app.UseModel("mongo", dbNameModel.ConnName(), true).(*mgo.Database)
	return &mongo.DBNameDao{DB:db}
}

func (p *DBNameModel) ConnName() string {
	return "default"
}