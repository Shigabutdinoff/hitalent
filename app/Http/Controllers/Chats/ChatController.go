package Chats

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	chatRequest "hitalent/app/Http/Requests/Chats"
	chatModel "hitalent/app/Models/Chats"
	messageModel "hitalent/app/Models/Messages"

	"gorm.io/gorm"
)

func Store(w http.ResponseWriter, r *http.Request) {
	title, err := chatRequest.ParseCreate(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, chatRequest.ErrorResponse{Error: err.Error()})
		return
	}

	chat, err := Create(title)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, chatRequest.ErrorResponse{Error: "failed to create chat"})
		return
	}

	writeJSON(w, http.StatusOK, chat)
}

func Show(w http.ResponseWriter, r *http.Request) {
	id, err := parseChatID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, chatRequest.ErrorResponse{Error: err.Error()})
		return
	}

	limit, err := chatRequest.ParseLimit(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, chatRequest.ErrorResponse{Error: err.Error()})
		return
	}

	chat, messages, err := Get(id, limit)
	if err != nil {
		if errors.Is(err, chatRequest.ErrChatNotFound) {
			writeJSON(w, http.StatusNotFound, chatRequest.ErrorResponse{Error: "chat not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, chatRequest.ErrorResponse{Error: "failed to load chat"})
		return
	}

	writeJSON(w, http.StatusOK, chatRequest.ChatWithMessagesResponse{
		Chat:     chat,
		Messages: messages,
	})
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	id, err := parseChatID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, chatRequest.ErrorResponse{Error: err.Error()})
		return
	}

	deleted, err := Delete(id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, chatRequest.ErrorResponse{Error: "failed to delete chat"})
		return
	}
	if !deleted {
		writeJSON(w, http.StatusNotFound, chatRequest.ErrorResponse{Error: "chat not found"})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func StoreMessage(w http.ResponseWriter, r *http.Request) {
	chatId, err := chatRequest.ParseChatID(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, chatRequest.ErrorResponse{Error: err.Error()})
		return
	}

	text, err := chatRequest.ParseCreateMessage(r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, chatRequest.ErrorResponse{Error: err.Error()})
		return
	}

	message, err := createMessage(chatId, text)
	if err != nil {
		if errors.Is(err, chatRequest.ErrChatNotFound) {
			writeJSON(w, http.StatusNotFound, chatRequest.ErrorResponse{Error: "chat not found"})
			return
		}
		writeJSON(w, http.StatusInternalServerError, chatRequest.ErrorResponse{Error: "failed to create message"})
		return
	}

	writeJSON(w, http.StatusOK, message)
}

func parseChatID(r *http.Request) (int64, error) {
	rawID := r.PathValue("id")
	if rawID == "" {
		return 0, errors.New("chat id is required")
	}
	parsedID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil || parsedID < 1 {
		return 0, errors.New("chat id must be a positive integer")
	}
	return parsedID, nil
}

func Create(title string) (chatRequest.ChatResponse, error) {
	now := time.Now().UTC()
	chat := chatModel.GetModel()
	chat.Title = title
	chat.CreatedAt = &now

	if err := chat.Model.GetConnection().Create(&chat).Error; err != nil {
		return chatRequest.ChatResponse{}, err
	}

	return chatRequest.ChatResponse{
		Id:        chat.Id,
		Title:     chat.Title,
		CreatedAt: chat.CreatedAt,
	}, nil
}

func Get(id int64, limit int) (chatRequest.ChatResponse, []chatRequest.MessageResponse, error) {
	chat, err := chatModel.Find(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return chatRequest.ChatResponse{}, nil, chatRequest.ErrChatNotFound
		}
		return chatRequest.ChatResponse{}, nil, err
	}

	var messages []messageModel.Message
	if err := messageModel.GetModel().Model.GetConnection().
		Where("chat_id = ?", chat.Id).
		Order("created_at desc").
		Limit(limit).
		Find(&messages).Error; err != nil {
		return chatRequest.ChatResponse{}, nil, err
	}

	reverseMessages(messages)

	responseMessages := make([]chatRequest.MessageResponse, 0, len(messages))
	for _, message := range messages {
		responseMessages = append(responseMessages, chatRequest.MessageResponse{
			Id:        message.Id,
			ChatId:    message.ChatId,
			Text:      message.Text,
			CreatedAt: message.CreatedAt,
		})
	}

	return chatRequest.ChatResponse{
		Id:        chat.Id,
		Title:     chat.Title,
		CreatedAt: chat.CreatedAt,
	}, responseMessages, nil
}

func Delete(id int64) (bool, error) {
	result := chatModel.GetModel().Model.GetConnection().Delete(&chatModel.Chat{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil
}

func createMessage(chatId int64, text string) (chatRequest.MessageResponse, error) {
	chat := chatModel.GetModel()
	if err := chat.Model.GetConnection().First(&chat, chatId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return chatRequest.MessageResponse{}, chatRequest.ErrChatNotFound
		}
		return chatRequest.MessageResponse{}, err
	}

	now := time.Now().UTC()
	message := messageModel.GetModel()
	message.ChatId = chatId
	message.Text = text
	message.CreatedAt = &now

	if err := message.Model.GetConnection().Create(&message).Error; err != nil {
		return chatRequest.MessageResponse{}, err
	}

	return chatRequest.MessageResponse{
		Id:        message.Id,
		ChatId:    message.ChatId,
		Text:      message.Text,
		CreatedAt: message.CreatedAt,
	}, nil
}

func reverseMessages(messages []messageModel.Message) {
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
