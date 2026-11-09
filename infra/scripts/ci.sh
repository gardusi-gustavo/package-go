#!/usr/bin/env bash
set -euo pipefail
ROOT="${CI_PROJECT_DIR:-$(git rev-parse --show-toplevel)}"
cd "$ROOT"
"$ROOT/infra/scripts/ci/validate-repo.sh" --root "$ROOT"
"$ROOT/infra/scripts/ci/check-size.sh" --root "$ROOT"
go test ./...
go vet ./...
(cd examples/http-app && go build ./...)
for f in infra/scripts/*.sh; do bash -n "$f"; done
echo "go checks ok"
