package service

import (
	"context"
	"fmt"

	"github.com/peng225/wolper/dictionary"
	"github.com/peng225/wolper/pb"
)

type WolperServiceServerImpl struct {
	dict dictionary.Dictionary
}

func (wssi *WolperServiceServerImpl) Init(input string) {
	wssi.dict.Load(input)
}

func (wssi *WolperServiceServerImpl) Query(ctx context.Context, sreq *pb.SearchRequest) (*pb.SearchResponse, error) {
	fmt.Printf("Query requested. (key = \"%v\", include = \"%v\", exclude = \"%v\", uniq = %v)\n",
		sreq.Key, sreq.Include, sreq.Exclude, sreq.Uniq)
	var result pb.SearchResponse
	result.Words = wssi.dict.Query(sreq.Key, sreq.Include, sreq.Exclude, sreq.Uniq)
	return &result, nil
}
