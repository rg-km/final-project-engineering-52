# **Overview**
### Framework
- [Gin](https://gin-gonic.com/)
### Modules
- [Sqlite3](https://github.com/mattn/go-sqlite3) - database driver
- [Viper](https://github.com/spf13/viper) - configuration
- [golang-migrate](https://github.com/golang-migrate/migrate) - migration
- [Faker](https://github.com/bxcodec/faker) - generate data
- [JWT](https://github.com/golang-jwt/jwt) - authentication

<br>

# **Available API**
## **Authentication**
Login : `http://localhost:8080/login` <br>
Register : `http://localhost:8080/register`

## **Users**
Fetch all : `http://localhost:8080/api/users` <br>
Create : `http://localhost:8080/api/users` <br>
Get by id : `http://localhost:8080/api/users/:id` <br>
Update : `http://localhost:8080/api/users/:id` <br>
Delete : `http://localhost:8080/api/users/:id` <br>

## **Scholarships**
Fetch all : `http://localhost:8080/api/scholarships` <br>
Create : `http://localhost:8080/api/scholarships` <br>
Get by id : `http://localhost:8080/api/scholarships/:id` <br>
Update : `http://localhost:8080/api/scholarships/:id` <br>
Delete : `http://localhost:8080/api/scholarships/:id` <br>

## **Categories**
Fetch all : `http://localhost:8080/api/categories` <br>
Create : `http://localhost:8080/api/categories` <br>
Get by id : `http://localhost:8080/api/categories/:id` <br>
Update : `http://localhost:8080/api/categories/:id` <br>
Delete : `http://localhost:8080/api/categories/:id` <br>

## **Comments**
Fetch all : `http://localhost:8080/api/comments` <br>
Create : `http://localhost:8080/api/comments` <br>
Get by id : `http://localhost:8080/api/comments/:id` <br>
Update : `http://localhost:8080/api/comments/:id` <br>
Delete : `http://localhost:8080/api/comments/:id` <br>

<br>

# **Install golang-migration**
```bash
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

<br>

# **Create schema (optional)**
```bash
migrate create -ext sql -dir db/migration -seq init_schema
```

<br>

# **Run the migration table using sqlite3 driver**
```bash
migrate -path db/migration -database "sqlite3://db/go_scholarship.db?sslmode=disable" -verbose up
```

<br>

# **Drop table (optional)**
```bash
migrate -path db/migration -database "sqlite3://db/go_scholarship.db?sslmode=disable" -verbose down
```

<br>

# **Seeding**
Run all seeder, example :
```bash
go run main.go seed
```

Run specific seeder, example :
```bash
go run main.go seed UserSeed
```

Run multiple seeder, example :
```bash
go run main.go seed UserSeed CategorySeed
```
