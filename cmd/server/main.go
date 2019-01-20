package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/kyubuem/airkorea"
	"github.com/kyubuem/airkorea/measure"
	"github.com/kyubuem/airkorea/occurrence"
	pb "github.com/kyubuem/airkorea/proto"
	"github.com/kyubuem/airkorea/station"
	"github.com/kyubuem/airkorea/statistics"
	opentracing "github.com/opentracing/opentracing-go"
	zipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) MeasureOp1Service(ctx context.Context, in *pb.MeasureOp1Request) (*pb.MeasureOp1Response, error) {
	res, err := airkorea.Get(measure.Op1(measure.Op1Request{
		StationName: in.GetStationName(),
		DataTerm:    in.GetDataTerm(),
		PageNo:      in.GetPageNo(),
		NumOfRows:   in.GetNumOfRows(),
		Version:     in.GetVersion(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("MeasureOp1Service failed due to some unknown reason")
	}

	items := make([]*pb.MeasureOp1Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*measure.Op1ResponseItem)
		items = append(items, &pb.MeasureOp1Response_Item{
			DataTime:      item.DataTime,
			ManageName:    item.ManageName,
			So2Value:      item.So2Value,
			CoValue:       item.CoValue,
			O3Value:       item.O3Value,
			No2Value:      item.No2Value,
			Pm10Value:     item.Pm10Value,
			Pm10Value_24H: item.Pm10Value24,
			Pm25Value:     item.Pm25Value,
			Pm25Value_24H: item.Pm25Value24,
			KhaiValue:     item.KhaiValue,
			KhaiGrade:     item.KhaiGrade,
			So2Grade:      item.So2Grade,
			CoGrade:       item.CoGrade,
			O3Grade:       item.O3Grade,
			No2Grade:      item.No2Grade,
			Pm10Grade:     item.Pm10Grade,
			Pm25Grade:     item.Pm25Grade,
			Pm10Grade_1H:  item.Pm10Grade1h,
			Pm25Grade_1H:  item.Pm25Grade1h,
		})
	}

	return &pb.MeasureOp1Response{
		Items: items,
	}, nil

}

func (s *server) MeasureOp2Service(ctx context.Context, in *pb.MeasureOp2Request) (*pb.MeasureOp2Response, error) {
	res, err := airkorea.Get(measure.Op2(measure.Op2Request{
		PageNo:    in.GetPageNo(),
		NumOfRows: in.GetNumOfRows(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("MeasureOp2Service failed due to some unknown reason")
	}

	items := make([]*pb.MeasureOp2Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*measure.Op2ResponseItem)
		items = append(items, &pb.MeasureOp2Response_Item{
			StationName: item.StationName,
			Address:     item.Address,
		})
	}

	return &pb.MeasureOp2Response{
		Items: items,
	}, nil
}

func (s *server) MeasureOp3Service(ctx context.Context, in *pb.MeasureOp3Request) (*pb.MeasureOp3Response, error) {
	res, err := airkorea.Get(measure.Op3(measure.Op3Request{
		SidoName:  in.GetSidoName(),
		PageNo:    in.GetPageNo(),
		NumOfRows: in.GetNumOfRows(),
		Version:   in.GetVersion(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("MeasureOp3Service failed due to some unknown reason")
	}

	items := make([]*pb.MeasureOp3Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*measure.Op3ResponseItem)
		items = append(items, &pb.MeasureOp3Response_Item{
			StationName:   item.StationName,
			ManageName:    item.ManageName,
			DataTime:      item.DataTime,
			So2Value:      item.So2Value,
			CoValue:       item.CoValue,
			O3Value:       item.O3Value,
			No2Value:      item.No2Value,
			Pm10Value:     item.Pm10Value,
			Pm10Value_24H: item.Pm10Value24,
			Pm25Value:     item.Pm25Value,
			Pm25Value_24H: item.Pm25Value24,
			KhaiValue:     item.KhaiValue,
			KhaiGrade:     item.KhaiGrade,
			So2Grade:      item.So2Grade,
			CoGrade:       item.CoGrade,
			O3Grade:       item.O3Grade,
			No2Grade:      item.No2Grade,
			Pm10Grade:     item.Pm10Grade,
			Pm25Grade:     item.Pm25Grade,
			Pm10Grade_1H:  item.Pm10Grade1h,
			Pm25Grade_1H:  item.Pm25Grade1h,
		})
	}

	return &pb.MeasureOp3Response{
		Items: items,
	}, nil
}

func (s *server) MeasureOp4Service(ctx context.Context, in *pb.MeasureOp4Request) (*pb.MeasureOp4Response, error) {
	res, err := airkorea.Get(measure.Op4(measure.Op4Request{
		SearchDate: in.GetSearchDate(),
		InformCode: in.GetInformCode(),
		PageNo:     in.GetPageNo(),
		NumOfRows:  in.GetNumOfRows(),
		Version:    in.GetVersion(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("MeasureOp4Service failed due to some unknown reason")
	}

	items := make([]*pb.MeasureOp4Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*measure.Op4ResponseItem)
		items = append(items, &pb.MeasureOp4Response_Item{
			DataTime:      item.DataTime,
			InformCode:    item.InformCode,
			InformOverall: item.InformOverall,
			InformCause:   item.InformCause,
			InformGrade:   item.InformGrade,
			ActionKnack:   item.ActionKnack,
			ImageUrl_1:    item.ImageUrl1,
			ImageUrl_2:    item.ImageUrl2,
			ImageUrl_3:    item.ImageUrl3,
			ImageUrl_4:    item.ImageUrl4,
			ImageUrl_5:    item.ImageUrl5,
			ImageUrl_6:    item.ImageUrl6,
			ImageUrl_7:    item.ImageUrl7,
			ImageUrl_8:    item.ImageUrl8,
			ImageUrl_9:    item.ImageUrl9,
			InformData:    item.InformData,
		})
	}

	return &pb.MeasureOp4Response{
		Items: items,
	}, nil
}

func (s *server) MeasureOp5Service(ctx context.Context, in *pb.MeasureOp5Request) (*pb.MeasureOp5Response, error) {
	res, err := airkorea.Get(measure.Op5(measure.Op5Request{
		DataGubun:       in.GetDataGubun(),
		SearchCondition: in.GetSearchCondition(),
		ItemCode:        in.GetItemCode(),
		PageNo:          in.GetPageNo(),
		NumOfRows:       in.GetNumOfRows(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("MeasureOp5Service failed due to some unknown reason")
	}

	items := make([]*pb.MeasureOp5Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*measure.Op5ResponseItem)
		items = append(items, &pb.MeasureOp5Response_Item{
			DataTime:  item.DataTime,
			ItemCode:  item.ItemCode,
			DataGubun: item.DataGubun,
			Seoul:     item.Seoul,
			Busan:     item.Busan,
			Daegu:     item.Daegu,
			Incheon:   item.Incheon,
			Gwangju:   item.Gwangju,
			Daejeon:   item.Daejeon,
			Ulsan:     item.Ulsan,
			Gyeonggi:  item.Gyeonggi,
			Gangwon:   item.Gangwon,
			Chungbuk:  item.Chungbuk,
			Chungnam:  item.Chungnam,
			Jeonbuk:   item.Jeonbuk,
			Jeonnam:   item.Jeonnam,
			Gyeongbuk: item.Gyeongbuk,
			Gyeongnam: item.Gyeongnam,
			Jeju:      item.Jeju,
			Sejong:    item.Sejong,
		})
	}

	return &pb.MeasureOp5Response{
		Items: items,
	}, nil
}

func (s *server) MeasureOp6Service(ctx context.Context, in *pb.MeasureOp6Request) (*pb.MeasureOp6Response, error) {
	res, err := airkorea.Get(measure.Op6(measure.Op6Request{
		SidoName:        in.GetSidoName(),
		SearchCondition: in.GetSearchCondition(),
		PageNo:          in.GetPageNo(),
		NumOfRows:       in.GetNumOfRows(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("MeasureOp6Service failed due to some unknown reason")
	}

	items := make([]*pb.MeasureOp6Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*measure.Op6ResponseItem)
		items = append(items, &pb.MeasureOp6Response_Item{
			DataTime:  item.DataTime,
			CityName:  item.CityName,
			So2Value:  item.So2Value,
			CoValue:   item.CoValue,
			O3Value:   item.O3Value,
			No2Value:  item.No2Value,
			Pm10Value: item.Pm10Value,
			Pm25Value: item.Pm25Value,
		})
	}

	return &pb.MeasureOp6Response{
		Items: items,
	}, nil
}

func (s *server) OccurrenceOp1Service(ctx context.Context, in *pb.OccurrenceOp1Request) (*pb.OccurrenceOp1Response, error) {
	res, err := airkorea.Get(occurrence.Op1(occurrence.Op1Request{
		Year:      in.GetYear(),
		PageNo:    in.GetPageNo(),
		NumOfRows: in.GetNumOfRows(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("OccurrenceOp1Service failed due to some unknown reason")
	}

	items := make([]*pb.OccurrenceOp1Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*occurrence.Op1ResponseItem)
		items = append(items, &pb.OccurrenceOp1Response_Item{
			DataTime:     item.DataTime,
			DistrictName: item.DistrictName,
			MoveName:     item.MoveName,
			IssueTime:    item.IssueTime,
			IssueVal:     item.IssueVal,
			ClearTime:    item.ClearTime,
			ClearVal:     item.ClearVal,
			MaxVal:       item.MaxVal,
		})
	}

	return &pb.OccurrenceOp1Response{
		Items: items,
	}, nil
}

func (s *server) OccurrenceOp2Service(ctx context.Context, in *pb.OccurrenceOp2Request) (*pb.OccurrenceOp2Response, error) {
	res, err := airkorea.Get(occurrence.Op2(occurrence.Op2Request{
		Year:      in.GetYear(),
		PageNo:    in.GetPageNo(),
		NumOfRows: in.GetNumOfRows(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("OccurrenceOp2Service failed due to some unknown reason")
	}

	items := make([]*pb.OccurrenceOp2Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*occurrence.Op2ResponseItem)
		items = append(items, &pb.OccurrenceOp2Response_Item{
			DataTime: item.DataTime,
			TmCnt:    item.TmCnt,
			TmArea:   item.TmArea,
		})
	}

	return &pb.OccurrenceOp2Response{
		Items: items,
	}, nil
}

func (s *server) StationOp1Service(ctx context.Context, in *pb.StationOp1Request) (*pb.StationOp1Response, error) {
	res, err := airkorea.Get(station.Op1(station.Op1Request{
		TmX:       in.GetTmX(),
		TmY:       in.GetTmY(),
		PageNo:    in.GetPageNo(),
		NumOfRows: in.GetNumOfRows(),
		Version:   in.GetVersion(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("StationOp1Service failed due to some unknown reason")
	}

	items := make([]*pb.StationOp1Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*station.Op1ResponseItem)
		items = append(items, &pb.StationOp1Response_Item{
			StationName: item.StationName,
			Address:     item.Address,
			Tm:          item.Tm,
		})
	}

	return &pb.StationOp1Response{
		Items: items,
	}, nil
}

func (s *server) StationOp2Service(ctx context.Context, in *pb.StationOp2Request) (*pb.StationOp2Response, error) {
	res, err := airkorea.Get(station.Op2(station.Op2Request{
		StationName: in.GetStationName(),
		Address:     in.GetAddress(),
		PageNo:      in.GetPageNo(),
		NumOfRows:   in.GetNumOfRows(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("StationOp2Service failed due to some unknown reason")
	}

	items := make([]*pb.StationOp2Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*station.Op2ResponseItem)
		items = append(items, &pb.StationOp2Response_Item{
			StationName: item.StationName,
			Address:     item.Address,
			Year:        item.Year,
			Oper:        item.Oper,
			Photo:       item.Photo,
			Vrml:        item.Vrml,
			Map:         item.Map,
			ManageName:  item.ManageName,
			Item:        item.Item,
			DmX:         item.DmX,
			DmY:         item.DmY,
		})
	}

	return &pb.StationOp2Response{
		Items: items,
	}, nil
}

func (s *server) StationOp3Service(ctx context.Context, in *pb.StationOp3Request) (*pb.StationOp3Response, error) {
	res, err := airkorea.Get(station.Op3(station.Op3Request{
		UmdName:   in.GetUmdName(),
		PageNo:    in.GetPageNo(),
		NumOfRows: in.GetNumOfRows(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("StationOp3Service failed due to some unknown reason")
	}

	items := make([]*pb.StationOp3Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*station.Op3ResponseItem)
		items = append(items, &pb.StationOp3Response_Item{
			SidoName: item.SidoName,
			SggName:  item.SggName,
			UmdName:  item.UmdName,
			TmX:      item.TmX,
			TmY:      item.TmY,
		})
	}

	return &pb.StationOp3Response{
		Items: items,
	}, nil
}

func (s *server) StatisticsOp1Service(ctx context.Context, in *pb.StatisticsOp1Request) (*pb.StatisticsOp1Response, error) {
	res, err := airkorea.Get(statistics.Op1(statistics.Op1Request{
		StationName:     in.GetStationName(),
		SearchCondition: in.GetSearchCondition(),
		PageNo:          in.GetPageNo(),
		NumOfRows:       in.GetNumOfRows(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("StatisticsOp1Service failed due to some unknown reason")
	}

	items := make([]*pb.StatisticsOp1Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*statistics.Op1ResponseItem)
		items = append(items, &pb.StatisticsOp1Response_Item{
			DataTime:    item.DataTime,
			So2Average:  item.So2Average,
			CoAverage:   item.CoAverage,
			O3Average:   item.O3Average,
			No2Average:  item.No2Average,
			Pm10Average: item.Pm10Average,
		})
	}

	return &pb.StatisticsOp1Response{
		Items: items,
	}, nil
}

func (s *server) StatisticsOp2Service(ctx context.Context, in *pb.StatisticsOp2Request) (*pb.StatisticsOp2Response, error) {
	res, err := airkorea.Get(statistics.Op2(statistics.Op2Request{
		SearchDataTime:       in.GetSearchDataTime(),
		StatArticleCondition: in.GetStatArticleCondition(),
		PageNo:               in.GetPageNo(),
		NumOfRows:            in.GetNumOfRows(),
	}))

	if err != nil {
		log.Fatalf("%v", err)
		return nil, errors.New("StatisticsOp2Service failed due to some unknown reason")
	}

	items := make([]*pb.StatisticsOp2Response_Item, 0)
	for _, v := range res.GetItem() {
		item := v.(*statistics.Op2ResponseItem)
		items = append(items, &pb.StatisticsOp2Response_Item{
			SidoName:    item.SidoName,
			DataTime:    item.DataTime,
			So2Average:  item.So2Average,
			CoAverage:   item.CoAverage,
			O3Average:   item.O3Average,
			No2Average:  item.No2Average,
			Pm10Average: item.Pm10Average,
			So2Max:      item.So2Max,
			CoMax:       item.CoMax,
			O3Max:       item.O3Max,
			No2Max:      item.No2Max,
			Pm10Max:     item.Pm10Max,
			So2Min:      item.So2Min,
			CoMin:       item.CoMin,
			O3Min:       item.O3Min,
			No2Min:      item.No2Min,
			Pm10Min:     item.Pm10Min,
		})
	}

	return &pb.StatisticsOp2Response{
		Items: items,
	}, nil
}

var (
	logrusLogger *logrus.Logger
	customFunc   grpc_logrus.CodeToLevel
)

func init() {
	// create collector.
	collector, err := zipkin.NewHTTPCollector("http://localhost:9411/api/v1/spans")
	if err != nil {
		fmt.Printf("unable to create Zipkin HTTP collector: %+v\n", err)
		os.Exit(-1)
	}

	// create recorder.
	recorder := zipkin.NewRecorder(collector, false, "127.0.0.1:50051", "forecast_server")

	// create tracer.
	tracer, err := zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)
	if err != nil {
		fmt.Printf("unable to create Zipkin tracer: %+v\n", err)
		os.Exit(-1)
	}

	// explicitly set our tracer to be the default tracer.
	opentracing.InitGlobalTracer(tracer)
}

func main() {
	/*
		keyFile := flag.String("key", "secure", "Input key file")
		rpcPort := flag.Int("port", 50051, "RPC server port")
		flag.Parse()

		if key, err := ioutil.ReadFile(*keyFile); err != nil {
			log.Fatalln(err)
		} else {
			log.Println("Success loading service key")
			airkorea.SetKey(string(key))
		}
	*/
	key := os.Getenv("KEY")
	if key == "" {
		log.Fatal("KEY environment variable was not set")
	}
	airkorea.SetKey(key)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":"+port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Logrus entry is used, allowing pre-definition of certain fields by the user.
	logrusLogger := &logrus.Logger{
		Out:          os.Stderr,
		Formatter:    new(logrus.JSONFormatter),
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
	logrusEntry := logrus.NewEntry(logrusLogger)
	// Shared options for the logger, with a custom gRPC code to log level function.

	customFunc = grpc_logrus.DefaultCodeToLevel
	opts := []grpc_logrus.Option{
		grpc_logrus.WithLevels(customFunc),
	}

	// Make sure that log statements internal to gRPC library are logged using the logrus Logger as well.
	grpc_logrus.ReplaceGrpcLogger(logrusEntry)

	s := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...),
			grpc_opentracing.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_logrus.StreamServerInterceptor(logrusEntry, opts...),
			grpc_opentracing.StreamServerInterceptor(),
		),
	)

	pb.RegisterForecasterServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
