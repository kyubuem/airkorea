package measure

import (
	"fmt"
)

//unityAirEnvrnIdexSnstiveAboveMsrstnListItem 은 통합대기환경 지수 나쁨 이상측정소를 조회 했을 경우 제공되는
//데이터를 저장히기 위한 구조체이다.
type unityAirEnvrnIdexSnstiveAboveMsrstnListItem struct {
	StationName string `xml:"stationName"` //측정소명
	Address     string `xml:"addr"`        //측정소 주소
}

//unityAirEnvrnIdexSnstiveAboveMsrstnListRequest 은 통합대기환경 지수 나쁨 이상측정소를
//조회 하기 위한 구조체이다. 별도의 파라미터값이 필요없다.
type unityAirEnvrnIdexSnstiveAboveMsrstnListRequest struct {
	PageNo    int64 //페이지 번호
	NumOfRows int64 //한 페이지 결과 수
}

//Op2ResponseItem 은 unityAirEnvrnIdexSnstiveAboveMsrstnListItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op2ResponseItem unityAirEnvrnIdexSnstiveAboveMsrstnListItem

//Op2Request 은 unityAirEnvrnIdexSnstiveAboveMsrstnListRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op2Request unityAirEnvrnIdexSnstiveAboveMsrstnListRequest

//Op2Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op2Operator struct {
	Req Op2Request
}

//Request는 unityAirEnvrnIdexSnstiveAboveMsrstnListRequest 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op2Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 2번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op2Operator) RequestString() string {
	return fmt.Sprintf("%s/%s?pageNo=%d&numOfRows=%d",
		serviceName, operationMap[2],
		o.Req.PageNo, o.Req.NumOfRows)
}

//Item은 Op2ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op2Operator) Item() interface{} {
	return &Op2ResponseItem{}
}
