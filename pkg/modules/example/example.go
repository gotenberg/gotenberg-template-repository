package example

import (
	"errors"

	flag "github.com/spf13/pflag"
	"go.uber.org/multierr"

	"github.com/gotenberg/gotenberg/v7/pkg/gotenberg"
)

func init() {
	gotenberg.MustRegisterModule(new(Example))
}

// Example is our module. It does nothing.
type Example struct {
	strProp string
	intProp int
}

// Descriptor returns an [Example]'s module descriptor.
func (ex *Example) Descriptor() gotenberg.ModuleDescriptor {
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
func (ex *Example) Provision(ctx *gotenberg.Context) error {
	flags := ctx.ParsedFlags()
	ex.strProp = flags.MustString("example-str-prop")
	ex.intProp = flags.MustInt("example-int-prop")

	return nil
}

// Validate validates the module properties.
func (ex *Example) Validate() error {
	var err error

	if ex.strProp == "bar" {
		err = multierr.Append(err, errors.New("str prop must be different than bar"))
	}

	if ex.intProp == 1337 {
		err = multierr.Append(err, errors.New("int prop must be different than 1337"))
	}

	return err
}

func (ex *Example) SystemMessages() []string {
	return []string{
		"Hello world!",
	}
}

// Interface guards.
var (
	_ gotenberg.Module       = (*Example)(nil)
	_ gotenberg.Provisioner  = (*Example)(nil)
	_ gotenberg.Validator    = (*Example)(nil)
	_ gotenberg.SystemLogger = (*Example)(nil)
)
