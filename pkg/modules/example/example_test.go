package example

import (
	"reflect"
	"testing"

	"github.com/gotenberg/gotenberg/v8/pkg/gotenberg"
)

func TestExample_Descriptor(t *testing.T) {
	descriptor := new(Example).Descriptor()

	actual := reflect.TypeOf(descriptor.New())
	expect := reflect.TypeFor[*Example]()

	if actual != expect {
		t.Errorf("expected '%s' but got '%s'", expect, actual)
	}
}

func TestExample_Provision(t *testing.T) {
	for _, tc := range []struct {
		scenario    string
		ctx         *gotenberg.Context
		expectError bool
	}{
		{
			scenario: "default flags",
			ctx: func() *gotenberg.Context {
				return gotenberg.NewContext(
					gotenberg.ParsedFlags{
						FlagSet: new(Example).Descriptor().FlagSet,
					},
					[]gotenberg.ModuleDescriptor{},
				)
			}(),
			expectError: false,
		},
	} {
		t.Run(tc.scenario, func(t *testing.T) {
			mod := new(Example)
			err := mod.Provision(tc.ctx)

			if tc.expectError && err == nil {
				t.Fatal("expected error but got none")
			}

			if !tc.expectError && err != nil {
				t.Fatalf("expected no error but got: %v", err)
			}
		})
	}
}

func TestExample_Validate(t *testing.T) {
	for _, tc := range []struct {
		scenario    string
		strProp     string
		intProp     int
		expectError bool
	}{
		{
			scenario:    "invalid strProp",
			strProp:     "bar",
			intProp:     1,
			expectError: true,
		},
		{
			scenario:    "invalid intProp",
			strProp:     "foo",
			intProp:     1337,
			expectError: true,
		},
		{
			scenario:    "successful validation",
			strProp:     "foo",
			intProp:     10,
			expectError: false,
		},
	} {
		t.Run(tc.scenario, func(t *testing.T) {
			mod := &Example{
				strProp: tc.strProp,
				intProp: tc.intProp,
			}

			err := mod.Validate()

			if tc.expectError && err == nil {
				t.Fatal("expected error but got none")
			}

			if !tc.expectError && err != nil {
				t.Fatalf("expected no error but got: %v", err)
			}
		})
	}
}
