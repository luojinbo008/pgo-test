package handler

import (
	"encoding/json"
	"github.com/penguinn/pgo/jsonrpc2"
	"context"
	log "github.com/cihub/seelog"
)

type Echo int

type AddParams struct {
	AddStart 	int 	`json:"addStart"`
	AddEnd 	  	int 	`json:"addEnd"`
}

type AddResult struct {
	Sum 	int    `json:"sum"`
}

type MulParams struct {
	MulStart 	int 	`json:"mulStart"`
	MulEnd 	  	int 	`json:"mulEnd"`
}

type MulResult struct {
	Product 	int    `json:"product"`
}

func(p *Echo)Add(c context.Context, params *json.RawMessage) (interface{}, *jsonrpc2.Error) {
	var param AddParams
	if err := jsonrpc2.Unmarshal(params, &param); err != nil {
		log.Error(err)
		return nil, err
	}

	return AddResult{
		Sum: param.AddStart + param.AddEnd,
	}, nil
}

func(p *Echo)Mul(c context.Context, params *json.RawMessage) (interface{}, *jsonrpc2.Error) {
	var param MulParams
	if err := jsonrpc2.Unmarshal(params, &param); err != nil {
		log.Error(err)
		return nil, err
	}

	return AddResult{
		Sum: param.MulStart * param.MulEnd,
	}, nil
}

