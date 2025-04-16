package group

import (
	"context"
	"errors"
	"fmt"

	"socialNetwork/entity"
)

func (g *group) GetGroupsByUserID(ctx context.Context,userID int, limit int, offset int) (entity.Groups, error) {
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
	WHERE gm.member_id = ?
	GROUP BY g.id
	ORDER BY g.id DESC
  LIMIT ? OFFSET ?
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err !=nil {
		return nil, errors.New(fmt.Sprintf("error preparing query in GetGroupsByUserID function: %v", err))
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, userID, limit, offset)
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
			JOIN group_members gm ON g.id = gm.group_id
			LEFT JOIN posts p ON g.id = p.group_id
			WHERE g.id = ? AND gm.member_id = ?
			GROUP BY g.id
			ORDER BY g.id DESC
			LIMIT 1
			`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return entity.Group{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, groupID, userID)
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
	result, err := stmt.ExecContext(ctx, group.Name,group.Description, group.Type, group.Admin)
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

func (g *group) GetAllGroupsRepository(ctx context.Context, limit, offset, typeGroup int) (entity.Groups, error) {
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
		JOIN group_members gm ON g.id = gm.group_id
		LEFT JOIN posts p ON g.id = p.group_id
		WHERE g.type = ?
		GROUP BY g.id
		ORDER BY g.id DESC
		LIMIT ? OFFSET ?
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, typeGroup, limit, offset)
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

func (g *group)GetGroupMembersRepository(ctx context.Context, groupID int) ([]entity.User, error) {
	query := `
		SELECT gm.member_id, u.nickname, u.avatar
		FROM group_members gm
		JOIN users u ON gm.member_id = u.id
		WHERE gm.group_id = ?
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, groupID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var groupMembers []entity.User
	for rows.Next() {
		var groupMember entity.User
		err := rows.Scan(
			&groupMember.ID,
			&groupMember.Nickname,
			&groupMember.Avatar,
		)
		if err != nil {
			return nil, err
		}
		groupMembers = append(groupMembers, groupMember)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return groupMembers, nil
}

func (g *group) RequestToJoinGroupRepository(ctx context.Context, groupID, userID int) error {
	query := `
		INSERT INTO group_members (group_id, member_id, status)
		VALUES (?, ?, ?)
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, groupID, userID, 0)
	if err != nil {
		return err
	}
	return nil
}
func (g *group)AcceptRequestToJoinGroupRepository(ctx context.Context, groupID, userID int) error {
	query := `
		UPDATE group_members
		SET status = 1
		WHERE group_id = ? AND member_id = ?
	`
	stmt, err := g.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, groupID, userID)
	if err != nil {
		return err
	}
	return nil
}