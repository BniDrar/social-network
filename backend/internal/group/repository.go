package group

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"socialNetwork/entity"
)

func (g *group) GetUserByIdRepository(ctx context.Context,  userID int) (entity.User, error) {
	query := `
		SELECT id, first_name
		FROM users
		WHERE id = ?
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return entity.User{}, err
	}
	defer stmt.Close()
	var user entity.User
	err = stmt.QueryRowContext(ctx, userID).Scan(
		&user.ID,
		&user.First,
	)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (g *group) GetGroupsByUserID(ctx context.Context, userID int, limit int, offset int) (entity.Groups, error) {
	query := `
	SELECT 
    g.id, 
    g.name, 
    g.type,
		g.description,
    g.admin,
    COUNT(DISTINCT gm.member_id) AS member_count, 
    COUNT(DISTINCT p.id) AS post_count
	FROM groups g
	JOIN group_members gm ON g.id = gm.group_id
	LEFT JOIN posts p ON g.id = p.group_id
	WHERE gm.member_id = ? OR g.admin = ?
	GROUP BY g.id
	ORDER BY g.id DESC
  LIMIT ? OFFSET ?
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error preparing query in GetGroupsByUserID function: %v", err)
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, userID, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var groups entity.Groups
	for rows.Next() {
		var group entity.Group
		err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.Type,
			&group.Description,
			&group.Admin,
			&group.MemberCount,
			&group.PostCount,
		)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return groups, nil
}

func (g *group) GetGroupByIdRepository(ctx context.Context, userID, groupID int) (entity.Group, error) {
	query := `
		SELECT 
			g.id, 
			g.name, 
			g.description, 
			g.type, 
			g.admin, 
			COUNT(DISTINCT gm.member_id) AS member_count, 
			COUNT(DISTINCT p.id) AS post_count
		FROM groups g
		LEFT JOIN group_members gm ON g.id = gm.group_id
		LEFT JOIN posts p ON g.id = p.group_id
		WHERE g.id = ?
		GROUP BY g.id
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return entity.Group{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, groupID)
	var group entity.Group
	err = row.Scan(
		&group.ID,
		&group.Name,
		&group.Description,
		&group.Type,
		&group.Admin,
		&group.MemberCount,
		&group.PostCount,
	)
	if err != nil {
		return entity.Group{}, err
	}
	return group, nil
}

func (g *group) CreateGroupRepository(ctx context.Context, group entity.Group) (entity.Group, error) {
	query := `
		INSERT INTO groups (name,description, type, admin)
		VALUES (?, ?, ?, ?)
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return entity.Group{}, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, group.Name, group.Description, group.Type, group.Admin)
	if err != nil {
		return entity.Group{}, err
	}
	groupID, err := result.LastInsertId()
	if err != nil {
		return entity.Group{}, err
	}
	group.ID = int(groupID)
	return group, nil
}

func (g *group) GetAllGroupsRepository(ctx context.Context, typeGroup int) (entity.Groups, error) {
	userId := ctx.Value(entity.ContextID).(int)
	query := `
		SELECT
			g.id,
			g.name,
			g.description,
			g.type,
			g.admin,
			(gm.member_id IS NOT NULL OR g.admin = ?) AS isMember
		FROM groups g
		LEFT JOIN group_members gm ON gm.group_id = g.id AND gm.member_id = ?
		WHERE g.type = ?
		ORDER BY g.id DESC`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, userId, userId, typeGroup)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var groups entity.Groups
	for rows.Next() {
		var group entity.Group
		err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.Description,
			&group.Type,
			&group.Admin,
			&group.IsMember,
		)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return groups, nil
}

func (g *group) GetGroupMembersRepository(ctx context.Context, groupID int) ([]entity.User, error) {
	query := `
		SELECT 
			u.id,
			u.nickname,
			u.avatar,
			(g.admin = u.id) as is_admin
		FROM (
			SELECT member_id as id FROM group_members WHERE group_id = ?
			UNION
			SELECT admin as id FROM groups WHERE id = ?
		) as combined_members
		JOIN users u ON u.id = combined_members.id
		LEFT JOIN groups g ON g.id = ?
		WHERE u.id IS NOT NULL
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, groupID, groupID, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var groupMembers []entity.User
	for rows.Next() {
		var groupMember entity.User
		var isAdmin bool
		err := rows.Scan(
			&groupMember.ID,
			&groupMember.Nickname,
			&groupMember.Avatar,
			&isAdmin,
		)
		if err != nil {
			return nil, err
		}
		groupMember.IsAdmin = isAdmin
		groupMembers = append(groupMembers, groupMember)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return groupMembers, nil
}

func (g *group) addGroupMember(ctx context.Context, userID, groupID int) (int, error) {
	query := `INSERT INTO group_members (member_id, group_id)
	VALUES (?, ?)`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, userID, groupID)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

// CHECK IF THE USER IS A MEMBER OF THE GROUP
func (g *group) IsMemberRepository(ctx context.Context, userID, groupID int) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM group_members
			WHERE member_id = ? AND group_id = ?
			
			UNION
			
			SELECT 1
			FROM groups
			WHERE admin = ? AND id = ?)`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var exists bool
	err = stmt.QueryRowContext(ctx, userID, groupID, userID, groupID).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (g *group) GetUsersThatCanJoinGroupRepository(ctx context.Context, groupID int) ([]entity.User, error) {
	query := `
		SELECT u.id, u.nickname, u.avatar
		FROM users u
		WHERE u.id NOT IN (
			SELECT gm.member_id
			FROM group_members gm
			WHERE gm.group_id = ?
		)
		AND u.id != (
			SELECT g.admin
			FROM groups g
			WHERE g.id = ?
		)
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, groupID, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Nickname, &user.Avatar)
		if err != nil {
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// --------------------events----------------------------
// --------------------events----------------------------
func (g *group) CreateEventRepository(ctx context.Context, event entity.Event) (int, int, error) {
	query := `
		INSERT INTO events (user_id, group_id, title, description, date, location)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, ctx.Value(entity.ContextID).(int), event.GroupID, event.Title, event.Description, event.Date, event.Location)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	eventID, err := result.LastInsertId()
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	return int(eventID), http.StatusOK, nil
}

func (g *group) GetEventsRepository(ctx context.Context, groupID int) ([]entity.Event, error) {
	query := `
		SELECT 
			e.id, e.group_id, e.user_id, e.title, e.description, e.date, e.created_at, e.location,
			COALESCE(en.status, 0) as going
		FROM events e
		LEFT JOIN engagement en ON en.event_id = e.id AND en.user_id = ?
		WHERE e.group_id = ?
	`

	var events []entity.Event
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return events, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, ctx.Value(entity.ContextID).(int), groupID)
	if err != nil {
		return events, err
	}
	defer rows.Close()

	for rows.Next() {
		var event entity.Event
		err := rows.Scan(
			&event.ID,
			&event.GroupID,
			&event.UserID,
			&event.Title,
			&event.Description,
			&event.Date,
			&event.CreatedAt,
			&event.Location,
			&event.Going, 
		)
		if err != nil {
			return events, err
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return events, err
	}
	return events, nil
}


func (g *group) GetEventRepository(ctx context.Context, eventID int, userID int) (entity.Event, error) {
	query := `
		SELECT 
			e.id, e.group_id, e.user_id, e.title, e.description, e.date, e.created_at, e.location,
			COALESCE(en.status, 0) as going
		FROM events e
		LEFT JOIN engagement en ON en.event_id = e.id AND en.user_id = ?
		WHERE e.id = ?
	`

	var event entity.Event
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return event, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, userID, eventID).Scan(
		&event.ID,
		&event.GroupID,
		&event.UserID,
		&event.Title,
		&event.Description,
		&event.Date,
		&event.CreatedAt,
		&event.Location,
		&event.Going, 
	)
	if err != nil {
		return event, err
	}
	return event, nil
}


func (g *group) VoteEventRepository(ctx context.Context, eventId, status int) (int, int, error) {
	// search in engagement table when user voted an event if exist then return status code 400
	// other wize create new vote with eventId and userId and status
	query := `SELECT COUNT(*) FROM engagement WHERE user_id = ? AND event_id = ?`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	defer stmt.Close()
	var count int
	err = stmt.QueryRowContext(ctx, ctx.Value(entity.ContextID).(int), eventId).Scan(&count)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	if count > 0 {
		return 0, http.StatusBadRequest, errors.New("you already voted this event")
	}
	query = `INSERT INTO engagement (user_id, event_id, status) VALUES (?, ?, ?)`
	stmt, err = g.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(ctx.Value(entity.ContextID).(int), eventId, status)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	eventID, err := result.LastInsertId()
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	return int(eventID), status, nil
}

/*-------------------notification--------------------*/

func (g *group) CreateInvitationNotification(ctx context.Context, invt entity.Invitation) (int, int, error) {
	query := `
		INSERT INTO notification (type, sender_id, receiver_id)
		VALUES (?, ?, ?)
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, entity.GroupInvitationNotification, invt.InviterID, invt.InvitedID)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	eventID, err := result.LastInsertId()
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	return int(eventID), http.StatusOK, nil
}

func (g *group) CreateRequestJoiningNotification(ctx context.Context, notif entity.Notification) (int, int, error) {
	query := `INSERT INTO notification (type, sender_id, group_id, receiver_id)
	VALUES (?, ?, ?, ?)`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, entity.GroupParticipationNotification, notif.SenderId, notif.GroupId, notif.ReceiverID)
	if err != nil {
		return 0, http.StatusBadRequest, err
	}
	eventID, err := result.LastInsertId()
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	return int(eventID), http.StatusOK, nil
}

func (g *group) CreateEventNotification(ctx context.Context, event entity.Event) (int, int, error) {
	query := `INSERT INTO notification (type, sender_id, group_id, event_id)
	VALUES (?, ?, ?, ?)`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}
	res, err := stmt.ExecContext(ctx, entity.EventNotification, ctx.Value(entity.ContextID), event.GroupID, event.ID)
	if err != nil {
		return 0, http.StatusBadRequest, errors.New("invalid event credentials")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, http.StatusInternalServerError, err
	}

	return int(id), http.StatusOK, nil
}

func (g *group) getNotificationById(ctx context.Context, id int) (entity.Notification, error) {
	query := `
		SELECT * FROM notification WHERE id = ?
	`
	var notification entity.Notification
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return notification, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, id)
	err = row.Scan(
		&notification.Id,
		&notification.Type,
		&notification.GroupId,
		&notification.SenderId,
		&notification.ReceiverID,
		&notification.Accepted,
	)
	if err != nil {
		return notification, err
	}
	return notification, nil
}

func (g *group) RemoveNotificationById(ctx context.Context, id int) error {
	query := `
		DELETE FROM notification WHERE id = ?`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
