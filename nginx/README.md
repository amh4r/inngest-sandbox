# Nginx

This is a simple nginx server that proxies requests to an example app.

## Usage

Set the following in the `.env` file:

```sh
INNGEST_SERVE_HOST=http://localhost:3980
INNGEST_SERVE_ORIGIN=http://localhost:3980
```

Start the server with:

```sh
make start
```

You can now access the example app via the proxy:

```sh
curl http://localhost:3980/api/inngest
```
