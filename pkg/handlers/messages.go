package handlers

import (
	"chatsapp/pkg/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func (handler *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	message := &models.Message{}
	err = json.Unmarshal(body, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = handler.Repo.GetChatByID(message.Chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = handler.Repo.GetUserByID(message.Author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := handler.Repo.SendMessage(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "\nThe message with ID: %v was created\n", id)
}

func (handler *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	chat := &models.Chat{}
	err = json.Unmarshal(body, chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chat, err = handler.Repo.GetChatByID(chat.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	messages, err := handler.Repo.GetChatMessages(chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Messages of chat: %s\n", chat.Name))

	for _, message := range messages {
		builder.WriteString(fmt.Sprintf("ID: %v\n", message.ID))
		builder.WriteString(fmt.Sprintf("Chat: %v\n", message.Chat))
		builder.WriteString(fmt.Sprintf("Author: %v\n", message.Author))
		builder.WriteString(fmt.Sprintf("Text: %s\n", message.Text))
		builder.WriteString(fmt.Sprintf("Created at: %s\n\n", message.CreatedAt))
	}
	fmt.Fprintf(w, builder.String())
}