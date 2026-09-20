package main

import (
	"fmt"

	"github.com/ArteShow/DoThat/internal/database"
)

func main() {
	db, err := database.Init("migrations", "data/dothat.db")
	if err != nil {
		panic(err)
	}

	fmt.Println(db.DB.Ping())
}
