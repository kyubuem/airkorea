## airkorea
airkorea 에서 제공되는 API문서를 golang버전으로 만든 라이브러리이다.
XML 파싱을 통하여 response 구조체로 받아온다.
해당 라이브러리를 사용하기 위해서는 airkorea에서 서비스를 신청해야 한다.
서비스 신청이 완료 되면 servicekey를 할당 받게 되는데 private 키이므로 별도의 키 파일을 만들어서 사용한다.
```shell
$ cd $(go to main.go dir)
$ touch key
$ vi key
$ put your private key
```
키 파일 생성이 완료 되면 라이브러리를 사용하고자 하는 패키지에서 키를 읽어서 airkorea 패키지에 설정해주면 된다.
```go
	if key, err := ioutil.ReadFile("key"); err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Success loading service key")
		airkorea.SetKey(string(key))
	}
```

### airkora Service

1. 측정소 정보 조회 서비스  
대기질 측정소 정보를 조회하기 위한 서비스로 TM 좌표 기반의 가까운 측정소 및 측정소 목록과  
측정소의 정보를 조회할 수 있다. 측정소정보 조회 서비스는 크게 3개의 명령이 제공된다.  
  
    * 근접측정소 목록 조회  
    TM 좌표를 이용하여 좌표 주변 측정소 정보와 측정장소와 좌표간의 거리 정보를 제공하는 서비스  
    ```go
        if stationOp1Res, err := airkorea.Get(station.Op1(station.Op1Request{
            TmX:       "945959.0381341814",
            TmY:       "1953851.7348996028",
            Version:   airkorea.Ver10,
            PageNo:    1,
            NumOfRows: 10,
        })); err != nil {
            log.Fatalln(err)
        } else {
            for _, v := range stationOp1Res.Body.Items.Item {
                log.Printf("%+v\n", v)
            }
        }
    ```
      
    * 측정소 목록 조회  
    측정소 주소 또는 측정소 명칭으로 측정소 목록 또는 단 건의 측정소 상세 정보를 제공하는 서비스  
    ```go
        if stationOp2Res, err := airkorea.Get(station.Op2(station.Op2Request{
            Address:     "서울",
            StationName: "종로구",
            PageNo:      1,
            NumOfRows:   10,
        })); err != nil {
            log.Fatalln(err)
        } else {
            for _, v := range stationOp2Res.Body.Items.Item {
                log.Printf("%+v\n", v)
            }
        }
    ```
      
    * TM 기준좌표 조회  
    검색서비스를 사용하여 읍면동 이름을 검색조건으로 기준좌표(TM좌표) 정보를 제공하는 서비스  
    ```go
        if stationOp3Res, err := airkorea.Get(station.Op3(station.Op3Request{
            UmdName:   "혜화동",
            PageNo:    1,
            NumOfRows: 10,
        })); err != nil {
            log.Fatalln(err)
        } else {
            for _, v := range stationOp3Res.Body.Items.Item {
                log.Printf("%+v\n", v)
            }
        }
    ```
      
2. 대기오염 정보 조회 서비스  
각 측정소별 대기오염정보를 조회하기 위한 서비스로 기간별, 시도별 대기오염 정보와 통합대기환경지수 나쁨 이상 측정소 내역,  
대기질(미세먼지/오존) 예보 통보 내역 등을 조회할 수 있다. 대기오염 정보 조회 서비스는 크게 6개의 명령이 제공된다.  
  
    * 측정소별 실시간 측정정보 조회  
    측정소명과 측정데이터 기간(일, 한달, 3개월)으로 해당 측정소의 일반 항목 측정정보를 제공하는 측정소별 실시간 측정정보 조회  
    ```go
    	if measureOp1Res, err := airkorea.Get(measure.Op1(measure.Op1Request{
    		StationName: "종로구",
    		DataTerm:    airkorea.Daily,
    		PageNo:      1,
    		NumOfRows:   10,
    		Version:     airkorea.Ver13,
    	})); err != nil {
    		log.Fatalln(err)
    	} else {
    		for _, v := range measureOp1Res.Body.Items.Item {
    			log.Printf("%+v\n", v)
    		}
    	}
    ```
      
    * 통합대기환경지수 나쁨 이상 측정소 목록 조회  
    통합대기환경지수가 나쁨 등급 이상인 측정소명과 주소 목록 정보를 제공하는 통합대기환경 지수 나쁨 이상 측정소 목록 조회  
    ```go
    	if measureOp2Res, err := airkorea.Get(measure.Op2(measure.Op2Request{
    		PageNo:    1,
    		NumOfRows: 10,
    	})); err != nil {
    		log.Fatalln(err)
    	} else {
    		for _, v := range measureOp2Res.Body.Items.Item {
    			log.Printf("%+v\n", v)
    		}
    	}
    ```
      
    * 시도별 실시간 측정정보 조회  
    시도명을 검색조건으로 하여 시도별 측정소 목록에 대한 일반 항목과 CAI 최종 실시간 측정값과 지수 정보 조회 기능을 제공하는 시도별 실시간 측정정보 조회
    ```go
    	if measureOp3Res, err := airkorea.Get(measure.Op3(measure.Op3Request{
    		SidoName:  airkorea.Seoul,
    		PageNo:    1,
    		NumOfRows: 10,
 		    Version: airkorea.Ver13,
    	})); err != nil {
    		log.Fatalln(err)
    	} else {
    		for _, v := range measureOp3Res.Body.Items.Item {
    			log.Printf("%+v\n", v)
    		}
    	}
    ```
      
    * 미세먼지/오존 예보통보 조회  
    통보코드와 통보시간으로 예보정보와 발생 원인 정보를 조회하는 대기질(미세먼지/오존) 예보통보 조회  
    ```go
        then := time.Now()
        if measureOp4Res, err := airkorea.Get(measure.Op4(measure.Op4Request{
            SearchDate: strings.Split(then.Format(time.RFC3339), "T")[0],
            InformCode: airkorea.Pm10,
            PageNo:     1,
            NumOfRows:  10,
            Version:    airkorea.Ver11,
        })); err != nil {
            log.Fatalln(err)
        } else {
            for _, v := range measureOp4Res.Body.Items.Item {
                log.Printf("%+v\n", v)
            }
        }
    ```
      
    * 시도별 실시간 평균정보 조회  
    시도별 측정소목록에 대한 일반 항목의 시간 및 일평균 자료 및 지역평균 정보를 제공하는 시도별 실시간 평균정보 조회  
    데이터 구분을 Hour로 요청시 검색 조건이 필요없다. Daily 입력시 Week, Month 을 입력해야 함
    ```go
    	if measureOp5Res, err := airkorea.Get(measure.Op5(measure.Op5Request{
    		ItemCode:        airkorea.Pm10,
    		DataGubun:       airkorea.Daily,
    		SearchCondition: airkorea.Week,
    		PageNo:          1,
    		NumOfRows:       10,
    	})); err != nil {
    		log.Fatalln(err)
    	} else {
    		for _, v := range measureOp5Res.Body.Items.Item {
    			log.Printf("%+v\n", v)
    		}
    	}
    ```
      
    * 시군구별 실시간 평균정보 조회  
    시도의 각 시군구별 측정소 목록의 일반 항목에 대한 시간대별 평균 농도를 제공하는 시군구별 실시간 평균정보 조회   
    ```go
    	if measureOp6Res, err := airkorea.Get(measure.Op6(measure.Op6Request{
    		SidoName:        airkorea.Gyeonggi,
    		SearchCondition: airkorea.Hour,
    		PageNo:          1,
    		NumOfRows:       30,
    	})); err != nil {
    		log.Fatalln(err)
    	} else {
    		for _, v := range measureOp6Res.Body.Items.Item {
    			log.Printf("%+v\n", v)
    		}
    	}
    ```
      
3. 대기오염 통계 서비스  
대기오염 통계 정보를 조회하기 위한 서비스로 각 측정소별 확정 농도 정보와 기간별 통계수치 정보를 조회할 수 있다.  
  
    * 측정소별 최종확정 농도 조회  
    측정소명을 입력하여 일별, 월별, 연별 확정농도 수치를 조회
    ```go
    	if statisticsOp1Res, err := airkorea.Get(statistics.Op1(statistics.Op1Request{
    		StationName:     "종로구",
    		SearchCondition: airkorea.Daily,
    		PageNo:          1,
    		NumOfRows:       10,
    	})); err != nil {
    		log.Fatalln(err)
    	} else {
    		for _, v := range statisticsOp1Res.Body.Items.Item {
    			log.Printf("%+v\n", v)
    		}
    	}
    ```
      
    * 기간별 오염통계 정보 조회  
    측정항목, 연도정보를 입력받아 지자체별 평균농도값을 조회  
    ```go
        if statisticsOp2Res, err := airkorea.Get(statistics.Op2(statistics.Op2Request{
            SearchDataTime:       "2016-12",
            StatArticleCondition: airkorea.CityAir,
            PageNo:               1,
            NumOfRows:            10,
        })); err != nil {
            log.Fatalln(err)
        } else {
            for _, v := range statisticsOp2Res.Body.Items.Item {
                log.Printf("%+v\n", v)
            }
        }
    ```
      
4. 오존/황사 발생정보 조회 서비스
오존, 황사 발생 정보를 조회하기 위한 서비스
  
    * 오존주의보 발생 정보  
    조회연도를 입력하여 오존주의보 발생정보를 조회  
    ```go
       	if occurrenceOp1Res, err := airkorea.Get(occurrence.Op1(occurrence.Op1Request{
       		Year:      2018,
       		PageNo:    1,
       		NumOfRows: 10,
       	})); err != nil {
       		log.Fatalln(err)
       	} else {
       		for _, v := range occurrenceOp1Res.Body.Items.Item {
       			log.Printf("%+v\n", v)
       		}
       	}
    ```
      
    * 황사주의보 발생 정보  
    조회연도에 황사발생정보를 조회
    ```go
    	if occurrenceOp2Res, err := airkorea.Get(occurrence.Op2(occurrence.Op2Request{
    		Year:      2018,
    		PageNo:    1,
    		NumOfRows: 10,
    	})); err != nil {
    		log.Fatalln(err)
    	} else {
    		for _, v := range occurrenceOp2Res.Body.Items.Item {
    			log.Printf("%+v\n", v)
    		}
    	}
    ```
    
5. 에러코드  
  
|No|Error Name|Decription|
| :---: | :--- | :--- |  
|1|ApplicationError|제공기관 서비스 제공 상태가 원할하지 않습니다.|
|2|DBError|제공기관 서비스 제공 상태가 원할하지 않습니다.|
|3|NoData|데이터없음 에러|
|4|HttpError|제공기관 서비스 제공 상태가 원할하지 않습니다|
|5|ServiceTimeOut|제공기관 서비스 제공 상태가 원할하지 않습니다|
|10|WrongRequestedParameter|OpenApi요청시 ServiceKey파라미터가 없음|
|11|NoMustRequestParameter|요청하신 OpenApi의 필수 파라미터가 누락되었습니다.|
|12|NoOpenApi|OpenApi 호출시 URL이 잘못됨|
|20|AccessDenied|활용승인이 되지 않은 OpenApi 호출|
|22|ServiceRequestCountLimit|일일 활용건수가 초과함(활용건수 증가 필요)| 
|30|NotRegisteredKey|잘못된 서비스키를 사용하였거나 서비스키를 URL인코딩하지 않음|
|31|ExpiredServiceKey|OpenApi 사용기간이 만료됨(활용언장신청 후 사용가능)|
|32|NotRegisteredDomainOrIp|활용신청한 서버의 IP와 실제 OpenAPI호출한 서버가 다를 경우|
