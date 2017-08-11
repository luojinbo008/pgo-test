package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DBNameDao struct {
	DB 		*mgo.Database
}

type HeHe struct {
	Id		 	float32 `json:"id"`
	Value 		float32 `json:"value"`
}

func (p *DBNameDao) GetOneById(id float32) (hehe HeHe, err error){
	hehe = HeHe{}
	err = p.DB.C("hehe").Find(bson.M{"id":id}).One(&hehe)
	return
}
