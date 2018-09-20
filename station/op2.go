package station

import (
	"fmt"
	"strings"
)

//msrstnListResponseItem 은 측정소 목록을 조회하여 리턴한다.
type msrstnListResponseItem struct {
	StationName string  `xml:"stationName"` //측정소 이름
	Address     string  `xml:"addr"`        //측정소가 위치한 주소
	Year        int64   `xml:"year"`        //측정소 설치년도
	Oper        string  `xml:"oper"`        //측정소 관리기관 이름
	Photh       string  `xml:"photo"`       //측정소 이미지
	Vrml        string  `xml:"vrml"`        //측정소 주변 전경
	Map         string  `xml:"map"`         //측정소가 설치된 장소 지도 이미지
	ManageName  string  `xml:"mangName"`    //측정망
	Item        string  `xml:"item"`        //측정소 측정 항목 (SO2, CO, O3, NO2, PM10)
	DmX         float64 `xml:"dmX"`         //WGS84기반 X좌표
	DmY         float64 `xml:"dmY"`         //WGS84기반 Y좌표
}

//msrstnListRequest 은 측정소 목록 조회시 필요한 데이터 구조체이다.
type msrstnListRequest struct {
	StationName string //측정소 이름
	Address     string //주소 이름
	PageNo      int64  //페이지 번호
	NumOfRows   int64  //한 페이지 결과 수
}

//Op2ResponseItem 은 msrstnListResponseItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op2ResponseItem msrstnListResponseItem

//GetSupportItems 은 측정소 측정 항목의 리스트를 리턴한다.
//측정소 별로 지원하는 측정정보가 다르며 기본적으로 SO2, CO, O3, NO2, PM10등이 지원된다.
func (o Op2ResponseItem) GetSupportItems() []string {
	return strings.Split(o.Item, ",")
}

//Op2Request 은 msrstnListRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op2Request msrstnListRequest

//Op2Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op2Operator struct {
	Req Op2Request
}

//Request는 msrstnListRequest의 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op2Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 2번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op2Operator) RequestString() string {
	url := fmt.Sprintf("%s/%s?addr=%s&stationName=%s&pageNo=%d&numOfRows=%d",
		serviceName, operationMap[2], o.Req.Address, o.Req.StationName, o.Req.PageNo, o.Req.NumOfRows)
	return url
}

//Item은 Op1ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op2Operator) Item() interface{} {
	return &Op2ResponseItem{}
}
