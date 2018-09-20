package occurrence

import "fmt"

//ozAdvsryOccrrncInfoResponseItem 은 오존주의보 발생정보를 조회했을 때 서버에서
//return 해주는 데이터 집합이다.
type ozAdvsryOccrrncInfoResponseItem struct {
	DataTime     string  `xml:"dataTime"`     //발령 날짜
	DistrictName string  `xml:"districtName"` //발령 지역 이름
	MoveName     string  `xml:"moveName"`     //발령 권역 이름
	IssueTime    string  `xml:"issueTime"`    //발령 시간
	IssueVal     float64 `xml:"issueVal"`     //발령시 오존 농도 (단위 : ppm)
	ClearTime    string  `xml:"clearTime"`    //해제 시간
	ClearVal     float64 `xml:"clearVal"`     //해제시 오존 농도 (단위 : ppm)
	MaxVal       float64 `xml:"maxVal"`       //오존 최고 농도 (단위 : ppm)
}

//ozAdvsryOccrrncInfoRequest 은 오존주의보 발생정보를 조회할 때 사용하는 데이터 집합이다.
type ozAdvsryOccrrncInfoRequest struct {
	Year      int64 //측정 연도
	PageNo    int64 //페이지 번호
	NumOfRows int64 //한 페이지 결과 수
}

//Op1ResponseItem 은 ozAdvsryOccrrncInfoResponseItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op1ResponseItem ozAdvsryOccrrncInfoResponseItem

//Op1Request 은 ozAdvsryOccrrncInfoRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op1Request ozAdvsryOccrrncInfoRequest

//Op1Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op1Operator struct {
	Req Op1Request
}

//Request는 ozAdvsryOccrrncInfoRequest 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op1Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 1번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op1Operator) RequestString() string {
	return fmt.Sprintf("%s/%s?year=%d&pageNo=%d&numOfRows=%d",
		serviceName, operationMap[1],
		o.Req.Year, o.Req.PageNo, o.Req.NumOfRows)
}

//Item은 Op1ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op1Operator) Item() interface{} {
	return &Op1ResponseItem{}
}
