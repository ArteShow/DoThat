package main

import (
	"fmt"
	"os"

	"github.com/ArteShow/DoThat/internal/database"
)

func main() {
	if err := os.Mkdir("data", 0755); err != nil {
		panic(err)
	}

	db, err := database.Init("migrations", "data/dothat.db")
	if err != nil {
		panic(err)
	}

	fmt.Println(db.DB.Ping())
}

//unnesesery
