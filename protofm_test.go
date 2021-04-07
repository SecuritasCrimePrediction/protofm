package protofm_test

import (
	"testing"

	"github.com/SecuritasCrimePrediction/protofm"
	"github.com/SecuritasCrimePrediction/protofm/testproto"

	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/proto"
	"gotest.tools/assert/cmp"
)

func Test_CreateNestedMap(t *testing.T) {
	input := []string{"a", "b", "c.X"}
	want := protofm.FieldMaskMap{
		"b": protofm.FieldMaskMap{},
		"a": protofm.FieldMaskMap{},
		"c": protofm.FieldMaskMap{
			"x": protofm.FieldMaskMap{},
		},
	}

	got := protofm.NewMask(input)

	c := cmp.DeepEqual(want, got)

	if !c().Success() {
		t.Errorf("not the same\nwanted: %v\ngot: %v", want, got)
	}
}

func Test_Filter(t *testing.T) {
	for _, tc := range []struct {
		reason string
		msg    proto.Message
		filter []string
		want   proto.Message
	}{
		{
			reason: "should be able to have all fields of the message in the filter",
			msg:    &testproto.SimpleObject{Pow: "pow", Wow: "wow", Foo: 1, Baz: 2},
			filter: []string{"pow", "wow", "foo", "baz"},
			want:   &testproto.SimpleObject{Pow: "pow", Wow: "wow", Foo: 1, Baz: 2},
		}, {
			reason: "should not matter which order the paths are in the filter",
			msg:    &testproto.SimpleObject{Pow: "pow", Wow: "wow", Foo: 1, Baz: 2},
			filter: []string{"baz", "foo", "pow", "wow"},
			want:   &testproto.SimpleObject{Pow: "pow", Wow: "wow", Foo: 1, Baz: 2},
		}, {
			reason: "should be able to have an empty filter and get all fields back",
			msg:    &testproto.SimpleObject{Pow: "pow", Wow: "wow", Foo: 1, Baz: 2},
			filter: []string{},
			want:   &testproto.SimpleObject{Pow: "pow", Wow: "wow", Foo: 1, Baz: 2},
		}, {
			reason: "should be able to have nil filter and get all fields back",
			msg:    &testproto.SimpleObject{Pow: "pow", Wow: "wow", Foo: 1, Baz: 2},
			filter: nil,
			want:   &testproto.SimpleObject{Pow: "pow", Wow: "wow", Foo: 1, Baz: 2},
		}, {
			reason: "should be able to get only selected fields",
			msg:    &testproto.SimpleObject{Pow: "pow", Wow: "wow", Foo: 1, Baz: 2},
			filter: []string{"baz", "pow"},
			want:   &testproto.SimpleObject{Pow: "pow", Baz: 2},
		}, {
			reason: "should be able to specify filter for empty fields without errors",
			msg:    &testproto.SimpleObject{Pow: "pow", Foo: 1},
			filter: []string{"baz", "pow"},
			want:   &testproto.SimpleObject{Pow: "pow"},
		}, {
			reason: "should be able to handle complex structures",
			msg: &testproto.NestedObject{
				Pow: 1,
				Wow: "wow",
				FooBaz: []*testproto.NestedObject_FooBaz{
					{Foo: 1, Baz: "baz"},
					{Foo: 1, Baz: "baz"},
				},
				ComplexObject: &testproto.ComplexObject{
					RepeatedAndSingleValue: &testproto.RepeatedAndSingle{
						SingleValue:   &testproto.Single{Value: "A"},
						RepeatedValue: []string{"A", "B"},
					},
					SingleValue: "C",
				},
			},
			filter: []string{"pow", "foo_baz.foo", "complex_object.repeated_and_single_value.repeated_value", "complex_object.single_value"},
			want: &testproto.NestedObject{
				Pow:    1,
				FooBaz: []*testproto.NestedObject_FooBaz{{Foo: 1}, {Foo: 1}},
				ComplexObject: &testproto.ComplexObject{
					RepeatedAndSingleValue: &testproto.RepeatedAndSingle{
						RepeatedValue: []string{"A", "B"},
					},
					SingleValue: "C",
				},
			},
		},
	} {
		t.Run(tc.reason, func(t *testing.T) {
			protofm.ApplyMask(tc.msg, tc.filter)
			c := cmp.DeepEqual(tc.want, tc.msg, cmpopts.IgnoreUnexported(
				testproto.NestedObject{}, testproto.SimpleObject{},
				testproto.NestedObject_FooBaz{}, testproto.ComplexObject{},
				testproto.Single{}, testproto.RepeatedAndSingle{},
			))

			if !c().Success() {
				t.Errorf("not the same\nwanted: %v\ngot: %v", tc.want, tc.msg)
			}
		})
	}
}

func Test_Validate(t *testing.T) {
	for _, tc := range []struct {
		reason string
		msg    proto.Message
		filter []string
		want   bool
	}{
		{
			"all fields may be set in the filter",
			&testproto.SimpleObject{},
			[]string{"pow", "wow", "foo", "baz"},
			true,
		}, {
			"none of the fields may be in the filter",
			&testproto.SimpleObject{},
			[]string{},
			true,
		}, {
			"filter may be nil",
			&testproto.SimpleObject{},
			nil,
			true,
		}, {
			"some fields may be set in the filter",
			&testproto.SimpleObject{},
			[]string{"wow", "baz"},
			true,
		}, {
			"should handle complex proto messages",
			&testproto.NestedObject{},
			[]string{"pow", "foo_baz.foo", "complex_object.repeated_and_single_value.repeated_value", "complex_object.single_value"},
			true,
		}, {
			"should fail if requested field does not exist on single value",
			&testproto.NestedObject{},
			[]string{"pow", "foo_baz.foo", "complex_object.repeated_and_single_value.repeated_value", "complex_object.single_value.doesNotExist"},
			false,
		}, {
			"should fail if requested field does not exist in the base of the object",
			&testproto.NestedObject{},
			[]string{"pow", "foo_baz.foo", "complex_object.repeated_and_single_value.repeated_value", "complex_object.single_value", "doesNotExist"},
			false,
		}, {
			"should fail if requested field does not exist in a repeated value",
			&testproto.NestedObject{},
			[]string{"pow", "foo_baz.foo", "complex_object.repeated_and_single_value.repeated_value.doesNotExist", "complex_object.single_value"},
			false,
		}, {
			"can't validate nil message",
			nil,
			[]string{"some", "strings"},
			false,
		},
	} {
		if got := protofm.ValidateMask(tc.msg, tc.filter); got != tc.want {
			t.Errorf("expected valid to be %t but got %t, reason: %s", tc.want, got, tc.reason)
		}
	}
}
