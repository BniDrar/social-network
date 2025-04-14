package notification

import (
	"database/sql"
	"errors"
	"fmt"
)

// need some changes to follow up with the macro image
func (n *Notification) GroupContainsMember(groupId, userId int) (bool, error) {
	query := `SELECT 1 FROM members WHERE group_id = ? AND user_id = ? LIMIT 1`
	var exists int

	err := n.db.QueryRow(query, groupId, userId).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, errors.New(fmt.Sprintf("Error checking group membership:", err))
	}

	return true, nil
}
