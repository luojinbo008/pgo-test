package model

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"github.com/penguinn/pgo/app"
)

var HeHeModel = new(HeHe)

type HeHe struct {
	Id		 	float32 `json:"id"`
	Value 		float32 `json:"value"`
}

func (p *HeHe) GetOneById(id float32) (hehe HeHe, err error){
	hehe = HeHe{}
	mongodb := app.UseModel("mongo", p, false)
	err = mongodb.(*mgo.Database).C(p.CollectionName()).Find(bson.M{"id":id}).One(&hehe)
	return
}

func (p *HeHe) CollectionName() string {
	return "hehe"
}

func (p *HeHe) ConnName() string {
	return "default"
}