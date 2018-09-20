package measure

import (
	"fmt"
	"strings"
)

//msrstnAcctoRltmMesureDnstyResponseItem 은 측정소별 실시간 측정정보를 조회했을 경우 제공되는
//데이터를 저장하기 위한 구조체이다.
type msrstnAcctoRltmMesureDnstyResponseItem struct {
	DataTime    string  `xml:"dataTime,omitempty"`    //오염도 측정 연-월-일 시간:분
	ManageName  string  `xml:"mangName,omitempty"`    //측정망 정보(국가배경, 교외대기, 도시대기, 도로변대기)
	So2Value    float64 `xml:"so2Value,omitempty"`    //아황산가스 농도
	CoValue     float64 `xml:"coValue,omitempty"`     //일산화탄소 농도
	O3Value     float64 `xml:"o3Value,omitempty"`     //오존 농도
	No2Value    float64 `xml:"no2Value,omitempty"`    //이산화질소 농도
	Pm10Value   int64   `xml:"pm10Value,omitempty"`   //미세먼지(PM10) 농도
	Pm10Value24 int64   `xml:"pm10Value24,omitempty"` //미세먼지(PM10) 24시간 예측이동 농도
	Pm25Value   int64   `xml:"pm25Value,omitempty"`   //미세먼지(PM25) 농도
	Pm25Value24 int64   `xml:"pm25Value24,omitempty"` //미세먼지(PM25) 24시간 예측이동 농도
	KhaiValue   int64   `xml:"khaiValue,omitempty"`   //통합대기환경 수치
	KhaiGrade   int64   `xml:"khaiGrade,omitempty"`   //통합대기환경 지수
	So2Grade    int64   `xml:"so2Grade,omitempty"`    //아황산가스 지수
	CoGrade     int64   `xml:"coGrade,omitempty"`     //일산화탄소 지수
	O3Grade     int64   `xml:"o3Grade,omitempty"`     //오존 지수
	No2Grade    int64   `xml:"no2Grade,omitempty"`    //이산화질소 지수
	Pm10Grade   int64   `xml:"pm10Grade,omitempty"`   //미세먼지(PM10) 24시간 등급자료
	Pm25Grade   int64   `xml:"pm25Grade,omitempty"`   //미세먼지(PM25) 24시간 등급자료
	Pm10Grade1h int64   `xml:"pm10Grade1h,omitempty"` //미세먼지(PM10) 1시간 등급자료
	Pm25Grade1h int64   `xml:"pm25Grade1h,omitempty"` //미세먼지(PM25) 1시간 등급자료
}

//msrstnAcctoRltmMesureDnstyRequest 은 측정소별 실시간 측정정보를 조회하기 위한 측정소명 및 주기를
//서버에 전송하기 위한 구조체이다.
type msrstnAcctoRltmMesureDnstyRequest struct {
	StationName string //측정소 이름
	DataTerm    string //요청 데이터 기간(1일: DAILY, 1개월: MONTH, 3개월: 3MONTH)
	PageNo      int64  //페이지 번호
	NumOfRows   int64  //한 페이지 결과 수
	Version     string //버전 미포함 : PM25 데이터 미포함, 1.0 : PM25데이터 포함, 1.1 PM10,PM25 24시간 예측이동 평균 데이터 포함, 1.2 : 측정망 정보 데이터 포함, 1.3 : PM10 PM25 1시간 등급 자료 포함
}

//Op1ResponseItem 은 msrstnAcctoRltmMesureDnstyResponseItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op1ResponseItem msrstnAcctoRltmMesureDnstyResponseItem

//Op1Request 은 msrstnAcctoRltmMesureDnstyRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op1Request msrstnAcctoRltmMesureDnstyRequest

//Op1Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op1Operator struct {
	Req Op1Request
}

//Request는 msrstnAcctoRltmMesureDnstyRequest 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op1Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 1번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op1Operator) RequestString() string {
	url := fmt.Sprintf("%s/%s?stationName=%s&dataTerm=%s&pageNo=%d&numOfRows=%d",
		serviceName, operationMap[1], o.Req.StationName, o.Req.DataTerm, o.Req.PageNo, o.Req.NumOfRows)
	if strings.Compare(o.Req.Version, "1.0") == 0 ||
		strings.Compare(o.Req.Version, "1.1") == 0 ||
		strings.Compare(o.Req.Version, "1.2") == 0 ||
		strings.Compare(o.Req.Version, "1.3") == 0 {
		url = fmt.Sprintf("%s&ver=%s", url, o.Req.Version)
	}
	return url
}

//Item은 Op1ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op1Operator) Item() interface{} {
	return &Op1ResponseItem{}
}
