package main

import (
	"database/sql"
	"fmt"
	"miniproject/database"
	"miniproject/controllers"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var (
	DB *sql.DB
	err error
)

func main() {
	err = godotenv.Load(".env")
    if err != nil {
       panic("Error loading .env file")
    }

    psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
       os.Getenv("PGHOST"),
       os.Getenv("PGPORT"),
       os.Getenv("PGUSER"),
       os.Getenv("PGPASSWORD"),
       os.Getenv("PGDATABASE"),
    )


	DB, err = sql.Open("postgres", psqlInfo)
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	router := gin.Default()
    router.GET("/persons", controllers.GetAllPerson)
    router.POST("/persons", controllers.InsertPerson)
    router.PUT("/persons/:id", controllers.UpdatePerson)
    router.DELETE("/persons/:id", controllers.DeletePerson)

	router.Run(":" + os.Getenv("PORT"))

	fmt.Println("Berhasil menyambung ke DB!")
}