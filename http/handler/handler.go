package handler

import (
	"net/http"
	log"github.com/cihub/seelog"
	"fmt"
	"encoding/json"
	"github.com/penguinn/pgo/app"
	"github.com/penguinn/pgo-test/http/model/mysql"
	"github.com/penguinn/pgo-test/http/model/mongo"
)

type TestStruct struct {
	Id		 	float32 `json:"id"`
	Value 		float32 `json:"value"`
}

type Controller int

func (p *Controller)DefaultHandler(w http.ResponseWriter, req *http.Request) {
	mDB := mysql.NewDBNameDao()
	data, err := mDB.SelectDevicenuByUseridSource(57082592, 1)
	if err != nil{
		log.Error(err)
	}
	fmt.Println(data)
	w.Write([]byte(data.DeviceNum))
}


func (p *Controller)TestHandler(w http.ResponseWriter, req *http.Request) {
	mongoDB := mongo.NewDBNameDao()
	heheStruct, err := mongoDB.GetOneById(1.0)
	if err != nil{
		log.Error(err)
	}
	heheJson, err := json.Marshal(heheStruct)
	if err != nil{
		log.Error(err)
	}

	redi, err := app.GetRedis("default")
	if err != nil{
		log.Error("get redis error,", err)
	}
	redisString, err:= redi.Get("sjy").Result()
	if err != nil{
		log.Error(err)
	}
	w.Write([]byte(heheJson))
	w.Write([]byte(redisString))
}