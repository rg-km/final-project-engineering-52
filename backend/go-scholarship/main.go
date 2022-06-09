package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"go-scholarship/api/handlers"
	"go-scholarship/api/repository"
	"go-scholarship/db/seeds"
	"github.com/spf13/viper"
)

func main() {
	// viper config
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	driver := viper.GetString("db.driver")
	dsn := viper.GetString("db.dsn")
	port := viper.GetString("db.port")
	debug := viper.GetString("debug")

	// debug check
	if debug == "true" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// database
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Println("Error :", err)
	}

	defer db.Close()

	// seeding
	arguments(db)

	r := gin.Default()

	// users
	u := repository.NewUserRepo(db)
	handlers.NewUserHandler(r, u)

	// scholarships
	s := repository.NewScholarshipRepository(db)
	handlers.NewScholarshipHandler(r, s)

	// comments
	co := repository.NewCommentRepository(db)
	handlers.NewCommentHandler(r, co)

	// categories
	ca := repository.NewCategoryRepository(db)
	handlers.NewCategoryHandler(r, ca)

	// start server
	r.Run(":" + port)
}

func arguments(db *sql.DB) {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			seeds.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
}
