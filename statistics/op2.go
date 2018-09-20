package statistics

import "fmt"

//datePollutnStatInfoResponseItem 은 기간별 오염 통계 정보를 조회했을 때 서버에서
//return 해주는 데이터 집합이다.
type datePollutnStatInfoResponseItem struct {
	SidoName    string  `xml:"sidoName"` //지자체 이름
	DataTime    string  `xml:"dataTime"` //측정일
	So2Average  float64 `xml:"so2Avg"`   //아황산가스 평균농도 (단위 : ppm)
	CoAverage   float64 `xml:"coAvg"`    //일산화탄소 평균농도 (단위 : ppm)
	O3Average   float64 `xml:"o3Avg"`    //오존 평균농도 (단위 : ppm)
	No2Average  float64 `xml:"no2Avg"`   //이산화질소 평균농도 (단위 : ppm)
	Pm10Average int64   `xml:"pm10Avg"`  //미세먼지 평균농도 (단위 : ㎍/㎥)
	So2Max      float64 `xml:"so2Max"`   //아황산가스 최고농도 (단위 : ppm)
	CoMax       float64 `xml:"coMax"`    //일산화탄소 최고농도 (단위 : ppm)
	O3Max       float64 `xml:"o3Max"`    //오존 최고농도 (단위 : ppm)
	No2Max      float64 `xml:"no2Max"`   //이산화질소 최고농도 (단위 : ppm)
	Pm10Max     int64   `xml:"pm10Max"`  //미세먼지 최고농도 (단위 : ㎍/㎥)
	So2Min      float64 `xml:"so2Min"`   //아황산가스 최저농도 (단위 : ppm)
	CoMin       float64 `xml:"coMin"`    //일산화탄소 최저농도 (단위 : ppm)
	O3Min       float64 `xml:"o3Min"`    //오존 최저농도 (단위 : ppm)
	No2Min      float64 `xml:"no2Min"`   //이산화질소 최저농도 (단위 : ppm)
	Pm10Min     int64   `xml:"pm10Min"`  //미세먼지 최저농도 (단위 : ㎍/㎥)
}

//datePollutnStatInfoRequest 은 기간별 오염 통계 정보를 조회할 때 사용하는 데이터 집합이다.
type datePollutnStatInfoRequest struct {
	SearchDataTime       string //조회 날짜
	StatArticleCondition string //측정망정보(도시대기, 도로변대기, 국가배경농도, 교외대기)
	PageNo               int64  //페이지 번호
	NumOfRows            int64  //한 페이지 결과 수
}

//Op2ResponseItem 은 datePollutnStatInfoResponseItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op2ResponseItem datePollutnStatInfoResponseItem

//Op2Request 은 datePollutnStatInfoRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op2Request datePollutnStatInfoRequest

//Op2Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op2Operator struct {
	Req Op2Request
}

//Request는 datePollutnStatInfoRequest 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op2Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 2번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op2Operator) RequestString() string {
	return fmt.Sprintf("%s/%s?searchDataTime=%s&statArticleCondition=%s&pageNo=%d&numOfRows=%d",
		serviceName, operationMap[2],
		o.Req.SearchDataTime, o.Req.StatArticleCondition, o.Req.PageNo, o.Req.NumOfRows)
}

//Item은 Op2ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op2Operator) Item() interface{} {
	return &Op2ResponseItem{}
}
