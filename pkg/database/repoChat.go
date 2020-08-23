package database

import (
	"chatsapp/pkg/models"
	"database/sql"
)

func (repo *Repository) CreateChat(chat *models.Chat) (int64, error){
	// Проверка наличия в БД всех пользователей
	userIDs := make([]uint32, len(chat.Users))
	for i, userID := range chat.Users {
		_, err := repo.GetUserByID(userID)
		if err != nil {
			return -1, err
		}
		userIDs[i] = userID
	}
	result, err := repo.DB.Exec(
		"INSERT INTO chats (`name`) " +
			"VALUES (?)", chat.Name)
	if err != nil {
		return -1, err
	}
	chatID, err := result.LastInsertId()

	// Добавление пользователей в отношение c чатом
	for i := range chat.Users {
		_, err := repo.DB.Exec(
			"INSERT INTO user_chat_relation (`chat_id`, `user_id`) " +
			"VALUES (?, ?)", chatID, userIDs[i])
		if err != nil {
			return -1, err
		}
	}
	return chatID, nil
}

func (repo *Repository) GetChatByID(id uint32) (*models.Chat, error) {
	chat := &models.Chat{}
	err := repo.DB.
		QueryRow(
			"SELECT id, name, created_at FROM chats " +
			"WHERE id = ?", id).
		Scan(&chat.ID, &chat.Name, &chat.CreatedAt)
	if err != nil {
		return nil, err
	}
	// Находим список пользователей в чате
	users, err := repo.GetUsersInChat(chat)
	if err != nil {
		return nil, err
	}
	chat.Users = users
	return chat, nil
}
// Получение списка чатов пользователя
func (repo *Repository) GetUserChats(user *models.User) ([]*models.Chat, error) {
	rows, err := repo.DB.Query(
		"SELECT chats.id, chats.name, chats.created_at, last_time FROM " +
			"(chats JOIN (SELECT chat, max(created_at) as last_time FROM messages GROUP BY chat) as t1 " +
			"ON chats.id=t1.chat) " +
			"JOIN user_chat_relation ON chats.id=user_chat_relation.chat_id " +
			"WHERE user_chat_relation.user_id = ? " +
			"ORDER BY t1.last_time DESC",
		user.ID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chats := []*models.Chat{}
	for rows.Next() {
		chat := &models.Chat{}
		lastMessageTime := ""
		err = rows.Scan(&chat.ID, &chat.Name, &chat.CreatedAt, &lastMessageTime)
		if err != nil {
			return nil, err
		}

		users, err := repo.GetUsersInChat(chat)
		if err != nil {
			return nil, err
		}
		chat.Users = users
		chats = append(chats, chat)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return chats, err
}
// Получение списка идентификаторов пользователей чата
func (repo *Repository) GetUsersInChat(chat *models.Chat) ([]uint32, error) {
	rows, err := repo.DB.Query(
		"SELECT user_id FROM user_chat_relation " +
		"WHERE chat_id = ?", chat.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []uint32
	for rows.Next() {
		user := 0
		err = rows.Scan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, uint32(user))
	}
	if err = rows.Err(); err != nil {
		if err == sql.ErrNoRows {
			return users, nil
		}
		return nil, err
	}
	return users, nil
}