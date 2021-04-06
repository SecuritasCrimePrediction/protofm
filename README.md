# Proto field mask 

Field masks can be used by API callers to list which fields in the response they are interested in, and let the server remove the rest of the fields to reduce the amount of data being transferred from the server.
This library is created to be able to apply field masks not only on simple proto messages but also proto messages in lists.

# Field mask paths
The field mask is created with a list of message field definitions which are dot notated strings.\
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
With a field mask `f.a, f.b.d` applied to the above message the result will be
```
f {
  a : 22
  b {
    d : 1
  }
}
```
If the list of paths is empty the message will not be filtered.

## Validate
There is a validation function for the field mask to validate it against a proto message.
The validate function return `false` if one or more of the paths in the field mask are not available in the proto message.
The validation function can be called in two ways, either by sending both proto message and field mask paths to the `ValidateMask` function
```
if valid := protofm.ValidateMask(pbMessage, []string{"f.a", "f.b.d"}); !valid {
    fmt.Println("the field mask was not valid for this message type")
}
```
or if you are going to use the field mask to apply to messages if it is valid you can first create a new instance of the `FieldMaskMap` and then validate it against a proto message
```
fm := protofm.NewMask([]string{"f.a", "f.b.d"})
if valid := fm.Validate(pbMessage); !valid {
    fmt.Println("the field mask was not valid for this message type")
}
```

## Apply
The functions to apply the field mask on a proto message does not validate the field mask against the message, the user should do this before applying.

There are two ways of applying the field mask on a proto message, either by sending both proto message and field mask paths to the `ApplyMask` function
```
protofm.ApplyMask(pbMessage, []string{"f.a", "f.b.d"})
```
or the preferred way is to first create a new instance of FieldMaskMap, validate it and then apply it on the proto message
```
fm := protofm.NewMask([]string{"f.a", "f.b.d"})
if valid := fm.Validate(pbMessage); !valid {
    fmt.Println("the field mask was not valid for this message type")
}
fm.Apply(pbMessage)
```

# gRPC hooks
This library contains gRPC hooks that you can use for your gRPC servers to apply field masks on the responses from the endpoints.
For responses to be filtered using field mask the request message need to implement the `FieldMaskable` interface which looks like this
```
type FieldMaskable interface {
    GetFieldMask() []string
}
```
For example a proto message definition of a request that will filter the response can look like this
```
message ExampleRequest {
    repeated string field_mask = 1;
    string some_other_field = 2;
}
```
The responses will be filtered if
* ... the interceptors are added to the gRPC server options
* ... the request implements the `FieldMaskable` interface
* ... the field mask paths are not empty
* ... the field mask paths are valid on the response message

There are interceptors for both Server stream and Unary endpoints.
To add them to the gRPC server options add them like this
```
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
