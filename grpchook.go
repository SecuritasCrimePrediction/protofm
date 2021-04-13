package protofm

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

// ReadMaskable must be implemented by the requests for the interceptors to apply the field map on the responses.
type ReadMaskable interface {
	GetReadMask() []string
}

// UnaryServerInterceptor will apply field masks on responses for unary endpoints if the request
// * ... implement the ReadMaskable interface
// * ... the field mask paths are not empty
// * ... the field mask paths are valid on the response message
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		// get the response
		resp, err := handler(ctx, req)
		if err != nil {
			return resp, err
		}

		if fm, ok := req.(ReadMaskable); ok {
			// cast to proto message if possible
			protoResp, isProtoResponse := resp.(proto.Message)
			if !isProtoResponse {
				return resp, err
			}

			// filter the response
			mask := NewMask(fm.GetReadMask())
			if mask.Validate(protoResp) {
				// only apply field mask if it is valid
				mask.Apply(protoResp)

				// set the filtered response
				resp = protoResp
			}
		}

		return resp, err
	}
}

// StreamServerInterceptor will apply field masks on all streamed responses from Server streamed endpoints if the request
// * ... implement the ReadMaskable interface
// * ... the field mask paths are not empty
// * ... the field mask paths are valid on the response message
func StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		// wrap the stream to be able to apply field masks
		s := &streamWrapper{
			wrappedStream: stream,
		}

		return handler(srv, s)
	}
}

// The field mask stream wrapper for StreamServerInterceptor
type streamWrapper struct {
	wrappedStream grpc.ServerStream
	mask          FieldMaskMap
	paths         []string
	validated     bool
}

func (w *streamWrapper) SetHeader(md metadata.MD) error {
	return w.wrappedStream.SetHeader(md)
}

func (w *streamWrapper) SendHeader(md metadata.MD) error {
	return w.wrappedStream.SendHeader(md)
}

func (w *streamWrapper) SetTrailer(md metadata.MD) {
	w.wrappedStream.SetTrailer(md)
}

func (w *streamWrapper) Context() context.Context {
	return w.wrappedStream.Context()
}

func (w *streamWrapper) RecvMsg(m interface{}) error {
	if err := w.wrappedStream.RecvMsg(m); err != nil {
		return err
	}

	protoMsg, isProto := m.(proto.Message)
	if !isProto {
		return nil
	}

	if sub, ok := protoMsg.(ReadMaskable); ok {
		w.paths = sub.GetReadMask()
	}

	return nil
}

func (w *streamWrapper) SendMsg(m interface{}) error {
	protoMsg, isProto := m.(proto.Message)
	if !isProto {
		return w.wrappedStream.SendMsg(m)
	}

	// create a validate field mask when applicable
	if w.mask == nil && !w.validated {
		w.validated = true
		if w.paths != nil {
			mask := NewMask(w.paths)
			if mask.Validate(protoMsg) {
				w.mask = mask
			}
		}
	}

	// apply the field mask if it is set
	if w.mask != nil {
		w.mask.Apply(protoMsg)
	}

	// set the filtered response
	m = protoMsg

	// send the filtered response
	return w.wrappedStream.SendMsg(m)
}
