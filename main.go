package main

import (
	"fmt"
	"hitalent/app/Services/DatabaseManager"
)

func main() {
	connection, err := DatabaseManager.Connection("")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(connection.Dialector.Name())
}
