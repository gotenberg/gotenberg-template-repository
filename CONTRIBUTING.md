# Contributing to the template

A GitHub template repository for building custom Gotenberg variants. Keep it small and easy to fork.

## Toolchain

- Module: `github.com/gotenberg/gotenberg-template-repository`
- Go: 1.27.3
- Docker
- Node.js (see `.node-version`), for Prettier linting
- golangci-lint: the Makefile runs golangci-lint at a pinned version through `go run`, never the copy on your `PATH`.

## Project layout

```
cmd/app/             -> Entry point only (wiring/startup). Imports the modules.
pkg/modules/         -> Your modules. Ships with an example module.
build/               -> Dockerfile of the custom Gotenberg image.
```

## Setup and Makefile

| Command              | Purpose                                    | When to use           |
| -------------------- | ------------------------------------------ | --------------------- |
| `make help`          | Show the available commands                | To discover targets   |
| `make build`         | Build the custom Gotenberg Docker image    | Before manual testing |
| `make fmt`           | Format Go code and tidy the dependencies   | Before committing     |
| `make prettify`      | Format non-Go files (Markdown, YAML, JSON) | Before committing     |
| `make lint`          | Lint Go code (zero warnings permitted)     | Before committing     |
| `make lint-prettier` | Lint non-Go files                          | Before committing     |
| `make lint-todo`     | Find TODOs in Go code                      | Before committing     |

## Documentation conventions

### Tone

- Short, declarative sentences. Say what it does, then stop.
- Lead with the action. "Validates the input", not "This function validates the input".
- Active voice. "the template checks the input", not "The input is checked by the template".
- No em dashes. Use a period, colon, or comma.
- No "we" hedging. "Don't...", not "We do not recommend...".

### Godoc

Every exported identifier has a doc comment starting with its name, and every package a `doc.go` with a `// Package foo ...` comment. Reference identifiers with `[Name]` for pkg.go.dev links. Test, fuzz, benchmark and example functions are exempt.

### Code comments

- Explain _why_, not _what_.
- No numbered steps, no section separators, no comments that restate the code.
- Mark debt with `// TODO: [context]`.

## Pull requests

### Commits

[Conventional Commits](https://www.conventionalcommits.org/), title only: `<type>(<scope>): <description>`. No body, no trailers (no `Co-Authored-By`, no `Signed-off-by`).

Common types: `feat`, `fix`, `refactor`, `test`, `docs`, `chore`, `ci`, `build`. The scope names the area of the change.

Stage specific files. Never `git add -A` or `git add .`.

### Checklist

- `make fmt && make lint && make prettify && make lint-prettier` pass.
- `go test ./...` passes.
- `make build` succeeds when the Dockerfile or modules change.
- One thing per PR.
