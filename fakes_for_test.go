package protofm_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"reflect"

	"github.com/SecuritasCrimePrediction/protofm"

	"github.com/SecuritasCrimePrediction/protofm/testproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

//
// Fake service to test StreamServerInterceptor
//
type fakeService struct {
	testproto.UnimplementedTestServiceServer
}

var _ testproto.TestServiceServer = fakeService{}

func (f fakeService) TestStreamFieldMask(req *testproto.FakeFieldMaskRequest,
	stream testproto.TestService_TestStreamFieldMaskServer) error {
	if req.GetRetError() != "" {
		return fmt.Errorf(req.GetRetError())
	}

	for i := 0; i < int(req.GetNrResponses()); i++ {
		if err := stream.SendMsg(req.GetWantedResponse()); err != nil {
			return err
		}
	}
	return nil
}

func (f fakeService) TestStreamNoFieldMask(req *testproto.NoFieldMaskRequest,
	stream testproto.TestService_TestStreamNoFieldMaskServer) error {
	if req.GetRetError() != "" {
		return fmt.Errorf(req.GetRetError())
	}

	for i := 0; i < int(req.GetNrResponses()); i++ {
		if err := stream.SendMsg(req.GetWantedResponse()); err != nil {
			return err
		}
	}
	return nil
}

//
// Fake server to test StreamServerInterceptor
//
const bufSize = 1024 * 1024

var lis *bufconn.Listener
var s *grpc.Server

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func StartNewFakeServer() {
	lis = bufconn.Listen(bufSize)
	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(protofm.StreamServerInterceptor()),
	}
	s = grpc.NewServer(opts...)
	testproto.RegisterTestServiceServer(s, fakeService{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func StopFakeServer() {
	s.Stop()
}

type SimpleObjectStream interface {
	Recv() (*testproto.SimpleObject, error)
}

//
// Fakes for testing UnaryServerInterceptor
//
type FieldMaskForTest struct {
	FieldMask     []string
	Response      interface{}
	ErrorResponse error
	NrResponses   int
}

func (f FieldMaskForTest) GetReadMask() []string {
	return f.FieldMask
}

type ResponseStruct struct {
	A string
	B int32
}

func FakeUnaryHandler(ctx context.Context, req interface{}) (interface{}, error) {
	var respValue interface{}
	var errValue error

	val := reflect.ValueOf(req)
	resp := val.FieldByName("Response")
	if resp.IsValid() {
		respValue = resp.Interface()
	}

	err := val.FieldByName("ErrorResponse")
	if err.IsValid() {
		if err.Interface() != nil {
			errValue = err.Interface().(error)
		}
	}

	return respValue, errValue
}
