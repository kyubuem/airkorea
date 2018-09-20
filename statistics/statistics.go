package statistics

//대기오염통계 서비스로, 측정소별 최종확정 농도, 기간별 오염통계정보를 조회 할 수 있다.
//평균 응답시간은 500ms이며 초당 최대 트랜젝션은 30tps이다.
var (
	serviceName  = "ArpltnStatsSvc"
	operationMap = map[int64]string{
		1: "getMsrstnAcctoLastDcsnDnsty", //측정소별 최종확정 농도 조회
		2: "getDatePollutnStatInfo",      //기간별 오염통계 정보 조회
	}
)

//Op1 은 측정소별 최종확정 농도 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op1(r Op1Request) *Op1Operator {
	return &Op1Operator{
		Req: r,
	}
}

//Op2 은 기간별 오염통계 정보 조회 서비스의 오퍼레이터를 생성하여 리턴한다.
func Op2(r Op2Request) *Op2Operator {
	return &Op2Operator{
		Req: r,
	}
}
