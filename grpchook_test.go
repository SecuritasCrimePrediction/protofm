package protofm_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/SecuritasCrimePrediction/protofm"
	"github.com/SecuritasCrimePrediction/protofm/testproto"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert/cmp"
)

func Test_UnaryFilter(t *testing.T) {
	unaryInterceptor := protofm.UnaryServerInterceptor()

	for _, tc := range []struct {
		reason   string
		req      interface{}
		wantResp interface{}
		wantErr  error
	}{
		{
			reason: "should not filter if request does not implement FieldMaskable interface",
			req: struct {
				Response       interface{}
				WrongFieldMask []string
			}{
				Response: &testproto.SimpleObject{
					Pow: "pow",
					Wow: "wow",
					Foo: 1,
					Baz: 2,
				},
				WrongFieldMask: []string{"pow"},
			},
			wantResp: &testproto.SimpleObject{
				Pow: "pow",
				Wow: "wow",
				Foo: 1,
				Baz: 2,
			},
			wantErr: nil,
		}, {
			reason: "should not filter if service return error",
			req: FieldMaskForTest{
				FieldMask:     []string{"A"},
				Response:      nil,
				ErrorResponse: fmt.Errorf("return error"),
			},
			wantResp: nil,
			wantErr:  fmt.Errorf("return error"),
		}, {
			reason: "should not filter if response is not proto message",
			req: FieldMaskForTest{
				FieldMask: []string{"A"},
				Response: &ResponseStruct{
					A: "A",
					B: 42,
				},
				ErrorResponse: nil,
			},
			wantResp: &ResponseStruct{
				A: "A",
				B: 42,
			},
			wantErr: nil,
		}, {
			reason: "should filter if request implement FieldMaskable interface, response is proto message and service return no error",
			req: FieldMaskForTest{
				FieldMask: []string{"pow", "foo"},
				Response: &testproto.SimpleObject{
					Pow: "pow",
					Wow: "wow",
					Foo: 1,
					Baz: 2,
				},
				ErrorResponse: nil,
			},
			wantResp: &testproto.SimpleObject{
				Pow: "pow",
				Foo: 1,
			},
			wantErr: nil,
		}, {
			reason: "should not filter if requested field mask is not valid",
			req: FieldMaskForTest{
				FieldMask: []string{"pow", "doesNotExist"},
				Response: &testproto.SimpleObject{
					Pow: "pow",
					Wow: "wow",
					Foo: 1,
					Baz: 2,
				},
				ErrorResponse: nil,
			},
			wantResp: &testproto.SimpleObject{
				Pow: "pow",
				Wow: "wow",
				Foo: 1,
				Baz: 2,
			},
			wantErr: nil,
		},
	} {
		gotResp, gotErr := unaryInterceptor(context.Background(), tc.req, &grpc.UnaryServerInfo{}, FakeUnaryHandler)

		if tc.wantErr != nil {
			if gotErr == nil {
				t.Errorf("expected error, got nil, reason: %s", tc.reason)
			} else {
				c := cmp.Error(gotErr, tc.wantErr.Error())
				if !c().Success() {
					t.Errorf("want error: %s\ngot: %s\nreason: %s", tc.wantErr.Error(), gotErr.Error(), tc.reason)
				}
			}
		} else {
			if gotErr != nil {
				t.Errorf("expected no error, got %v\nreason: %s", gotErr, tc.reason)
			}
		}

		c := cmp.DeepEqual(tc.wantResp, gotResp, cmpopts.IgnoreUnexported(testproto.SimpleObject{}))
		if !c().Success() {
			t.Errorf("want response: %v\ngot: %v\nreason: %s", tc.wantResp, gotResp, tc.reason)
		}
	}
}

func Test_StreamFilter(t *testing.T) {
	StartNewFakeServer()
	defer StopFakeServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "fakeServer", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial fake server: %v", err)
	}
	defer conn.Close()
	client := testproto.NewTestServiceClient(conn)

	for _, tc := range []struct {
		reason     string
		req        proto.Message
		wantResp   *testproto.SimpleObject
		wantNrResp int
		wantErr    error
	}{
		{
			reason: "should not filter if request does not implement FieldMaskable interface",
			req: &testproto.NoFieldMaskRequest{
				NrResponses: 3,
				WantedResponse: &testproto.SimpleObject{
					Pow: "pow",
					Wow: "wow",
					Foo: 1,
					Baz: 2,
				},
			},
			wantResp: &testproto.SimpleObject{
				Pow: "pow",
				Wow: "wow",
				Foo: 1,
				Baz: 2,
			},
			wantNrResp: 3,
			wantErr:    nil,
		}, {
			reason: "should not filter if service return error",
			req: &testproto.FakeFieldMaskRequest{
				FieldMask:   []string{"pow"},
				NrResponses: 3,
				RetError:    "return error",
				WantedResponse: &testproto.SimpleObject{
					Pow: "pow",
					Wow: "wow",
					Foo: 1,
					Baz: 2,
				},
			},
			wantResp:   nil,
			wantNrResp: 0,
			wantErr:    fmt.Errorf("return error"),
		}, {
			reason: "should filter if request implement FieldMaskable interface and service return no error",
			req: &testproto.FakeFieldMaskRequest{
				FieldMask:   []string{"pow", "baz"},
				NrResponses: 3,
				WantedResponse: &testproto.SimpleObject{
					Pow: "pow",
					Wow: "wow",
					Foo: 1,
					Baz: 2,
				},
			},
			wantResp: &testproto.SimpleObject{
				Pow: "pow",
				Baz: 2,
			},
			wantNrResp: 3,
			wantErr:    nil,
		}, {
			reason: "should not filter if requested field mask is not valid",
			req: &testproto.FakeFieldMaskRequest{
				FieldMask:   []string{"pow", "doesNotExist"},
				NrResponses: 3,
				WantedResponse: &testproto.SimpleObject{
					Pow: "pow",
					Wow: "wow",
					Foo: 1,
					Baz: 2,
				},
			},
			wantResp: &testproto.SimpleObject{
				Pow: "pow",
				Wow: "wow",
				Foo: 1,
				Baz: 2,
			},
			wantNrResp: 3,
			wantErr:    nil,
		},
	} {
		var err error
		var stream SimpleObjectStream
		if req, ok := tc.req.(*testproto.FakeFieldMaskRequest); ok {
			stream, err = client.TestStreamFieldMask(context.Background(), req)
		}
		if req, ok := tc.req.(*testproto.NoFieldMaskRequest); ok {
			stream, err = client.TestStreamNoFieldMask(context.Background(), req)
		}

		if err != nil {
			t.Errorf("unexpected error when calling server, err: %v", err)
			continue
		}

		nrResponses := 0
		for {
			resp, gotErr := stream.Recv()
			if tc.wantErr != nil {
				if gotErr == nil {
					t.Errorf("expected error, got nil, reason: %s", tc.reason)
				} else {
					st, ok := status.FromError(gotErr)
					if !ok {
						t.Error("error was not a status error")
						break
					}
					c := cmp.Equal(st.Message(), tc.wantErr.Error())
					if !c().Success() {
						t.Errorf("want error: %s\ngot: %s\nreason: %s", tc.wantErr.Error(), gotErr.Error(), tc.reason)
					}
					break
				}
			} else {
				if gotErr != nil {
					if errors.Is(gotErr, io.EOF) {
						break
					}
					t.Errorf("expected no error, got %v\nreason: %s", gotErr, tc.reason)
					break
				}
			}

			c := cmp.DeepEqual(tc.wantResp, resp, cmpopts.IgnoreUnexported(testproto.SimpleObject{}))
			if !c().Success() {
				t.Errorf("want response: %v\ngot: %v\nreason: %s", tc.wantResp, resp, tc.reason)
			}
			nrResponses++
		}

		if nrResponses != tc.wantNrResp {
			t.Errorf("wanted %d responses, got: %d\nreason: %s", tc.wantNrResp, nrResponses, tc.reason)
		}
	}
}
