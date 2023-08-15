package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"git.garena.com/shopee/recommend/srec_protos/srec/go_gen_code/srec_proto"
	"io"
	"log"
	"net/http"
)

// recall ...
type RecallResponse struct {
	Version string       `json:"version"`
	Status  int32        `json:"status"`
	Message string       `json:"message"`
	Value   []RecallData `json:"value"`
}

type RecallData struct {
	Items []RecallItem `json:"items"`
}

type RecallItem struct {
	ItemId uint64  `json:"item_id"`
	Score  float32 `json:"score"`
}

// Detail ...
type DetailRequest struct {
	Type       int
	Parameters DetailParameter
}

type DetailParameter struct {
	Ids []uint64
}

type DetailResponse struct {
	Data DetailData
}

type DetailData struct {
	Items []ItemInfo
}

type ItemInfo struct {
	Itemid         uint64
	Shopid         uint64
	Price          float64
	Recommendable  bool
	GlobalCat      uint32
	GlobalSubCat   uint32
	GlobalThirdCat uint32
}

const (
	PDPURLTemp = "http://pdp-recall-k8s.recommend.shopee.io/v1/GetModelItems?model=pop_v2@%s&country=%s&key=%s"
	PDPSecKey  = "you_may_also_like"
	model      = "pop_v2"
	region     = "CL"
	fetchNum   = 500

	bundle       = "product_detail_page"
	ItemInfoTemp = "EmergencyStaticFile,BND:%s,QUE:%s,QUES:%s,SECKEY:%s"
)

const (
	DetailURL = "http://viz.recommend.shopee.io/api/v1/discovery/item"
	Cookies   = `_ga_7N8T3QXGY7=GS1.1.1675219045.4.0.1675219045.0.0.0; _ga_QMDMFHB4T7=GS1.1.1683086752.102.1.1683086752.0.0.0; _ga_8SBPY6GX69=GS1.1.1684328703.616.0.1684328703.0.0.0; _ga_XWL9LM2WJ1=GS1.1.1684328703.590.0.1684328703.0.0.0; _ga_3KMTQ5WCQH=GS1.1.1684328703.590.0.1684328703.0.0.0; _ga_N3F0MMP8D4=GS1.1.1684978965.13.1.1684979735.0.0.0; space_auth_live=MTY4NTQxMjQyNXxOd3dBTkZwUU56VktOVk5TUmtsTlF6VktOakpGVVVsQldVWlFTbFpHV1ZwWVZWSk5Oa2hNTWpKWVdFUTJVVVJIVlZReU1reFhWMEU9fJ6FQlsjUQuOBW4YOjsl_LrSNnhDg92EKOY0GpmoJJhB; _ga_6VXJESRVZJ=GS1.1.1685436163.382.0.1685436163.0.0.0; _ga=GA1.2.959695384.1668505378; _ga_VPBMX0QP83=GS1.1.1685513681.88.0.1685513681.60.0.0; _ga_V8CYP9D8PH=GS1.1.1685512760.5.1.1685514415.0.0.0; SSO_C=svbvx3n86jhx4nso3bvb38vxrom94yb6smr3o6fy; SSO_A=15d8bgclp9ankd61njx3dh9ghzyx2xqmbmr3pany; _ga_N1FF831FCD=GS1.1.1685519857.550.1.1685520018.0.0.0`
)

var (
	client = &http.Client{}
)

func getItemDetail(recallItem *RecallItem) (*srec_proto.RUnit, error) {
	req := DetailRequest{
		Type: 0,
		Parameters: DetailParameter{
			Ids: []uint64{recallItem.ItemId},
		},
	}
	reqData, err := json.Marshal(req)
	if err != nil {
		return nil, errors.New("req marshal failed")
	}

	request, _ := http.NewRequest(http.MethodPost, DetailURL, bytes.NewBuffer(reqData))
	request.Header.Set("Cookie", Cookies)

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	detailResp := &DetailResponse{}
	if err := json.Unmarshal(respData, detailResp); err != nil {
		return nil, err
	}

	item := detailResp.Data.Items[0]

	runit := &srec_proto.RUnit{
		Key:           fmt.Sprintf("item::%d", item.Itemid),
		DataType:      "item",
		Score:         recallItem.Score,
		Recommendable: item.Recommendable,
		Info:          fmt.Sprintf(ItemInfoTemp, bundle, model, model, PDPSecKey),
		Label:         fmt.Sprintf("item::%d", item.Itemid),
		Itemid:        item.Itemid,
		Shopid:        item.Shopid,
		Catid:         item.GlobalCat,
		L2Catid:       item.GlobalSubCat,
		L3Catid:       item.GlobalThirdCat,
		Price:         item.Price,
	}
	return runit, nil
}

func main() {

	url := fmt.Sprintf(PDPURLTemp, region, region, region)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	recallData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("recall: %v\n", string(recallData))

	recallResp := &RecallResponse{}
	if err = json.Unmarshal(recallData, recallResp); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("recall: %v\n", *recallResp)

	mixerResp := srec_proto.RecommendResponse{}
	section := srec_proto.Section{}
	data := make([]*srec_proto.RUnit, 0, fetchNum)

	if len(recallResp.Value) == 0 {
		fmt.Println("empty recallResp")
		return
	}

	count := 1
	for _, recallItem := range recallResp.Value[0].Items {
		if count > fetchNum {
			break
		}
		runit, err := getItemDetail(&recallItem)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, runit)
		count++
	}

	section.Key = PDPSecKey
	section.Total = uint32(len(data))
	section.Data = data

	mixerResp.Sections = append(mixerResp.Sections, &section)

	//fmt.Println(url)

	marshal, _ := json.Marshal(mixerResp)
	fmt.Printf("%v", string(marshal))
}
