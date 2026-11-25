# Account Management API

A small REST API in Go for viewing accounts and transferring money between them. It loads accounts from `accounts.json` into memory at startup and serves them over HTTP.

## Design

The code is split into layers, each with one job:

| Layer | Path | Responsibility |
|---|---|---|
| Handlers | `handler/` | HTTP routing (gorilla/mux), request parsing and JSON responses |
| Service | `service/` | Transfer rules: both accounts must exist, and the sender must have enough funds |
| Repository | `repository/` | In-memory account store, guarded by a read/write mutex for concurrent requests |
| Domain | `domain/` | `Account` and `Transfer` types |

## Run

Requires Go.

```bash
go build -o account-management && ./account-management
```

The server listens on port 8080.

## API

Get all accounts:

```
GET localhost:8080/accounts
```

Get one account by ID:

```
GET localhost:8080/account/{account_id}
```

Transfer money between accounts:

```
POST localhost:8080/transfer
{
    "from": "db6d4e22-ed9a-4bac-8b50-859d17e83591",
    "to": "d4679bee-1294-4133-a2c3-572416599f0d",
    "amount": 125.5
}
```

A transfer fails if either account ID is unknown or the sender's balance is lower than the amount.

## Tests

```bash
go test ./...
```

The handler tests cover a successful transfer, an unknown sender and insufficient funds.
