package measure

import (
	"fmt"
	"strings"
)

//minuDustFrcstDspthItem 은 대기질 예보통보를 조회했을 경우 제공되는
//데이터를 저장하기 위한 구조체이다.
type minuDustFrcstDspthItem struct {
	DataTime      string `xml:"dataTime,omitempty"`      //통보시간
	InformCode    string `xml:"informCode,omitempty"`    //통보코드
	InformOverall string `xml:"informOverall,omitempty"` //예보개황
	InformCause   string `xml:"informCause,omitempty"`   //발생원인
	InformGrade   string `xml:"informGrade,omitempty"`   //예보등급
	ActionKnack   string `xml:"actionKnack,omitempty"`   //행동요령
	ImageUrl1     string `xml:"imageUrl1,omitempty"`     //시간대별 예측 모댈 결과 사진(6:00, 12:00, 18:00, 24:00 KST)
	ImageUrl2     string `xml:"imageUrl2,omitempty"`     //시간대별 예측 모댈 결과 사진(6:00, 12:00, 18:00, 24:00 KST)
	ImageUrl3     string `xml:"imageUrl3,omitempty"`     //시간대별 예측 모댈 결과 사진(6:00, 12:00, 18:00, 24:00 KST)
	ImageUrl4     string `xml:"imageUrl4,omitempty"`     //시간대별 예측 모댈 결과 사진(6:00, 12:00, 18:00, 24:00 KST)
	ImageUrl5     string `xml:"imageUrl5,omitempty"`     //시간대별 예측 모댈 결과 사진(6:00, 12:00, 18:00, 24:00 KST)
	ImageUrl6     string `xml:"imageUrl6,omitempty"`     //시간대별 예측 모댈 결과 사진(6:00, 12:00, 18:00, 24:00 KST)
	ImageUrl7     string `xml:"imageUrl7,omitempty"`     //미세먼지(PM10) 한반도 대기질 예측모델결과 애니메이션 이미지
	ImageUrl8     string `xml:"imageUrl8,omitempty"`     //미세먼지(PM2.5) 한반도 대기질 예측모델결과 애니메이션 이미지
	ImageUrl9     string `xml:"imageUrl9,omitempty"`     //오존(O3) 한반도 대기질 예측모델결과 애니메이션 이미
	InformData    string `xml:"informData,omitempty"`    //예측통보시간
}

//minuDustFrcstDspthRequest 은 대기질 예보통보를 조회하기 위한 측정소명 및 주기를
//서버에 전송하기 위한 구조체이다.
type minuDustFrcstDspthRequest struct {
	SearchDate string //통보시간 검색(조회 날짜 입력이 없을 경우 한달동안 예보통보 발령 날짜의 리스트 정보를 확인)
	InformCode string //통보코드 검색 (PM10 : 미세먼지, PM25 : 초미세먼지, O3 : 오존)
	PageNo     int64  //페이지 번호(조회날짜로 검색 시 사용 안함)
	NumOfRows  int64  //한 페이지 결과수 (조회 날짜로 검색 시 사용 안함)
	Version    string //새로운 애니메이션이 포함된 API호출시 버전 사용 (&ver=1.1)
}

//Op4ResponseItem 은 minuDustFrcstDspthItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op4ResponseItem minuDustFrcstDspthItem

//Op4Request 은 minuDustFrcstDspthRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op4Request minuDustFrcstDspthRequest

//Op4Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op4Operator struct {
	Req Op4Request
}

//Request는 minuDustFrcstDspthRequest 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op4Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 4번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op4Operator) RequestString() string {
	url := fmt.Sprintf("%s/%s?searchDate=%s&informCode=%s&pageNo=%d&numOfRows=%d",
		serviceName, operationMap[4], o.Req.SearchDate, o.Req.InformCode, o.Req.PageNo, o.Req.NumOfRows)
	if strings.Compare(o.Req.Version, "1.1") == 0 {
		url = fmt.Sprintf("%s&ver=%s", url, o.Req.Version)
	}
	return url
}

//Item은 Op1ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op4Operator) Item() interface{} {
	return &Op4ResponseItem{}
}
