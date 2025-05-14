// Use DBML to define your database structure
// Docs: https://dbml.dbdiagram.io/docs
//link: https://dbdiagram.io/d/67a6742f263d6cf9a06e3470
// creating a web hood and now testing
// /some not change
Table follows {
  following_user_id integer [ref: > user.id] // Reference to user id
  followed_user_id integer [ref: > user.id] // Reference to user id
  created_at timestamp
}

Table user {
    id int [primary key]
    Email text
    Password text
    First text
    Last text
    Date_Of_Birth time
    Avatar blob
    Nickname text
    About_Me text
    status int // Global or private account
}

Table post {
  id integer [primary key]
  title text
  image blob
  content text [note: 'Content of the post']
  user_id integer [ref: > user.id] // Reference to user id
  status int // Global, for friends, or private
  group int [ref: > group.id] // Reference to group id 
  created_at timestamp
  updated_at timestamp
}

Table comment {
  id integer [primary key]
  post_id integer [ref: > post.id] // Reference to post id
  user_id integer [ref: > user.id] // Reference to user id
  content text
}

Table event {
  id integer [primary key]
  group_id int [ref: > group.id]
  user_id int [ref: > user.id]
  title text
  description text
  event_time timestamp
  create_at timestamp
}

Table engagement {
  id integer [primary key]
  comment_id int [ref: > comment.id] // Reference to comment
  post_id int [ref: > post.id] // Reference to post
  event_id int [ref: > event.id]
  user_id int [ref: > user.id] // Reference to user
  status int // Like, dislike, or not
}

Table notification {
  id integer [primary key]
  type int 
  group_id int [ref: > group.id] // Reference to group
  sender_id int [ref: > user.id] // Reference to sender user
  accepted boolean
}

TABLE session {
    id integer [primary key]
    user_id integer [ref: > user.id] // Reference to user id
    token text
    data text
    expiry timestamp
}

Table group {
  id integer [primary key]
  name text
  type int // Real group, fake, or messages group
  admin int [ref: > user.id] // Reference to user id
}

Table group_members {
  id integer [primary key]
  member_id integer [ref: > user.id] // Reference to user id
  group_id integer [ref: > group.id] // Reference to group id
}

Table chat {
  id integer [primary key]
  name text
}

Table message {
  id integer [primary key]
  chat_id integer [ref: > chat.id] // Reference to chat id
  receiver_id integer [ref: > user.id] // Reference to user id
  sender_id integer [ref: > user.id] // Reference to user id
  group integer [ref: > group.id] // Reference to group id
  created_at timestamp 
  content text
}



paths = {
  user => 
      /valid-token (post)
      /login
      /register
      /profile/id(update, read)
      /profile/followers/id (get)
      /profile/followed/id
      /follow (delete, post)
  post => 
    /posts
    /post/id
    /post(delete, create, update)
    /comment(create, delete, update)
    /engage
    /group/id
    /group(create, update, delete)
    /group/invite/nickname
    /group/request/group_id
  chat => 
    /chat/ (ws)
    /chat/group-name (ws)

}
