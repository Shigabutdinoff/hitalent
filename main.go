package main

import (
	"fmt"
	"hitalent/app/Services/DatabaseManager"
)

func main() {
	fmt.Println(DatabaseManager.Connection(""))
}
