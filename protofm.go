package protofm

import (
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// FieldMaskMap holds the information about which fields should be filtered
// in a protobuf message.
type FieldMaskMap map[string]FieldMaskMap

// NewMask creates a FieldMaskMap from the paths provided.
// Path parts are separated by "." for example "base.sub.key".
func NewMask(paths []string) FieldMaskMap {
	ret := FieldMaskMap{}

	for _, path := range paths {
		// get all parts of the path
		parts := strings.Split(strings.ToLower(path), ".")
		currMap := ret
		for _, part := range parts {
			// fetch the map for the given part
			fieldMap, ok := currMap[part]
			if !ok {
				// create it if it does not exist
				currMap[part] = FieldMaskMap{}
				fieldMap = currMap[part]
			}
			currMap = fieldMap
		}
	}

	return ret
}

// ValidateMask validates that the paths are valid paths in the given protobuf message.
// A FieldMaskMap is created and then validated against the message.
func ValidateMask(msg proto.Message, paths []string) bool {
	if msg == nil {
		return false
	}
	return NewMask(paths).Validate(msg)
}

// ApplyMask keep the value of all given paths in the given proto message and clears the rest.
// Does not validate the field mask, this should be done before calling this function.
// A FieldMaskMap is created and then used for filtering the given proto message.
func ApplyMask(msg proto.Message, paths []string) {
	NewMask(paths).Apply(msg)
}

// ValidateMask the FieldMaskMap against the given proto message.
func (fm FieldMaskMap) Validate(msg proto.Message) bool {
	msgReflect := msg.ProtoReflect()
	msgDescription := msgReflect.Descriptor()
	for fieldName, mask := range fm {
		field := msgDescription.Fields().ByName(protoreflect.Name(strings.ToLower(fieldName)))
		if field == nil {
			return false
		}
		if len(mask) > 0 {
			if field.IsList() {
				elm := msgReflect.Get(field).List().NewElement()
				if _, ok := elm.Interface().(protoreflect.Message); !ok {
					return false
				}

				if !mask.Validate(elm.Message().Interface()) {
					return false
				}
			} else {
				if field.Kind() == protoreflect.MessageKind {
					if !mask.Validate(msgReflect.Get(field).Message().Interface()) {
						return false
					}
				} else {
					return false
				}
			}
		}
	}

	return true
}

// Apply the given proto message with the FieldMaskMap.
// Does not validate the field mask, this should be done before calling this function.
// Will keep all fields specified in the FieldMaskMap as is and clear the rest.
// If there are no paths in the FieldMaskMap all fields will be left as is.
func (fm FieldMaskMap) Apply(msg proto.Message) {
	if len(fm) == 0 {
		return
	}

	msgReflect := msg.ProtoReflect()
	msgReflect.Range(func(descriptor protoreflect.FieldDescriptor, value protoreflect.Value) bool {
		name := string(descriptor.Name())

		// check if this message should be filtered
		mask, ok := fm[name]
		if ok {
			if len(mask) == 0 {
				// no children to filter, keep the message as is
				return true
			}
			if descriptor.IsList() {
				// get the list
				msgList := msgReflect.Get(descriptor).List()

				for i := 0; i < msgList.Len(); i++ {
					// filter each message in the list
					mask.Apply(msgList.Get(i).Message().Interface())
				}
			} else {
				// filter the message
				mask.Apply(msgReflect.Get(descriptor).Message().Interface())
			}

		} else {
			// message should not be kept
			msgReflect.Clear(descriptor)
		}

		return true
	})
}
