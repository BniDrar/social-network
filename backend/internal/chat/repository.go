package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"socialNetwork/entity"
)

func (r *chat) GetChats(id int) ([]entity.Chat, error) {
	query := `SELECT * FROM chats WHERE user_id = $1 OR friend_id = $1`
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var chats []entity.Chat
	for rows.Next() {
		var chat entity.Chat
		if err := rows.Scan(&chat.ID, &chat.UserID, &chat.FriendID, &chat.Message, &chat.CreatedAt); err != nil {
			return nil, err
		}
		chats = append(chats, chat)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return chats, nil
}

func (r *chat) GetChatById(id int) (entity.Chat, error) {
	query := `SELECT * FROM chats WHERE id = $1`
	row := r.db.QueryRow(query, id)
	var chat entity.Chat
	if err := row.Scan(&chat.ID, &chat.UserID, &chat.FriendID, &chat.Message, &chat.CreatedAt); err != nil {
		return chat, err
	}
	return chat, nil
}

func (r *chat) CreateChat(chat entity.Chat) (int, error) {
	query := `INSERT INTO chats (user_id, friend_id, message) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(query, chat.UserID, chat.FriendID, chat.Message).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *chat) GetChatMessages(chatID int) ([]entity.Message, error) {
	query := `SELECT * FROM chat_messages WHERE chat_id = $1`
	rows, err := r.db.Query(query, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var messages []entity.Message
	for rows.Next() {
		var message entity.Message
		if err := rows.Scan(&message.ID, &message.GroupID, &message.SenderID, &message.Content, &message.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *chat) CreateMessage(message entity.Message) (int, error) {
	query := `INSERT INTO chat_messages (chat_id, sender_id, content) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := r.db.QueryRow(query, message.GroupID, message.SenderID, message.Content).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *chat) SaveMessage(message entity.Message) error {
	query := `INSERT INTO messages (sender_id, receiver_id,created_at ,content) VALUES (?, ?, ?, ?)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(message.SenderID, message.ReceiverID, message.CreatedAt, message.Content)
	if err != nil {
		return err
	}
	return nil
}

/*
type Contact struct {
	ID int `json:"id"`
	GrouptName 'json:"group_name"
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Avatar NullString `json:"avatar"`
	Online bool
}*/

func (c *chat) getContacts(ctx context.Context, userId int) ([]entity.Contact, error) {
	query := `
		SELECT DISTINCT 
			u.id,
			NULL AS group_name,
			u.first_name,
			u.last_name,
			u.avatar
		FROM users u
		JOIN messages m ON (u.id = m.sender_id AND m.receiver_id = ?) OR (u.id = m.receiver_id AND m.sender_id = ?)
		WHERE u.id != ?

		UNION

		SELECT DISTINCT
			g.id,
			g.name AS group_name,
			NULL AS first_name,
			NULL AS last_name,
			NULL AS avatar
		FROM groups g
		JOIN group_members gm ON gm.group_id = g.id
		WHERE gm.member_id = ? AND g.type IN (0, 2)
	`

	rows, err := c.db.QueryContext(ctx, query, userId, userId, userId, userId)
	if err != nil {
		return nil, fmt.Errorf("query contacts: %w", err)
	}
	defer rows.Close()

	var contacts []entity.Contact
	for rows.Next() {
		var contact entity.Contact
		var avatar entity.NullString
		var groupName entity.NullString
		var firstName entity.NullString
		var lastName entity.NullString

		if err := rows.Scan(
			&contact.ID,
			&groupName,
			&firstName,
			&lastName,
			&avatar,
		); err != nil {
			return nil, fmt.Errorf("scan contact: %w", err)
		}

		contact.GroupName = groupName.String
		contact.FirstName = firstName.String
		contact.LastName = lastName.String
		contact.Avatar = avatar.String

		contacts = append(contacts, contact)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return contacts, nil
}

func (c *chat) getOnlines(ctx context.Context, userId int) ([]entity.Contact, error) {
	query := `
		SELECT DISTINCT 
			u.id,
			NULL AS group_name,
			u.first_name,
			u.last_name,
			u.avatar
		FROM users u
		JOIN messages m ON (u.id = m.sender_id AND m.receiver_id = ?) OR (u.id = m.receiver_id AND m.sender_id = ?)
		WHERE u.id != ?

		UNION

		SELECT DISTINCT
			g.id,
			g.name AS group_name,
			NULL AS first_name,
			NULL AS last_name,
			NULL AS avatar
		FROM groups g
		JOIN group_members gm ON gm.group_id = g.id
		WHERE gm.member_id = ? AND g.type IN (0, 2)
	`

	rows, err := c.db.QueryContext(ctx, query, userId, userId, userId, userId)
	if err != nil {
		return nil, fmt.Errorf("query contacts: %w", err)
	}
	defer rows.Close()

	var contacts []entity.Contact
	for rows.Next() {
		var contact entity.Contact
		var avatar entity.NullString
		var groupName entity.NullString
		var firstName entity.NullString
		var lastName entity.NullString

		if err := rows.Scan(
			&contact.ID,
			&groupName,
			&firstName,
			&lastName,
			&avatar,
		); err != nil {
			return nil, fmt.Errorf("scan contact: %w", err)
		}

		contact.GroupName = groupName.String
		contact.FirstName = firstName.String
		contact.LastName = lastName.String
		contact.Avatar = avatar.String

		contacts = append(contacts, contact)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return contacts, nil
}
func DecodeMessage(data []byte, msg *entity.Message) (*entity.Message, error) {
	if err := json.Unmarshal(data, msg); err != nil {
		return msg, err
	}
	return msg, nil
}
