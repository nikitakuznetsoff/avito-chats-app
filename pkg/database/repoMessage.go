package database

import (
	"chatsapp/pkg/models"
)

func (repo *Repository) SendMessage(message *models.Message) (int64, error) {
	result, err := repo.DB.Exec("INSERT INTO messages (`chat`, `author`, `text`) VALUES (?, ?, ?)",
		message.Chat,
		message.Author,
		message.Text,
	)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

func (repo *Repository) GetChatMessages(chat *models.Chat) ([]*models.Message, error) {
	rows, err := repo.DB.Query(
		"SELECT id, chat, author, text, created_at " +
		"FROM messages WHERE chat = ? " +
		"ORDER BY messages.created_at DESC", chat.ID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	messages := []*models.Message{}
	for rows.Next() {
		message := &models.Message{}
		err = rows.Scan(&message.ID, &message.Chat,
			&message.Author, &message.Text, &message.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}