# Account Management

Mange accounts with REST api.



### Executing program
In the root directory of the project:
```
go build -o account-management && ./account-management 
```
### APIs
Get all the accounts:
```
GET localhost:8080/accounts
```
Get one account by ID:
```
GET localhost:8080/account/{account_id}
```
Transfer from an account to other:
```
POST localhost:8080/transfer
{
    "from": "db6d4e22-ed9a-4bac-8b50-859d17e83591",
    "from": "d4679bee-1294-4133-a2c3-572416599f0d",
    "amount": 125.5
}

```
