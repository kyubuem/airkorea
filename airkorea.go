package airkorea

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

const (
	_                        = iota
	ApplicationError         = 1  //제공기관 서비스 제공 상태가 원할하지 않습니다.
	DBError                  = 2  //제공기관 서비스 제공 상태가 원할하지 않습니다.
	NoData                   = 3  //데이터없음 에러
	HttpError                = 4  //제공기관 서비스 제공 상태가 원할하지 않습니다.
	ServiceTimeOut           = 5  //제공기관 서비스 제공 상태가 원할하지 않습니다.
	WrongRequestedParameter  = 10 //OpenApi요청시 ServiceKey파라미터가 없음
	NoMustRequestParameter   = 11 //요청하신 OpenApi의 필수 파라미터가 누락되었습니다.
	NoOpenApi                = 12 //OpenApi 호출시 URL이 잘못됨
	AccessDenied             = 20 //활용승인이 되지 않은 OpenApi 호출
	ServiceRequestCountLimit = 22 //일일 활용건수가 초과함(활용건수 증가 필요)
	NotRegisteredKey         = 30 //잘못된 서비스키를 사용하였거나 서비스키를 URL인코딩하지 않음
	ExpiredServiceKey        = 31 //OpenApi 사용기간이 만료됨(활용언장신청 후 사용가능)
	NotRegisteredDomainOrIp  = 32 //활용신청한 서버의 IP와 실제 OpenAPI호출한 서버가 다를 경우
)

var (
	endPoint   = "http://openapi.airkorea.or.kr/openapi/services/rest"
	serviceKey = ""
)

var (
	Hour       = "HOUR"
	Daily      = "DAILY"
	Week       = "WEEK"
	Month      = "MONTH"
	ThreeMonth = "3MONTH"
)

var (
	Pm10 = "PM10"
	Pm25 = "PM25"
	O3   = "O3"
)

var (
	Ver10 = "1.0"
	Ver11 = "1.1"
	Ver12 = "1.2"
	Ver13 = "1.3"
)

var (
	Seoul     = "서울"
	Busan     = "부산"
	Daegu     = "대구"
	Incheon   = "인천"
	Gwangju   = "광주"
	Daejeon   = "대전"
	Ulsan     = "울산"
	Gyeonggi  = "경기"
	Gangwon   = "강원"
	Chungbuk  = "충북"
	Chungnam  = "충남"
	Jeonbuk   = "전북"
	Jeonnam   = "전남"
	Gyeongbuk = "경북"
	Gyeongnam = "경남"
	Jeju      = "제주"
	Sejong    = "세종"
)

var (
	CityAir    = "도시대기"
	RoadAir    = "도로변대기"
	CountryAir = "국가배경농도"
	OutsideAir = "교외대기"
)

//SetKey 는 airkorea에서 개인별로 지급되는 서비스 키로
//키 파일을 별도로 만들어서 읽어온 뒤, 설정하도록 한다.
//	if key.conf, err := ioutil.ReadFile("key.conf"); err != nil {
//		log.Fatalln(err)
//	} else {
//		log.Println("Success loading service key.conf")
//		airkorea.SetKey(string(key.conf))
//	}
func SetKey(s string) {
	serviceKey = s
}

//Response 는 airkorea에서 데이터 요청시 리턴되는 request로 모든 오퍼레이션이 동일한 타입으로 리턴된다.
type Response struct {
	XMLName xml.Name `xml:"response"`
	Header  Header   `xml:"header"`
	Body    Body     `xml:"body"`
}

func (r *Response) GetItem() []interface{} {
	return r.Body.Items.Item
}

//Header 은 요청한 데이터의 결과를 나타낸다
type Header struct {
	XMLName       xml.Name `xml:"header"`
	ResultCode    int64    `xml:"resultCode"`
	ResultMessage string   `xml:"resultMsg"`
}

//Body 는 요청한 데이터정보를 갖는다.
type Body struct {
	XMLName    xml.Name `xml:"body"`
	Items      ItemList `xml:"items"`
	NumOfRows  int64    `xml:"numOfRows"`
	PageNo     int64    `xml:"pageNo"`
	TotalCount int64    `xml:"totalCount"`
}

//ItemList 는 오퍼레이션 별로 다른 데이터를 받아오기 때문에 interface타입으로 사용한다.
//오퍼레이션별 Item구조체를 동적으로 바인딩하여 결과를 취합하여 저장한다.
type ItemList struct {
	XMLName  xml.Name      `xml:"items"`
	Item     []interface{} `xml:"item"`
	ItemType reflect.Type
}

//UnmarshalXML 은 요청한 데이터의 xml을 사용자 정의에 맞도록 파싱하기 위하여 정의된 함수이다.
//item의 경우 empty interface 타입이므로 Decoder가 어떤 타입으로 파싱해야 하는지 모르기 때문에
//각각의 오퍼레이션에 맞는 구조체로 파싱하도록 하는 역할을 한다.
//오퍼레이션 별 아이템들은 키값 "item"으로 입력되고, 오퍼레이션의 ResponseItem 구조체의 타입으로 동적할당 하여
//파싱한 데이터를 입력한다.
func (i *ItemList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch s := t.(type) {
		case xml.StartElement:
			if s.Name.Local == "item" {
				it := reflect.New(i.ItemType)
				if err := d.DecodeElement(it.Interface(), &start); err != nil {
					return err
				}
				i.Item = append(i.Item, it.Interface())
			}
		case xml.EndElement:
			return nil
		}
	}
	return nil
}

//ServiceOperator 은 서비스별 오퍼레이션을 동작시키기 위한 인터페이스이다.
//Request() 함수는 각 오퍼레이션의 Request 구조체를 리턴한다.
//RequestString() 함수는 오퍼레이션의 API Url을 Request 구조체를 이용하여 생성후 리턴한다.
//Item()은 ResponseItem 구조체를 리턴한다.
type ServiceOperator interface {
	Request() interface{}
	RequestString() string
	Item() interface{}
}

//Get 은 각 오퍼레이션별 request를 받아서 response를 리턴하도록한다
//response된 []byte타입의 데이터는 Unmarshal을 통해서 사용자에게 필요한 구조체형태로 변환된다.
func Get(s ServiceOperator) (*Response, error) {
	resp, err := http.Get(request(s))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := response(s)
	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//request 는 유저가 요청한 데이터를 바탕으로 API url을 생성한다.
func request(s ServiceOperator) string {
	str := fmt.Sprintf("%s/%s&ServiceKey=%s",
		endPoint, s.RequestString(), serviceKey)
	//log.Println(str)
	return str
}

//response는 요청한 오퍼레이션에 맞는 Item구조체를 만들어 Unmarshal시 필요한 구조체로 만드는 역할을 한다.
func response(s ServiceOperator) *Response {
	return &Response{
		Header: Header{},
		Body: Body{
			Items: ItemList{
				ItemType: reflect.TypeOf(s.Item()).Elem(),
			},
		},
	}
}
