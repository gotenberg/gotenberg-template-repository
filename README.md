> Use this GitHub template repository for your custom modules and Gotenberg's Docker image variants.
> 
> Want to share your work? Open a PR to the [awesome list](https://github.com/gotenberg/awesome-gotenberg)! 🚀

## Quick Start

Update the following variables from the `Makefile`:

* `APP_NAME` - the name of your Gotenberg's Docker image variant.
* `APP_VERSION`
* `APP_AUTHOR`
* `APP_REPOSITORY`
* `DOCKER_REGISTRY`

Good 🤓? Now run:

```bash
make it
```

This command builds both your Gotenberg's Docker image variant (`$(DOCKER_REGISTRY)/gotenberg:7-$(APP_NAME)-$(APP_VERSION)`)
and a tests' Docker image.

## Next steps

1. Update the `go.mod` file with your Go module's name.
2. Update the `build/Docker` with your instructions.
3. Create your module(s) in `pkg/modules`.
4. Import your module(s) in `cmd/app/main.go`.

## Testing

First, run:

```bash
make tests
```

Once your in the testing container, you have access to the following commands:

```
golint   Run the linter.
gotest   Run the tests.
gotodos  Display TODOs in your Go source.
```
