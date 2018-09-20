package measure

import (
	"reflect"
	"strings"
	"unicode"
)

//각 측정소별 대기오염정보를 조회하기 위한 서비스로 기간별, 시도별
//대기오염 정보와 통합대기환경지수 나쁨 이상 측정소 내역, 대기질(미세먼지/오존) 예보 통보
//내역 등을 조회할 수 있다.
//평균 응답시간은 500ms이며 초당 최대 트랜젝션은 30tps이다.
//데이터 갱신 주기는 실시간 정보의 경우 10분(매 시간 시간자료 갱신은 20분 전후로 반영됨)이고,
//대기질 예보 정보는 매 시간 22분, 57분에 갱신된다.
var (
	serviceName  = "ArpltnInforInqireSvc"
	operationMap = map[int64]string{
		1: "getMsrstnAcctoRltmMesureDnsty",              //측정소별 실시간 측정정보 조회
		2: "getUnityAirEnvrnIdexSnstiveAboveMsrstnList", //통합대기환경지수 나쁨 이상 측정소 목록 조회
		3: "getCtprvnRltmMesureDnsty",                   //시도별 실시간 측정정보 조회
		4: "getMinuDustFrcstDspth",                      //미세먼지/오존 예보통보 조회
		5: "getCtprvnMesureLIst",                        //시도별 실시간 평균정보 조회
		6: "getCtprvnMesureSidoLIst",                    //시구군별 실시간 평균정보 조회
	}
)

//Op1 은 측정소별 실시간 측정정보 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op1(r Op1Request) *Op1Operator {
	return &Op1Operator{
		Req: r,
	}
}

//Op2 은 통합대기환경지수 나쁨 이상 측정소 목록 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op2(r Op2Request) *Op2Operator {
	return &Op2Operator{
		Req: r,
	}
}

//Op3 은 시도별 실시간 측정정보 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op3(r Op3Request) *Op3Operator {
	return &Op3Operator{
		Req: r,
	}
}

//Op4 은 미세먼지/오존 예보통보 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op4(r Op4Request) *Op4Operator {
	return &Op4Operator{
		Req: r,
	}
}

//Op5 은 시도별 실시간 평균정보 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op5(r Op5Request) *Op5Operator {
	return &Op5Operator{
		Req: r,
	}
}

//Op6 은 시군구별 실시간 평균정보 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op6(r Op6Request) *Op6Operator {
	return &Op6Operator{
		Req: r,
	}
}

//getFieldMap은 제공된 xml태그를 기준으로 Operation별 Item 구조체의 필드를 map[string]string
//형태로 리턴한다.
//해당 구조체는 태그의 첫번째 값을 기준으로 key를 생성하도록 설계되어있다.
//따라서, Item 구조체 설계시 태그값의 2번째 인자를 키값으로 설정해야한다.
//예제
//type Item struct {
//    Name string `xml:"name"`
//}
//FieldMap는 name이라는 key를 갖게 되며 Name의 Address를 value로 저장하게 된다.
func getFieldMap(i interface{}) (fieldMap map[string]string) {
	fieldMap = make(map[string]string)
	rt := reflect.Indirect(reflect.ValueOf(i)).Type()
	for i := 0; i < rt.NumField(); i++ {
		tags := strings.FieldsFunc(string(rt.Field(i).Tag), func(c rune) bool {
			return !unicode.IsLetter(c) && !unicode.IsNumber(c)
		})
		fieldMap[tags[1]] = rt.Field(i).Name
	}
	return
}
