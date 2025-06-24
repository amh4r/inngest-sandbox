# Inngest Sandbox

Polyglot sandbox for Inngest. Useful for testing features and troubleshooting.

## Usage

Run Dev Server:
```sh
npx inngest-cli@latest dev -u localhost:3939/api/inngest
```

Start an app:
```sh
(cd go && make dev)
(cd py-fastapi && make dev)
(cd py-flask && make dev)
(cd ts-express && make dev)
```

## Tips

If you're using a VSCode IDE, you can use the `root.code-workspace` file to view each example as a workspace:
```sh
code root.code-workspace
```
