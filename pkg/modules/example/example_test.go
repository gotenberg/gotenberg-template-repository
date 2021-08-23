package example

import (
	"reflect"
	"testing"

	"github.com/gotenberg/gotenberg/v7/pkg/gotenberg"
)

func TestExample_Descriptor(t *testing.T) {
	descriptor := Example{}.Descriptor()

	actual := reflect.TypeOf(descriptor.New())
	expect := reflect.TypeOf(new(Example))

	if actual != expect {
		t.Errorf("expected '%s' but got '%s'", expect, actual)
	}
}

func TestExample_Provision(t *testing.T) {
	for i, tc := range []struct {
		ctx       *gotenberg.Context
		expectErr bool
	}{
		{
			ctx: func() *gotenberg.Context {
				return gotenberg.NewContext(
					gotenberg.ParsedFlags{
						FlagSet: new(Example).Descriptor().FlagSet,
					},
					[]gotenberg.ModuleDescriptor{},
				)
			}(),
		},
	} {
		mod := new(Example)
		err := mod.Provision(tc.ctx)

		if tc.expectErr && err == nil {
			t.Errorf("test %d: expected error but got: %v", i, err)
		}

		if !tc.expectErr && err != nil {
			t.Errorf("test %d: expected no error but got: %v", i, err)
		}
	}
}

func TestExample_Validate(t *testing.T) {
	for i, tc := range []struct {
		strProp   string
		intProp   int
		expectErr bool
	}{
		{
			strProp:   "bar",
			expectErr: true,
		},
		{
			intProp:   1337,
			expectErr: true,
		},
		{
			strProp: "foo",
			intProp: 10,
		},
	} {
		mod := Example{
			strProp: tc.strProp,
			intProp: tc.intProp,
		}

		err := mod.Validate()

		if tc.expectErr && err == nil {
			t.Errorf("test %d: expected error but got: %v", i, err)
		}

		if !tc.expectErr && err != nil {
			t.Errorf("test %d: expected no error but got: %v", i, err)
		}
	}
}
