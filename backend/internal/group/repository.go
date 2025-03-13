package group

import (
	"errors"

	"socialNetwork/entity"
)

func (g *group) GetGroupsByUserID(userID int, limit int, offset int) (entity.Groups, error) {
	query := `
		SELECT 
			g.id, 
			g.name, 
			g.type, 
			g.admin,
			COUNT(DISTINCT gm.user_id) AS member_count,
			COUNT(DISTINCT p.id) AS post_count
		FROM 
			groups g
		JOIN 
			group_members gm ON g.id = gm.group_id
		LEFT JOIN 
			posts p ON g.id = p.group_id
		WHERE 
			gm.user_id = ?
		GROUP BY 
			g.id
		ORDER BY 
			g.id DESC
		LIMIT ? OFFSET ?
	`

	rows, err := g.db.Query(query, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groups := make(entity.Groups, 0)
	for rows.Next() {
		group := entity.Group{}
		err := rows.Scan(
			&group.ID,
			&group.Name,
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

	if len(groups) == 0 {
		return nil, errors.New("no groups found for user")
	}

	return groups, nil
}

// this function is used to get the group by id and the user id
func (g *group) GetGroupByIdRepository(userID, groupID int) (entity.Group, error) {
	query := `
		SELECT
			g.id,
			g.name,
			g.type,
			g.admin,
			COUNT(DISTINCT gm.user_id) AS member_count,
			COUNT(DISTINCT p.id) AS post_count
			FROM groups g
			JOIN group_members gm ON g.id = gm.group_id
			LEFT JOIN posts p ON g.id = p.group_id
			WHERE g.id = ? AND gm.user_id = ?
			GROUP BY g.id
			ORDER BY g.id DESC
			LIMIT 1
			`
	rows, err := g.db.Query(query, groupID, userID)
	if err != nil {
		return entity.Group{}, err
	}
	defer rows.Close()
	group := entity.Group{}
	if rows.Next() {
		err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.Type,
			&group.Admin,
			&group.MemberCount,
			&group.PostCount,
		)
		if err != nil {
			return entity.Group{}, err
		}
	}
	if err := rows.Err(); err != nil {
		return entity.Group{}, err
	}
	if group.ID == 0 {
		return entity.Group{}, errors.New("group not found")
	}
	return group, nil
}
