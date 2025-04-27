# 📄 Comment API Documentation

## Table of Contents
- [Add a Comment](#-add-a-comment)
- [Get Comments for a Post](#-get-comments-for-a-post)
- [Notes](#-notes)

---

## 🔹 Add a Comment

**Endpoint:**  
`POST /api/comment/add`

**Headers:**

```makefile
Content-Type: multipart/form-data
```

**Authentication:**  
- User must be authenticated via **session** (cookie must include a valid session).

---

### Request (multipart/form-data)

| Field         | Type    | Required | Description                         |
|---------------|---------|----------|-------------------------------------|
| post_id       | integer | ✅        | ID of the post to comment on        |
| user_id       | integer | ✅        | ID of the user making the comment   |
| creater_name  | string  | ✅        | Username of the commenter           |
| content       | string  | ✅        | The comment text                   |
| image         | file    | ❌        | Optional image file attached        |

---

### Example Request (no image)

```json
{
  "post_id": 123,
  "user_id": 45,
  "creater_name": "john_doe",
  "content": "This is a comment."
}
```

---

### Response (application/json)

```json
{
  "id": 789
}
```

---

### Status Codes

| Code | Meaning                                   |
|------|-------------------------------------------|
| 200  | OK – Comment added successfully           |
| 400  | Bad Request – Invalid data or upload error |
| 401  | Unauthorized – User session invalid or missing |
| 405  | Method Not Allowed – Only POST is accepted |

---

## 🔹 Get Comments for a Post

**Endpoint:**  
`GET /api/comment/get?id=<post_id>`

**Headers:**

```makefile
Content-Type: application/json
```

**Authentication:**  
- User must be authenticated via **session** (cookie must include a valid session).

---

### Query Parameters

| Param | Type    | Required | Description                       |
|-------|---------|----------|-----------------------------------|
| id    | integer | ✅        | ID of the post to fetch comments for |

---

### Example Request

```pgsql
GET /api/comment/get?id=123
```

---

### Response (application/json)

```json
[
  {
    "id": 789,
    "creater_name": "john_doe",
    "post_id": 123,
    "user_id": 45,
    "content": "This is a comment.",
    "image": "https://yourdomain.com/uploads/image.png"
  },
  {
    "id": 790,
    "creater_name": "jane_smith",
    "post_id": 123,
    "user_id": 46,
    "content": "Another comment.",
    "image": ""
  }
]
```

---

### Status Codes

| Code | Meaning                                  |
|------|------------------------------------------|
| 200  | OK – Comments returned successfully      |
| 400  | Bad Request – Invalid post ID             |
| 401  | Unauthorized – User session invalid or missing |
| 405  | Method Not Allowed – Only GET is accepted |

---

## 📌 Notes
- All requests require a valid authenticated session.
- If the session is missing, invalid, or expired, the server will return a **401 Unauthorized** error.
- Uploaded images are accessible via the returned URL in the `image` field.
- Only `POST` and `GET` methods are supported for the endpoints mentioned.

---