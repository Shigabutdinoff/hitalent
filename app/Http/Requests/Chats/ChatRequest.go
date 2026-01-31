package Chats

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type ChatResponse struct {
	Id        int64      `json:"id"`
	Title     string     `json:"title"`
	CreatedAt *time.Time `json:"created_at"`
}

type MessageResponse struct {
	Id        int64      `json:"id"`
	ChatId    int64      `json:"chat_id"`
	Text      string     `json:"text"`
	CreatedAt *time.Time `json:"created_at"`
}

type ChatWithMessagesResponse struct {
	Chat     ChatResponse      `json:"chat"`
	Messages []MessageResponse `json:"messages"`
}

var ErrChatNotFound = errors.New("chat not found")

type ErrorResponse struct {
	Error string `json:"error"`
}

type createChatRequest struct {
	Title string `json:"title"`
}

type createMessageRequest struct {
	Text string `json:"text"`
}

func ParseCreate(r *http.Request) (string, error) {
	var req createChatRequest
	if err := decodeJSONBody(r, &req); err != nil {
		return "", errors.New("invalid request body")
	}
	return normalizeTitle(req.Title)
}

func ParseCreateMessage(r *http.Request) (string, error) {
	var req createMessageRequest
	if err := decodeJSONBody(r, &req); err != nil {
		return "", errors.New("invalid request body")
	}
	return normalizeText(req.Text)
}

func ParseLimit(r *http.Request) (int, error) {
	limit := 20
	rawLimit := r.URL.Query().Get("limit")
	if rawLimit == "" {
		return limit, nil
	}
	parsed, err := strconv.Atoi(rawLimit)
	if err != nil || parsed < 1 || parsed > 100 {
		return 0, errors.New("limit must be between 1 and 100")
	}
	return parsed, nil
}

func normalizeTitle(title string) (string, error) {
	trimmed := strings.TrimSpace(title)
	if trimmed == "" {
		return "", errors.New("title is required")
	}
	length := utf8.RuneCountInString(trimmed)
	if length < 1 || length > 200 {
		return "", errors.New("title length must be between 1 and 200")
	}
	return trimmed, nil
}

func ParseChatID(r *http.Request) (int64, error) {
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

func normalizeText(text string) (string, error) {
	trimmed := strings.TrimSpace(text)
	if trimmed == "" {
		return "", errors.New("text is required")
	}
	length := utf8.RuneCountInString(trimmed)
	if length < 1 || length > 5000 {
		return "", errors.New("text length must be between 1 and 5000")
	}
	return trimmed, nil
}

func decodeJSONBody(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(dst); err != nil {
		return err
	}
	if decoder.Decode(&struct{}{}) != io.EOF {
		return errors.New("request body must be a single JSON object")
	}
	return nil
}
