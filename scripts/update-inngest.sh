#!/usr/bin/env bash
set -euo pipefail

# Update Inngest dependencies across all example projects
# Usage: ./scripts/update-inngest.sh

cd "$(dirname "$0")/.."

echo "==> Updating Go"
(cd go && go get -u github.com/inngest/inngestgo@latest && go mod tidy)

echo "==> Updating Python"
py_version=$(curl -s https://pypi.org/pypi/inngest/json | python3 -c "import sys, json; print(json.load(sys.stdin)['info']['version'])")
echo "    Latest version: $py_version"
for dir in py-*/; do
  sed -i '' "s/\"inngest==[^\"]*\"/\"inngest==$py_version\"/" "$dir/pyproject.toml"
  sed -i '' "s/\"inngest\[connect\]==[^\"]*\"/\"inngest[connect]==$py_version\"/" "$dir/pyproject.toml"
done
uv sync

echo "==> Updating TypeScript"
pnpm update 'inngest' '@inngest/*' --recursive --latest

echo "==> Done"
