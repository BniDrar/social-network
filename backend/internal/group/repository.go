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