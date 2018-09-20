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

//ctprvnMesureSidoLIstItem 은 시군구별 실시간 평균정보를 조회했을 경우 제공되는
//데이터를 저장하기 위한 구조체이다.
type ctprvnMesureSidoLIstItem struct {
	DataTime  string  `xml:"dataTime"`
	CityName  string  `xml:"cityName"`
	So2Value  float64 `xml:"so2Value"`
	CoValue   float64 `xml:"coValue"`
	O3Value   float64 `xml:"o3Value"`
	No2Value  float64 `xml:"no2Value"`
	Pm10Value int64   `xml:"pm10Value"`
	Pm25Value int64   `xml:"pm25Value"`
}

//ctprvnMesureSidoLIstRequest 은 시군구별 실시간 평균 정보를 조회하기 위한 시도명, 검색조건을
//서버에 전송하기 위한 구조체이다.
type ctprvnMesureSidoLIstRequest struct {
	SidoName        string //시도 이름(서울, 부산, 대구, 인천, 광주, 대전, 울산, 경기, 강원, 충북, 충남, 전북, 전남, 경북, 경남, 제주, 세종)
	SearchCondition string //요청 데이터기간(시간 : HOUR, 하루 : DAILY)
	PageNo          int64  // 페이지 번호
	NumOfRows       int64  // 한 페이지 결과 수
}

//Op6ResponseItem 은 ctprvnMesureSidoLIstItem 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op6ResponseItem ctprvnMesureSidoLIstItem

//UnmarshalXML은 입력되는 xml데이터를 사용자 정의에 맞도록 파싱하기 위하여 정의된 함수이다.
//ctprvnRltmMesureDnstyRequest를 통해서 데이터 요청시 ctprvnRltmMesureDnstyItem에 타입에 맞지 않는 특수 기호가 입력된다.
//'-' 값이 입력되면 xml의 디코더 인터페이스는 해당 파싱을 정지하고 에러를 리턴하므로 정상적인 파싱을 하기 위하여
//해당 값들을 string 타입으로 받아서 '-'값이 아닐경우에만 구조체의 멤버 타입에 맞도록 재변환하도록 한다.
func (o *Op6ResponseItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

//Op6Request 은 ctprvnMesureSidoLIstRequest 구조체를 다른 모듈에서 지원하도록 하기 위한 데이터 타입이다.
type Op6Request ctprvnMesureSidoLIstRequest

//Op6Operator 은 근접측정소 목록 OpenAPI를 실행하기 위한 명령을 만들도록 도와주는 인터페이스이다.
type Op6Operator struct {
	Req Op6Request
}

//Request는 ctprvnMesureSidoLIstRequest 객체 주소를 리턴한다.
//interface{}타입으로 리턴되므로 다른 모듈에서 데이터 사용시 타입지정이 필요하다.
func (o *Op6Operator) Request() interface{} {
	return &o.Req
}

//RequestString은 근접측정소 목록 조회의 Operation 6번에 대한 Url을 입력된
//데이터를 바탕으로 만들어서 리턴한다.
func (o Op6Operator) RequestString() string {
	return fmt.Sprintf("%s/%s?sidoName=%s&searchCondition=%s&pageNo=%d&numOfRows=%d",
		serviceName, operationMap[6],
		o.Req.SidoName, o.Req.SearchCondition, o.Req.PageNo, o.Req.NumOfRows)
}

//Item은 Op1ResponseItem 데이터를 생성하여 주소를 리턴한다.
func (o Op6Operator) Item() interface{} {
	return &Op6ResponseItem{}
}
