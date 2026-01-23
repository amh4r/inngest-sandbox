# Inngest Sandbox

Polyglot sandbox for Inngest. Useful for testing features and troubleshooting.

## Getting started

```sh
# Install Python dependencies
uv sync --all-packages

# Install TypeScript dependencies
pnpm install

# Create root env var file
cp example.env .env
```

## Usage

Run Dev Server:
```sh
npx inngest-cli@latest dev -u localhost:3939/api/inngest
```

Start an app:
```sh
# Go app
(cd go && make dev)

# Python apps
uv run --directory py-fastapi make dev
uv run --directory py-flask make dev
uv run --directory py-connect make dev

# TypeScript apps
pnpm -C ts-express run dev
pnpm -C ts-cloudflare-worker run dev
pnpm -C ts-connect run dev
pnpm -C ts-next run dev
```

Env vars in the root `.env` file reflect in all apps.

## Updating Inngest dependencies

```sh
go run ./cli update-inngest
```
