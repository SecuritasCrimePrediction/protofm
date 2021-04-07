# Proto field mask 

[![Go Reference](https://pkg.go.dev/badge/github.com/SecuritasCrimePrediction/protofm.svg)](https://pkg.go.dev/github.com/SecuritasCrimePrediction/protofm)

[Field masks](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/field-mask) allows gRPC consumers to select fields of interest to keep in the response. This reduces the amount of data being transferred from the server to the client.

This library extends the formal field mask definition by adding support for masking of fields inside repeated messages.

This library has the following features:

* Field mask application with list support (`protofm.Apply`)
* Field mask validation (`protofm.Validate`)
* gRPC hook for convention-based application of field masks.

## Background

A field mask is a list of strings with dot-separated field names from the response message.

Example

```
f {
  a : 22
  b {
    d : 1
    x : 2
  }
  y : 13
}
z: 8
```

Applying the field mask `["f.a", "f.b.d"]` the above message gives:

```
f {
  a : 22
  b {
    d : 1
  }
}
```

If a field mask is missing or empty, no filtering takes place.

## Applying a field mask

NOTE: it is **strongly recommended** to validate field masks before application to a Protobuf message. See `protofm.Validate` below.

There are two ways of applying a field mask on a proto message, either by sending both proto message and field mask paths to the `ApplyMask` function...

```go
protofm.ApplyMask(pbMessage, []string{"f.a", "f.b.d"})
```

...or the preferred way is to first create a new instance of FieldMaskMap, validate it and then apply it on the proto message:

```go
fm := protofm.NewMask([]string{"f.a", "f.b.d"})
if valid := fm.Validate(pbMessage); !valid {
    fmt.Println("the field mask was not valid for this message type")
}
fm.Apply(pbMessage)
```

## Validating a field mask

`protofm.Validate` returns `false` if one or more of the paths in the field mask are not available in the proto message.

Validate is called either with proto message and field mask paths to the `protofm.ValidateMask` function...

```go
if valid := protofm.ValidateMask(pbMessage, []string{"f.a", "f.b.d"}); !valid {
    fmt.Println("the field mask was not valid for this message type")
}
```

...or with an existing `protofm.FieldMaskMap`:

```go
fm := protofm.NewMask([]string{"f.a", "f.b.d"})
if valid := fm.Validate(pbMessage); !valid {
    fmt.Println("the field mask was not valid for this message type")
}
```

## gRPC hook

To automatically validate and apply field masks when `.field_mask` is present and non-empty on the request, add the following interceptors:

```go
opts := []grpc.ServerOption{
    grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
        protofm.UnaryServerInterceptor(),
    )),
    grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
        protofm.StreamServerInterceptor(),
    )),
}

grpcServer = grpc.NewServer(opts...)
```

Under the hood, this checks whether the incoming request is `FieldMaskable`:

```protobuf
type FieldMaskable interface {
    GetFieldMask() []string
}
```

Example proto definition that would trigger the interceptor:

```protobuf
message ExampleRequest {
    repeated string field_mask = 1;
    string some_other_field = 2;
}
```