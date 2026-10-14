package main

import (
	"encoding/json"
	"fmt"
	"git.garena.com/shopee/recommend/srec_protos/srec/go_gen_code/srec_proto"
	"io"
	"log"
	"net/http"
)

type RecallResponse struct {
	List []Item
}

type Item struct {
	Catid         uint32
	Itemid        uint64
	Shopid        uint64
	Score         float32
	DataType      string
	Price         float64
	Key           string
	Recommendable bool
	CategoryInfo  Category
}

type Category struct {
	L1cat uint32
	L2cat uint32
	L3cat uint32
}

const (
	DDURLTemp = "http://dd-recall-k8s.recommend.shopee.io/v1/DebugStrategy?instance_name=%s&country=%s&fetchnum=%d"
	DDSecKey  = "dd_main_sec"
	strategyA = "BayesStrategy_instance_order_2_exp_quota_500"
	region    = "TW"
	count     = 500

	bundle       = "daily_discover_main"
	ItemInfoTemp = "EmergencyStaticFile,BND:%s,QUE:%s,QUES:%s,SECKEY:%s"
)

func main() {
	url := fmt.Sprintf(DDURLTemp, strategyA, region, count)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	recallResp := &RecallResponse{}
	err = json.Unmarshal(bytes, recallResp)
	if err != nil {
		log.Fatal(err)
	}

	size := len(recallResp.List)

	mixerResp := srec_proto.RecommendResponse{}
	section := srec_proto.Section{}
	data := make([]*srec_proto.RUnit, 0, size)

	info := fmt.Sprintf(ItemInfoTemp, bundle, strategyA, strategyA, DDSecKey)

	for _, item := range recallResp.List {
		runit := &srec_proto.RUnit{
			Key:           item.Key,
			DataType:      item.DataType,
			Score:         item.Score,
			Recommendable: item.Recommendable,
			Info:          info,
			Label:         item.Key,
			Itemid:        item.Itemid,
			Shopid:        item.Shopid,
			Catid:         item.Catid,
			L2Catid:       item.CategoryInfo.L2cat,
			L3Catid:       item.CategoryInfo.L3cat,
			Price:         item.Price,
		}
		data = append(data, runit)
	}
	section.Key = DDSecKey
	section.Total = uint32(len(data))
	section.Data = data

	mixerResp.Sections = append(mixerResp.Sections, &section)

	//marshal, _ := json.Marshal(mixerResp)

	fmt.Println(url)
	//fmt.Printf("%v", string(marshal))
}
