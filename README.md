> GitHub template repository for building custom modules and Gotenberg Docker image variants.
> 
> Built something cool? Open a PR to the [awesome list](https://github.com/gotenberg/awesome-gotenberg)! 🚀

## Quick Start

Set these variables in the `Makefile`:

* `APP_NAME`: your Gotenberg Docker image variant name.
* `APP_VERSION`
* `APP_AUTHOR`
* `APP_REPOSITORY`
* `DOCKER_REGISTRY`
* `DOCKER_REPOSITORY`

Then build:

```bash
make build
```

Produces a Docker image tagged `$(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY):$(GOTENBERG_VERSION)-$(APP_NAME)-$(APP_VERSION)`.

## Next Steps

1. Rename the Go module in `go.mod`.
2. Add your instructions to `build/Dockerfile`.
3. Create your module(s) in `pkg/modules`.
4. Import your module(s) in `cmd/app/main.go`.

## FAQ

> How do I check the underlying Gotenberg version?

Inspect the `version` label on the image:

```
docker inspect $(DOCKER_REGISTRY)/$(DOCKER_REPOSITORY):$(GOTENBERG_VERSION)-$(APP_NAME)-$(APP_VERSION)
```

---

> What commands are available?

Run `make help` to list them 💡
