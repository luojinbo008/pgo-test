package handler

import (
	"net/http"
	log"github.com/cihub/seelog"
	"github.com/penguinn/pgo-test/web/model"
	"fmt"
	"encoding/json"
	"github.com/penguinn/pgo/app"
)

type TestStruct struct {
	Id		 	float32 `json:"id"`
	Value 		float32 `json:"value"`
}

type Controller int

func (p *Controller)DefaultHandler(w http.ResponseWriter, req *http.Request) {
	data, err := model.UserDeviceModel.SelectDevicenuByUseridSource(57076573, 1)
	if err != nil{
		log.Error(err)
	}
	fmt.Println(data)
	w.Write([]byte(data.DeviceNum))
}


func (p *Controller)TestHandler(w http.ResponseWriter, req *http.Request) {
	heheStruct, err := model.HeHeModel.GetOneById(1.0)
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