package station

import "fmt"

//nearbyMsrstnListResponseItem 은 근접측정소 목록 조회했을 경우 제공되는
//데이터를 저장하기 위한 구조체이다.
type nearbyMsrstnListResponseItem struct {
	StationName string  `xml:"stationName"` //측정소 이름
	Address     string  `xml:"addr"`        //측정소가 위치한 주소
	Tm          float64 `xml:"tm"`          //요청한 TM좌표와 측정소간의 거리(KM 단위)
}

//nearbyMsrstnListRequest  은 근접측정소 목록을 조회하기 위한 Tm측정방식의 좌표를
//서버에 전송하기 위한 구조체이다.
type nearbyMsrstnListRequest struct {
	TmX       string //TM측정방식 X좌표, ver1.0일경우 도로명주소API X좌표를 입력해야 한다.
	TmY       string //TM측정방식 Y좌표, ver1.0일경우 도로명주소API Y좌표를 입력해야 한다.
	PageNo    int64  //페이지 번호
	NumOfRows int64  //한 페이지 결과 수
	Version   string //버전 1.0을 호출할 경우 도로명주소 검색(juso.go.kr) API가 제공하는 API의 X,Y 좌표로 가까운 측정소를 표출한다.
}

//Op1ResponseItem 은 nearbyMsrstnListResponseItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op1ResponseItem nearbyMsrstnListResponseItem

//Op1Request 은 nearbyMsrstnListRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op1Request nearbyMsrstnListRequest

//Op1Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op1Operator struct {
	Req Op1Request
}

//Request는 nearbyMsrstnListRequest의 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op1Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 1번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op1Operator) RequestString() string {
	url := fmt.Sprintf("%s/%s?tmX=%s&tmY=%s",
		serviceName, operationMap[1], o.Req.TmX, o.Req.TmY)
	if o.Req.Version == "1.0" {
		url = fmt.Sprintf("%s&ver=%s", url, o.Req.Version)
	}
	return url
}

//Item은 Op1ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op1Operator) Item() interface{} {
	return &Op1ResponseItem{}
}
