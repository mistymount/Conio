package conio

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/hako/branca"
	"github.com/mistymount/Conio/internal/service"
)

const (
	databaseURL = "postgresql://root@127.0.0.1:26257/conio?sslmode=disable"
	port        = 3000
)

func main() {
	fmt.Println("Conio - Your personal secure social network")
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
		return
	}

	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v\n", err)
		return
	}

	codec := branca.NewBranca("supersecretkeywhichissecureandencrypted")
	service.New(db, codec)
}
