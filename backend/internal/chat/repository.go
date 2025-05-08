package chat

import (
	"context"
	"fmt"
	"socialNetwork/entity"
	"strings"
)

func (c *chat) getMessagesRepo(ctx context.Context, cursor entity.Cursor, userId int) ([]entity.Message, error) {
	var query string
	var params []interface{}
	var whereClauses []string

	if cursor.IsGroup {
		query = `SELECT * FROM messages AS m WHERE `
		whereClauses = append(whereClauses, "(group_id = ?)")
		params = append(params, cursor.ID)
	} else {
		query = `SELECT * FROM messages WHERE `
		whereClauses = append(whereClauses, "((sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?))")
		params = append(params, userId, cursor.ID, cursor.ID, userId)
	}

	// Pagination clause (only if cursor provided)
	if cursor.Time != nil && cursor.LastId != nil {
		whereClauses = append(whereClauses, "(created_at < ? OR (created_at = ? AND id < ?))")
		params = append(params, cursor.Time, cursor.Time, cursor.LastId)
	}

	// Final query assembly
	query += strings.Join(whereClauses, " AND ") + " ORDER BY created_at DESC, id DESC LIMIT ?"
	params = append(params, cursor.Limit)

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []entity.Message
	for rows.Next() {
		var message entity.Message
		err := rows.Scan(&message.ID, &message.SenderID, &message.ReceiverID, &message.GroupID, &message.CreatedAt, &message.Content)
		if err != nil {
			c.loger.Error.Println(err)
			continue
		}
		messages = append(messages, message)
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

func (r *chat) SaveMessage(ctx context.Context, message entity.Message) error {
	query := `INSERT INTO messages (sender_id, receiver_id,created_at ,content) VALUES (?, ?, ?, ?)`
	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, message.SenderID, message.ReceiverID, message.CreatedAt, message.Content)
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

func (c *chat) getOnlines(ctx context.Context, onlineIDs []uint, excludeID int) ([]entity.Contact, error) {
	// Filter out the excludeID from onlineIDs
	filtered := make([]uint, 0, len(onlineIDs))
	for _, id := range onlineIDs {
		if id != (uint)(excludeID) {
			filtered = append(filtered, id)
		}
	}

	if len(filtered) == 0 {
		return []entity.Contact{}, nil
	}

	query := `SELECT id, nickname, first_name, last_name, avatar FROM users WHERE id IN (?` + strings.Repeat(",?", len(filtered)-1) + `)`

	args := make([]interface{}, len(filtered))
	for i, id := range filtered {
		args[i] = id
	}

	stmt, err := c.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []entity.Contact
	for rows.Next() {
		var contact entity.Contact
		if err := rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Avatar); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}
