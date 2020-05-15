package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go-template-cctv/xmlfmt"
	"text/template"
)

const (
	OBJECT_SERIES  = "Series"
	OBJECT_PICTURE = "Picture"
	OBJECT_PROGRAM = "Program"
	OBJECT_MOVIE   = "Movie"
)

func main() {
	xml := generateSOAPRequest(populateRequest())
	fmt.Println(xml)
}

type Request struct {
	Objects  []Object
	Mappings []Mapping
}

type Object struct {
	ID          string // 接口中的唯一标识
	ElementType string // 元素类型
	Action      string // 操作类型 注册－REGIST 更新－UPDATE 删除－DELETE
	Code        string // 全局唯一标识
	Series      Series
	Picture     Picture
	Movie       Movie
}

type Mapping struct {
	ID               string
	Action           string
	ParentType       string
	ElementType      string
	ParentID         string
	ElementID        string
	ParentCode       string
	ElementCode      string
	Type             string
	ValidStart       string
	ValidEnd         string
	Sequence         string
	Result           string
	ErrorDescription string
}

// Objects模板生成
func (o Object) GetObjectTemplate(elementType string, values interface{}) string {
	var objectTemplate string
	switch elementType {

	case OBJECT_SERIES:
		objectTemplate = "./template/object_series.gohtml"
	case OBJECT_PROGRAM:
		objectTemplate = "./template/object_program.gohtml"
	case OBJECT_MOVIE:
		objectTemplate = "./template/object_movie.gohtml"
	case OBJECT_PICTURE:
		objectTemplate = "./template/object_picture.gohtml"
	}

	// Template with struct
	template, err := template.ParseFiles(objectTemplate)
	if err != nil {
		fmt.Println("Error while marshling object. %s ", err.Error())
		return ""
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	switch elementType {

	case OBJECT_SERIES:
		err = template.Execute(doc, values.(Series))
	case OBJECT_PROGRAM:
		err = template.Execute(doc, values.(Program))
	case OBJECT_MOVIE:
		err = template.Execute(doc, values.(Movie))
	case OBJECT_PICTURE:
		err = template.Execute(doc, values.(Picture))
	}

	if err != nil {
		fmt.Println("template.Execute error. %s ", err.Error())
		return ""
	}
	return doc.String()
}

func (m Mapping) GetMappingTemplate() string {

	// Template with struct
	template, err := template.ParseFiles("./template/mapping.gohtml")
	if err != nil {
		fmt.Println("Error while marshling object. %s ", err.Error())
		return ""
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, m)
	if err != nil {
		fmt.Println("template.Execute error. %s ", err.Error())
		return ""
	}
	return doc.String()
}

// 节目
type Program struct {
	Name                 string
	OrderNumber          string
	OriginalName         string
	SortName             string
	SearchName           string
	ActorDisplay         string
	WriterDisplay        string
	OriginalCountry      string
	Language             string
	ReleaseYear          string
	OrgAirDate           string
	LicensingWindowStart string
	LicensingWindowEnd   string
	DisplayAsNew         string
	DisplayAsLastChance  string
	Macrovision          string
	Description          string
	PriceTaxIn           string
	Status               string
	SourceType           string
	SeriesFlag           string
	Kpeople              string
	Director             string
	ScriptWriter         string
	Compere              string
	Guest                string
	Reporter             string
	OPIncharge           string
	CpId                 string
	CpName               string
	KindDisplay          string
	Type                 string
	Tags                 string
}

// 连续剧
type Series struct {
	Name                 string // 节目名称
	OrderNumber          string // 节目订购编号
	OriginalName         string // 原名
	SortName             string // 索引发布时间供界面排序
	SearchName           string // 搜索名称供界面搜索
	OrgAirDate           string // 首播时间
	LicensingWindowStart string // 有效开始时间 (YYYYMMDDHH24MiSS)
	LicensingWindowEnd   string // 有效结束时间 (YYYYMMDDHH24MiSS)
	DisplayAsNew         string // 新到天数
	DisplayAsLastChance  string // 剩余天数
	Macrovision          string // 拷贝保护标志 0:无拷贝保护 1:有拷贝保护
	Type                 string
	Keywords             string
	Tags                 string
	Price                string
	VolumnCount          string
	Status               string // 状态标志 0:失效 1:生效
	Description          string // 节目描述
	Kpeople              string // 主要人物
	Director             string // 导演
	ScriptWriter         string // 编剧
	Compere              string // 节目主持人
	Guest                string // 受访者
	Reporter             string // 记者
	OPIncharge           string // 其他责任人
	SeriesType           string // 0: 影片 1: 单集
	OriginalCountry      string // 国家地区
	Language             string // 语言
	ReleaseYear          string // 上映年份
	CpId                 string
	CpName               string
	KindDisplay          string
}

// 图片信息
type Picture struct {
	FileURL     string
	Description string
}

// 媒体内容信息
type Movie struct {
	Type             string
	FileURL          string
	SourceDRMType    string
	DestDRMType      string
	AudioType        string
	ScreenFormat     string
	ClosedCaptioning string
}

func populateRequest() *Request {
	req := Request{}
	req.Objects = []Object{}
	req.Mappings = []Mapping{}

	req.Objects = append(req.Objects, Object{
		ID:          "6660",
		ElementType: "Series",
		Action:      "REGIST",
		Code:        "aff899e9238e49f38ec520c6495e6450",
		Series: Series{
			Name:                 "探险活宝第三季",
			OrderNumber:          "",
			OriginalName:         "",
			SortName:             "",
			SearchName:           "txhbdsj",
			OrgAirDate:           "20110711",
			LicensingWindowStart: "20170101000000",
			LicensingWindowEnd:   "21170101000000",
			DisplayAsNew:         "",
			DisplayAsLastChance:  "",
			Macrovision:          "0",
			Type:                 "少儿",
			Keywords:             "",
			Tags:                 "冒险,儿童,动画",
			Price:                "0",
			VolumnCount:          "22",
			Status:               "1",
			Description:          "探险活宝全集故事讲述的是两个好朋友—人类小男孩Finn（台译阿宝、直译芬）和拥有魔力的狗Jake（台译老皮、直译杰克）在LandofOoo（台译哇赛秘境、直译噢噢噢大陆、）的冒险旅程。",
			Kpeople:              "杰里米·沙达,约翰·迪·玛吉欧",
			Director:             "LarryLeichliter",
			ScriptWriter:         "",
			Compere:              "",
			Guest:                "",
			Reporter:             "",
			OPIncharge:           "",
			SeriesType:           "1",
			OriginalCountry:      "美国",
			Language:             "国语",
			ReleaseYear:          "2011",
			CpId:                 "138",
			CpName:               "行云媒资",
			KindDisplay:          "冒险,儿童,动画",
		},
	})

	req.Objects = append(req.Objects, Object{
		ID:          "13631",
		ElementType: "Picture",
		Action:      "REGIST",
		Code:        "ec43663a4e944043b0ff080ae82a65a2",
		Picture: Picture{
			FileURL:     "http://u-storage.oss.seeingtv.com/138_7a0baae5fdb24178b38b88cc3b11ae6b.jpg",
			Description: "海报",
		},
	})

	req.Mappings = append(req.Mappings, Mapping{
		ID:          "00800000003020190710101020183133",
		Action:      "REGIST",
		ParentType:  "Picture",
		ElementType: "Series",
		ParentID:    "13631",
		ElementID:   "6660",
		ParentCode:  "ec43663a4e944043b0ff080ae82a65a2",
		ElementCode: "aff899e9238e49f38ec520c6495e6450",
		Type:        "1",
		Sequence:    "90",
	})
	return &req
}

func generateSOAPRequest(req *Request) string {
	// Using the var getTemplate to construct request
	template, err := template.ParseFiles("./template/cctv_album.gohtml")
	if err != nil {
		fmt.Println("Error while marshling object. %s ", err.Error())
		return ""
	}

	doc := &bytes.Buffer{}
	// Replacing the doc from template with actual req values
	err = template.Execute(doc, req)
	if err != nil {
		fmt.Println("template.Execute error. %s ", err.Error())
		return ""
	}

	buffer := &bytes.Buffer{}
	encoder := xml.NewEncoder(buffer)
	err = encoder.Encode(doc.String())
	if err != nil {
		fmt.Println("encoder.Encode error. %s ", err.Error())
		return ""
	}

	return xmlfmt.FormatXML(doc.String(), "", "")
}
