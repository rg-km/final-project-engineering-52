package seeds

import (
	"log"
)

func (s Seed) CommentSeed() {
	comment := []string{"Beasiswa nya bagus", "Beasiswa terpercaya", "Alhamdulillah lolos beasiswa ini", "Cukup bagus beasiswanya", "Sangat bagus"}

	for i := 0; i < 5; i++ {
		sqlStmt := `INSERT INTO comments (content, user_id, scholarship_id) VALUES (?, ?, ?)`

		_, err := s.db.Exec(sqlStmt, comment[i], i, i+1)
		if err != nil {
			log.Println(err)
		}
	}
}
