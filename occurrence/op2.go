package occurrence

import "fmt"

//ylwsndAdvsryOccrrncInfoResponseItem 은 황사주의보 발생정보를 조회했을 때 서버에서
//return 해주는 데이터 집합이다.
type ylwsndAdvsryOccrrncInfoResponseItem struct {
	DataTime string `xml:"dataTime"` //발령 날짜 연-월-일
	TmCnt    int64  `xml:"tmCnt"`    //황사 발생 회차
	TmArea   string `xml:"tmArea"`   //황사 발생 지역
}

//ylwsndAdvsryOccrrncInfoRequest 은 황사주의보 발생정보를 조회할 때 사용하는 데이터 집합이다.
type ylwsndAdvsryOccrrncInfoRequest struct {
	Year      int64 // 발령 연도
	PageNo    int64 // 페이지 번호
	NumOfRows int64 // 한 페이지 결과 수
}

//Op2ResponseItem 은 ylwsndAdvsryOccrrncInfoResponseItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op2ResponseItem ylwsndAdvsryOccrrncInfoResponseItem

//Op2Request 은 ylwsndAdvsryOccrrncInfoRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op2Request ylwsndAdvsryOccrrncInfoRequest

//Op2Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op2Operator struct {
	Req Op2Request
}

//Request는 ylwsndAdvsryOccrrncInfoRequest 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op2Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 2번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op2Operator) RequestString() string {
	return fmt.Sprintf("%s/%s?year=%d&pageNo=%d&numOfRows=%d",
		serviceName, operationMap[2],
		o.Req.Year, o.Req.PageNo, o.Req.NumOfRows)
}

//Item은 Op1ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op2Operator) Item() interface{} {
	return &Op2ResponseItem{}
}
