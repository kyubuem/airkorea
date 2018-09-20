package measure

import (
	"fmt"
)

//ctprvnMesureLIstItem 은 시도별 실시간 평균 정보를 조회 햇을 경우 제공되는
//데이터를 저장하기 위한 구조체이다.
type ctprvnMesureLIstItem struct {
	DataTime  string `xml:"dataTime"`  //평균자로 조회일 연-월-일
	ItemCode  string `xml:"itemCode"`  //조회항목 구분(SO2, CO, O3, NO2, PM10, PM25)
	DataGubun string `xml:"dataGubun"` //조회 자료 구분 (시간평균, 일평균)
	Seoul     int64  `xml:"seoul"`     //서울 지역 평균
	Busan     int64  `xml:"busan"`     //부산 지역 평균
	Daegu     int64  `xml:"daegu"`     //대구 지역 평균
	Incheon   int64  `xml:"incheon"`   //인천 지역 평균
	Gwangju   int64  `xml:"gwangju"`   //광주 지역 평균
	Daejeon   int64  `xml:"daejeon"`   //대전 지역 평균
	Ulsan     int64  `xml:"ulsan"`     //울산 지역 평균
	Gyeonggi  int64  `xml:"gyeonggi"`  //경기 지역 평균
	Gangwon   int64  `xml:"gangwon"`   //강원 지역 평균
	Chungbuk  int64  `xml:"chungbuk"`  //충북 지역 평균
	Chungnam  int64  `xml:"chungnam"`  //충남 지역 평균
	Jeonbuk   int64  `xml:"jeonbuk"`   //전북 지역 평균
	Jeonnam   int64  `xml:"jeonnam"`   //전남 지역 평균
	Gyeongbuk int64  `xml:"gyeongbuk"` //경북 지역 평균
	Gyeongnam int64  `xml:"gyeongnam"` //경남 지역 평균
	Jeju      int64  `xml:"jeju"`      //제주 지역 평균
	Sejong    int64  `xml:"sejong"`    //세종 지역 평균
}

//ctprvnMesureLIstRequest 은 시도별 실시간 평균 정보를 조회 하기 위한
//데이터 구분, 데이터 기간, 측정 항목을 서버에 전송하기 위한 구조체이다.
type ctprvnMesureLIstRequest struct {
	DataGubun       string //요청 자료 구분(시간평균 : HOUR, 일평균 : DAILY)
	SearchCondition string //요청 데이터기간 (일주일 : WEEK, 한달 : MONTH), 요청 자료 구분이 HOUR일경우 SearchCondition은 입려되더라도 무시된다.
	ItemCode        string //측정항목 구분(SO2, CO, O3, NO2, PM10, PM25)
	PageNo          int64  // 페이지 번호
	NumOfRows       int64  //한페이지 결과수
}

//Op5ResponseItem 은 ctprvnMesureLIstItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op5ResponseItem ctprvnMesureLIstItem

//Op5Request 은 ctprvnMesureLIstRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op5Request ctprvnMesureLIstRequest

//Op5Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op5Operator struct {
	Req Op5Request
}

//Request는 ctprvnMesureLIstRequest 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op5Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 5번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op5Operator) RequestString() string {
	url := fmt.Sprintf("%s/%s?itemCode=%s&dataGubun=%s",
		serviceName, operationMap[5], o.Req.ItemCode, o.Req.DataGubun)
	if o.Req.DataGubun == "HOUR" {
		return fmt.Sprintf("%s&pageNo=%d&numOfRows=%d", url, o.Req.PageNo, o.Req.NumOfRows)
	}
	return fmt.Sprintf("%s&searchCondition=%s&pageNo=%d&numOfRows=%d", url, o.Req.SearchCondition, o.Req.PageNo, o.Req.NumOfRows)
}

//Item은 Op5ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op5Operator) Item() interface{} {
	return &Op5ResponseItem{}
}
