package routes

import (
	"net/http"

	ChatController "hitalent/app/Http/Controllers/Chats"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /chats/{id}", ChatController.Show)
	mux.HandleFunc("POST /chats", ChatController.Store)
	mux.HandleFunc("POST /chats/{id}/messages", ChatController.StoreMessage)
	mux.HandleFunc("DELETE /chats/{id}", ChatController.Destroy)
}
