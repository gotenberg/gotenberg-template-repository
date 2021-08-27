package example

import (
	"fmt"

	"github.com/gotenberg/gotenberg/v7/pkg/gotenberg"
	flag "github.com/spf13/pflag"
	"go.uber.org/multierr"
)

func init() {
	gotenberg.MustRegisterModule(Example{})
}

// Example is our module. It does nothing.
type Example struct {
	strProp string
	intProp int
}

// Descriptor returns an Example's module descriptor.
func (mod Example) Descriptor() gotenberg.ModuleDescriptor {
	return gotenberg.ModuleDescriptor{
		ID: "example",
		FlagSet: func() *flag.FlagSet {
			fs := flag.NewFlagSet("example", flag.ExitOnError)
			fs.String("example-str-prop", "", "Set the string property")
			fs.Int("example-int-prop", 0, "Set the int property")

			return fs
		}(),
		New: func() gotenberg.Module { return new(Example) },
	}
}

// Provision sets the module properties.
func (mod *Example) Provision(ctx *gotenberg.Context) error {
	flags := ctx.ParsedFlags()
	mod.strProp = flags.MustString("example-str-prop")
	mod.intProp = flags.MustInt("example-int-prop")

	return nil
}

// Validate validates the module properties.
func (mod Example) Validate() error {
	var err error

	if mod.strProp == "bar" {
		err = multierr.Append(err, fmt.Errorf("str prop must be different than bar"))
	}

	if mod.intProp == 1337 {
		err = multierr.Append(err, fmt.Errorf("int prop must be different than 1337"))
	}

	return err
}

// Interface guards.
var (
	_ gotenberg.Module      = (*Example)(nil)
	_ gotenberg.Provisioner = (*Example)(nil)
	_ gotenberg.Validator   = (*Example)(nil)
)
