package conio

import "fmt"

const (
	databaseURL = "postgresql://root@127.0.0.1:26257/conio?sslmode=disable"
	port        = 3000
)

func main() {
	fmt.Println("Conio - Your personal secure social network")
}
