# package-go

Go reference package: a small library, a full HTTP example, tests, and release
notes. It is the “how do I set up Go here?” repository.

## Layout

| Path | Holds |
| --- | --- |
| `package.go` | publishable library |
| `examples/http-app/` | full HTTP app using the library |
| `docs/` | usage, build, and release notes |

## Build

```bash
go test ./...
go vet ./...
(cd examples/http-app && go build ./...)
```
