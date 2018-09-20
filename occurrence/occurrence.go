package occurrence

//오존황사 발생정보조회 서비스로, 기간별 오존, 황사 정보를 조회 할 수 있다.
//평균 응답시간은 500ms이며 초당 최대 트랜젝션은 30tps이다.
var (
	serviceName  = "OzYlwsndOccrrncInforInqireSvc"
	operationMap = map[int64]string{
		1: "getOzAdvsryOccrrncInfo",     //오존주의보 발생정보 조회
		2: "getYlwsndAdvsryOccrrncInfo", //황사주의보 발생정보 조회
	}
)

//Op1 은 오존주의보 발생 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op1(r Op1Request) *Op1Operator {
	return &Op1Operator{
		Req: r,
	}
}

//Op2 은 황사주의보 발생 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op2(r Op2Request) *Op2Operator {
	return &Op2Operator{
		Req: r,
	}
}
