package seeds

import (
	"log"
)

func (s Seed) CategorySeed() {
	category := []string{"SMA/SMK", "Mahasiswa"}

	for i := 0; i < 2; i++ {
		sqlStmt := `INSERT INTO categories (category_name) VALUES (?)`

		_, err := s.db.Exec(sqlStmt, category[i])
		if err != nil {
			log.Println(err)
		}
	}
}
