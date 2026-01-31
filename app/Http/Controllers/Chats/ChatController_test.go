package Chats

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStoreRejectsInvalidJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/chats", strings.NewReader(`{"title":`))
	rec := httptest.NewRecorder()

	Store(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestShowRejectsInvalidID(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/chats/abc", nil)
	req.SetPathValue("id", "abc")
	rec := httptest.NewRecorder()

	Show(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestDestroyRejectsMissingID(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/chats/", nil)
	rec := httptest.NewRecorder()

	Destroy(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestStoreMessageRejectsInvalidBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/chats/1/messages", strings.NewReader(`{"text":""}`))
	req.SetPathValue("id", "1")
	rec := httptest.NewRecorder()

	StoreMessage(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}
