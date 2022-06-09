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
