package seeds

import (
	"log"

	"github.com/bxcodec/faker/v3"
	"go-scholarship/utils/hash"
)

func (s Seed) UserSeed() {
	for i := 0; i < 5; i++ {
		sqlStmt := `INSERT INTO users (name, email, password) VALUES (?, ?, ?)`

		// hash password
		password := "1234"
		hashed, _ := hash.HashPassword(password)

		_, err := s.db.Exec(sqlStmt, faker.Name(), faker.Email(), hashed)
		if err != nil {
			log.Println(err)
		}
	}
}
