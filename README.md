# Product Management Full-Stack Project

A development project containing a Go REST API for managing products and a Next.js portal. The API stores product records in PostgreSQL and exposes create, read, update, and delete endpoints.

## Project structure

```text
apps/
  portal/        Next.js frontend
  test_fullstack/ Go API, PostgreSQL configuration, and tests
rest-test.rest   Example API requests (VS Code REST Client format)
```

## Requirements

- Go version specified by `apps/test_fullstack/go.mod` (Go 1.27.1)
- Node.js and npm
- Docker with Docker Compose

## Run locally

### 1. Configure the API and database

Create `apps/test_fullstack/.env` with local development settings:

```dotenv
PORT=8080
BLUEPRINT_DB_DATABASE=product_db
BLUEPRINT_DB_USERNAME=product_user
BLUEPRINT_DB_PASSWORD=product_password
BLUEPRINT_DB_HOST=localhost
BLUEPRINT_DB_PORT=5442
```

The `.env` file is ignored by Git. Compose uses the database settings above and maps PostgreSQL to port `5442` on your machine.

### 2. Start PostgreSQL

```bash
cd apps/test_fullstack
docker compose up -d
```

To stop the database when you are finished:

```bash
docker compose down
```

### 3. Start the Go API

From `apps/test_fullstack`:

```bash
go run ./cmd/api
```

The API listens at `http://localhost:8080`. It creates or updates the product table on startup. Alternatively, use the backend Makefile from this directory:

```bash
make run
```

### 4. Start the frontend

In another terminal:

```bash
cd apps/portal
npm ci
npm run dev
```

Open [http://localhost:3000](http://localhost:3000). The portal currently displays the default Next.js starter page. Its product API client is in `features/api.ts` and targets the local API at `http://localhost:8080`.

**CORS note:** The API currently allows browser requests from `http://localhost:5173`, not the Next.js default `http://localhost:3000`. The portal can start as described above, but browser-based API calls from it will need the API's allowed origin updated before they can succeed.

## API

All endpoints use JSON. Product records have this shape:

```json
{
  "id": 1,
  "name": "Example product",
  "price": 12.5,
  "stock": 4
}
```

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/` | Returns a hello message |
| `GET` | `/health` | Returns the API's health response |
| `POST` | `/products` | Creates a product |
| `GET` | `/products` | Lists products |
| `GET` | `/products/{id}` | Gets a product by ID |
| `PUT` | `/products/{id}` | Replaces a product's name, price, and stock |
| `DELETE` | `/products/{id}` | Deletes a product |

Create and update requests require `name`, `price`, and `stock`; price and stock must be greater than zero. Example:

```bash
curl -X POST http://localhost:8080/products \
  -H 'Content-Type: application/json' \
  -d '{"name":"Example product","price":12.5,"stock":4}'
```

For more request examples, see [`rest-test.rest`](./rest-test.rest).

## Build and test

Run these commands from `apps/test_fullstack`:

```bash
make build   # Build the API binary
make test    # Run all Go tests
make itest   # Run database integration tests
make clean   # Remove the built binary
```

The database integration tests use Testcontainers and require Docker to be running. For live reload, run `make watch` (it uses Air and offers to install it if it is not already available).

To check the frontend lint rules, run from `apps/portal`:

```bash
npm run lint
```
