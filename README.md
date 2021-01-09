# Golang banking app

This is app for the Duomly Golang Course - Learn Golang by building banking App

## Run

- Setup a PostgreSQL db \* Troubleshoot section for more details
- You need to setup your connection string in the two files, vulnerable-db.go, and migrations.go
- Start migration by commenting migration in main.go and commenting API

```
func main() {
	// Init database
	database.InitDatabase()
	// Do migration
	migrations.Migrate()
	// api.StartApi()
}
```

- Run (this migrate to your db):

```
go run main.go
```

- Comment out migration and uncomment api (to start server)

```
func main() {
	// Init database
	database.InitDatabase()
	// Do migration
	// migrations.Migrate()
	api.StartApi()
}
```

- Run:

```
go run main.go
```

This should start your server on localhost:8888

### API Examples

To register a new user:

Use this json as the body of your payload in a POST request to localhost:8888/register

```
{
	"username": "xxxxxx",
	"email": "xxxxxx@gmail.com",
	"password": "some_password"
}
```

Example response from /register:

```
{
  "data": {
    "ID": 5,
    "Username": "xxxxxx",
    "Email": "xxxxxx@gmail.com",
    "Accounts": [
      {
        "ID": 5,
        "Name": "xxxxxx's account",
        "Balance": 0
      }
    ]
  },
  "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcnkiOjE2MTAxODQ1ODgsInVzZXJfaWQiOjV9.E8FXVYmz_N_IL0WnC9COO2XyrHc7HFPjmQUaeg3MKjk",
  "message": "all is fine"
}
```

To send a transaction:
Use this json as the body of your payload in a POST request to localhost:8888/transaction

Make sure that you add your "jwt" token you obtain from /login or /register response and add that as your request's Authorization Bearer token.

```
{
	"UserId" : 1,
	"From": 1,
	"To": 3,
	"Amount" : 1000
}
```

### Troubleshoot (Mac)

When you're setting up your postgres db for the first time on a mac you can run:

```
createuser -s postgres - fixes role "postgres" does not exist
createdb bankapp - fixes database "bankapp" does not exist
```
