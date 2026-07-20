#!/usr/bin/env bash
set -euo pipefail

OUTPUT="${1:-bin/anime-list-server}"
cd "$(dirname "$0")/.."

go build -o "$OUTPUT" ./cmd/api/
echo "Build succeeded: $OUTPUT"
