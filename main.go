package main

import (
	"fmt"

	"hitalent/app/Models/Chat"
)

func main() {
	chat, err := Chat.Find(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(chat)
}
