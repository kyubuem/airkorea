package station

//측정소 정보 조회 서비스이며, 근접측정소 목록조회, 측정소 목록 조회, TM 기준좌표 조회 서비스가 제공된다.
//평균 응답시간은 500ms이며 초당 최대 트랜젝션은 30tps이다.
//최대 메세지 사이즈는 1000KB이다. 또한 데이터 갱신 주기는 하루 1회(새벽 4시)이다.
var (
	serviceName  = "MsrstnInfoInqireSvc"
	operationMap = map[int64]string{
		1: "getNearbyMsrstnList", //근접측정소 목록 조회
		2: "getMsrstnList",       //측정소 목록 조회
		3: "getTMStdrCrdnt",      //TM 기준좌표 조회
	}
)

//Op1 은 근접측정소 목록 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op1(r Op1Request) *Op1Operator {
	return &Op1Operator{
		Req: r,
	}
}

//Op2 은 측정소 목록 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op2(r Op2Request) *Op2Operator {
	return &Op2Operator{
		Req: r,
	}
}

//Op3 은 TM 기준좌표 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op3(r Op3Request) *Op3Operator {
	return &Op3Operator{
		Req: r,
	}
}
