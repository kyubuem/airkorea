package measure

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
)

//ctprvnRltmMesureDnstyItem 은 시도별 실시간 측정정보를 조회했을 경우 제공되는
//데이터를 저장하기 위한 구조체이다.
type ctprvnRltmMesureDnstyItem struct {
	StationName string  `xml:"stationName,omitempty"` //측정소 이름
	ManageName  string  `xml:"mangName,omitempty"`    //측정망 정보 (국가배경, 교외대기, 도시대기, 도로변대기)
	DataTime    string  `xml:"dataTime,omitempty"`    //오염도 측정 연-월-일 시간:분
	So2Value    float64 `xml:"so2Value,omitempty"`    //아황산가스 농도(단위 : ppm)
	CoValue     float64 `xml:"coValue,omitempty"`     //일산화탄소 농도(단위 : ppm)
	O3Value     float64 `xml:"o3Value,omitempty"`     //오존 농도(단위 : ppm)
	No2Value    float64 `xml:"no2Value,omitempty"`    //이산화질소 농도(단위 : ppm)
	Pm10Value   int64   `xml:"pm10Value,omitempty"`   //미세먼지(PM10) 농도 (단위 : ㎍/㎥)
	Pm10Value24 int64   `xml:"pm10Value24,omitempty"` //미세먼지(PM10)24시간예측이동농도(단위 : ㎍/㎥)
	Pm25Value   int64   `xml:"pm25Value,omitempty"`   //미세먼지(PM2.5) 농도(단위 : ㎍/㎥)
	Pm25Value24 int64   `xml:"pm25Value24,omitempty"` //미세먼지(PM2.5) 24시간예측이동농도(단위 : ㎍/㎥)
	KhaiValue   int64   `xml:"khaiValue,omitempty"`   //통합대기환경수치
	KhaiGrade   int64   `xml:"khaiGrade,omitempty"`   //통합대기환경지수
	So2Grade    int64   `xml:"so2Grade,omitempty"`    //아황산가스 지수
	CoGrade     int64   `xml:"coGrade,omitempty"`     //일산화탄소 지수
	O3Grade     int64   `xml:"o3Grade,omitempty"`     //오존 지수
	No2Grade    int64   `xml:"no2Grade,omitempty"`    //이산화질소 지수
	Pm10Grade   int64   `xml:"pm10Grade,omitempty"`   //미세먼지(PM10) 24시간 등급자료
	Pm25Grade   int64   `xml:"pm25Grade,omitempty"`   //미세먼지(PM2.5) 24시간 등급자료
	Pm10Grade1h int64   `xml:"pm10Grade1h,omitempty"` //미세먼지(PM10) 24시간 등급자료
	Pm25Grade1h int64   `xml:"pm25Grade1h,omitempty"` //미세먼지(PM2.5) 24시간 등급자료
}

//ctprvnRltmMesureDnstyRequest 은 시도별 실시간 측정 정보를 조회하기 위한 시도명을
//서버에 전송하기 위한 구조체이다.
type ctprvnRltmMesureDnstyRequest struct {
	SidoName  string //시도 이름(서울, 부산, 대구, 인천, 광주, 대전, 울산, 경기, 강원, 충북, 충남, 전북, 전남, 경북, 경남, 제주, 세종)
	PageNo    int64  //페이지 번호
	NumOfRows int64  //한 페이지 결과 수
	Version   string //버전 미포함 : PM25 데이터 미포함, 1.0 : PM25데이터 포함, 1.1 PM10,PM25 24시간 예측이동 평균 데이터 포함, 1.2 : 측정망 정보 데이터 포함, 1.3 : PM10 PM25 1시간 등급 자료 포함
}

//Op3ResponseItem 은 ctprvnRltmMesureDnstyItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op3ResponseItem ctprvnRltmMesureDnstyItem

//UnmarshalXML은 입력되는 xml데이터를 사용자 정의에 맞도록 파싱하기 위하여 정의된 함수이다.
//ctprvnRltmMesureDnstyRequest를 통해서 데이터 요청시 ctprvnRltmMesureDnstyItem에 타입에 맞지 않는 특수 기호가 입력된다.
//'-' 값이 입력되면 xml의 디코더 인터페이스는 해당 파싱을 정지하고 에러를 리턴하므로 정상적인 파싱을 하기 위하여
//해당 값들을 string 타입으로 받아서 '-'값이 아닐경우에만 구조체의 멤버 타입에 맞도록 재변환하도록 한다.
func (o *Op3ResponseItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	fieldMap := getFieldMap(o)
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch se := t.(type) {
		case xml.StartElement:
			var temp string
			if err := d.DecodeElement(&temp, &start); err != nil {
				return err
			} else {
				if strings.Compare(temp, "-") == 0 {
					break
				} else {
					x := reflect.ValueOf(o).Elem().FieldByName(fieldMap[se.Name.Local])
					switch x.Kind() {
					case reflect.Float64:
						v, _ := strconv.ParseFloat(temp, 64)
						x.Set(reflect.ValueOf(v))
					case reflect.Int64:
						v, _ := strconv.ParseInt(temp, 10, 64)
						x.Set(reflect.ValueOf(v))
					case reflect.String:
						x.Set(reflect.ValueOf(temp))
					default:
						return errors.New("Not Supported Kind type " + x.Kind().String())
					}
				}
			}
		case xml.EndElement:
			return nil
		}
	}
	return nil
}

//Op3Request 은 ctprvnRltmMesureDnstyRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op3Request ctprvnRltmMesureDnstyRequest

//Op3Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op3Operator struct {
	Req Op3Request
}

//Request는 ctprvnRltmMesureDnstyRequest 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op3Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 3번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op3Operator) RequestString() string {
	url := fmt.Sprintf("%s/%s?sidoName=%s&pageNo=%d&numOfRows=%d",
		serviceName, operationMap[3], o.Req.SidoName, o.Req.PageNo, o.Req.NumOfRows)
	if strings.Compare(o.Req.Version, "1.0") == 0 ||
		strings.Compare(o.Req.Version, "1.1") == 0 ||
		strings.Compare(o.Req.Version, "1.2") == 0 ||
		strings.Compare(o.Req.Version, "1.3") == 0 {
		url = fmt.Sprintf("%s&ver=%s", url, o.Req.Version)
	}
	return url
}

//Item은 Op3ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op3Operator) Item() interface{} {
	return &Op3ResponseItem{}
}
