package station

import "fmt"

//tMStdrCrdntResponseItem 은 TM기준좌표를 조회했을 때 서버에서
//return 해주는 데이터 집합이다.
type tMStdrCrdntResponseItem struct {
	SidoName string  `xml:"sidoName"` //시도명
	SggName  string  `xml:"sggName"`  //시군구명
	UmdName  string  `xml:"umdName"`  //읍면동명
	TmX      float64 `xml:"tmX"`      //TM측정방식 X좌표
	TmY      float64 `xml:"tmY"`      //TM측정방식 Y좌표
}

//tMStdrCrdntRequest 은 TM기준좌표를 조회할 때 사용하는 데이터 집합이다.
type tMStdrCrdntRequest struct {
	UmdName   string //읍면동명
	PageNo    int64  //페이지 번호
	NumOfRows int64  //한 페이지 결과 수
}

//Op3ResponseItem 은 tMStdrCrdntResponseItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op3ResponseItem tMStdrCrdntResponseItem

//Op3Request 은 tMStdrCrdntRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op3Request tMStdrCrdntRequest

//Op3Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op3Operator struct {
	Req Op3Request
}

//Request는 tMStdrCrdntRequest의 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op3Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 1번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op3Operator) RequestString() string {
	url := fmt.Sprintf("%s/%s?umdName=%s&pageNo=%d&numOfRows=%d",
		serviceName, operationMap[3], o.Req.UmdName, o.Req.PageNo, o.Req.NumOfRows)
	return url
}

//Item은 Op3ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op3Operator) Item() interface{} {
	return &Op3ResponseItem{}
}
