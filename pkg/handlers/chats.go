package handlers

import (
	"chatsapp/pkg/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (handler *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chat := &models.Chat{}
	err = json.Unmarshal(body, chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Проверка на уникальность идентификаторов пользователей
	uniqueMap := make(map[uint32]bool)
	for _, userID := range chat.Users {
		if !uniqueMap[userID] {
			uniqueMap[userID] = true
		} else {
			http.Error(w, "user id's aren't unique", http.StatusBadRequest)
			return
		}
	}

	id, err := handler.Repo.CreateChat(chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "\nThe chat with ID: %v was created\n", id)
}

func (handler *Handler) GetChats(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := &models.User{}
	err = json.Unmarshal(body, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chats, err := handler.Repo.GetUserChats(user)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "user not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("\nUser chats, user_id=%v\n", user.ID))

	for _, chat := range chats {
		builder.WriteString(fmt.Sprintf("ID: %v\n", chat.ID))
		builder.WriteString(fmt.Sprintf("Name: %s\n", chat.Name))
		builder.WriteString(fmt.Sprintf("Users: %v\n", chat.Users))
		builder.WriteString(fmt.Sprintf("Created at: %s\n\n", chat.CreatedAt))
	}
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintf(w, builder.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}