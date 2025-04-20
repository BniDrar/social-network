package comment

func (c *comment) CanSeePost(userid, postid int) bool {
	prep, err := c.db.Prepare(`SELECT
		    1
		FROM posts AS post 
		INNER JOIN users AS user ON user.id = post.user_id
		LEFT JOIN group_members AS gm ON gm.group_id=post.group_id AND gm.member_id = $1
		LEFT JOIN follows AS follow ON follow.followed_id = post.user_id AND follow.follower_id = 1
		WHERE
		    (post.group_id IS NOT NULL AND gm.member_id = $1 AND post.id=$2)
		    OR 
		    (post.group_id IS NULL AND follow.follower_id = $1 AND post.id=$2);`)
	if err != nil {
		return false
	}
	row := prep.QueryRow(userid, postid)
	var res bool
	err = row.Scan(&res)
	if err != nil {
		return false
	}
	return res
}
