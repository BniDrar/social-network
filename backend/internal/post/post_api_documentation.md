
# 📘 Post API Documentation

This document outlines the API endpoints related to post functionalities.

---

## 🔹 GET `/api/posts`

Retrieve a list of posts with pagination.

### Query Parameters

| Name   | Type   | Required | Description            |
|--------|--------|----------|------------------------|
| limit  | int    | Yes      | Number of posts to return |
| offset | int    | Yes      | Offset for pagination     |

### Responses

- `200 OK` – Returns an array of posts.
- `400 Bad Request` – Invalid or missing query parameters.

### Example Request

```
GET /api/posts?limit=10&offset=0
```

---

## 🔹 POST `/api/post/create`

Create a new post. This endpoint accepts `multipart/form-data`.

### Form Data Parameters

| Field           | Type            | Required | Description                                                                 |
|----------------|-----------------|----------|-----------------------------------------------------------------------------|
| content        | string          | Yes      | Text content of the post                                                   |
| status         | int (0,1,2)     | Yes      | Visibility status: `0` = custom, `1` = friends, `2` = global               |
| allowed_viewers| []int (form key: `allowed_viewers[]`) | Required if `status=0` | Array of user IDs allowed to view (custom visibility only) |
| media          | file (image)    | No       | Optional image file to attach                                              |

> ⚠️ All fields should be sent using `multipart/form-data`.

### Status Values

| Value | Meaning        | Behavior                                      |
|-------|----------------|-----------------------------------------------|
| 0     | Custom         | Requires `allowed_viewers[]`                 |
| 1     | Friends Only   | Visible to friends                            |
| 2     | Global         | Visible to everyone                           |

### Responses

- `200 OK` – Returns created post ID.
- `400 Bad Request` – Missing required fields, invalid `status`, or `allowed_viewers` missing when `status=0`.

### Example Form Submission (via `curl`)

```bash
curl -X POST http://localhost:8080/api/post/create \
  -F "content=Hello World" \
  -F "status=0" \
  -F "allowed_viewers[]=2" \
  -F "allowed_viewers[]=5" \
  -F "media=@/path/to/image.jpg"
```

### Example Response

```json
{
  "post_id": 42
}
```

---

## 🔹 GET `/api/post/`

Retrieve a single post by ID.

### Query Parameters

| Name     | Type | Required | Description         |
|----------|------|----------|---------------------|
| post_id  | int  | Yes      | ID of the post      |

### Responses

- `200 OK` – Returns the post data.
- `400 Bad Request` – Missing or invalid `post_id`.

### Example Request

```
GET /api/post/?post_id=42
```

---

## 🔹 POST `/api/post/vote`

React to a post (like, upvote, etc.).

### Query Parameters

| Name     | Type | Required | Description         |
|----------|------|----------|---------------------|
| post_id  | int  | Yes      | ID of the post      |

### Responses

- `200 OK` – Reaction was successful.
- `400 Bad Request` – Invalid post ID or missing query.

### Example Request

```
POST /api/post/vote?post_id=42
```

---

## 🔹 GET `/api/pictures/{filename}`

Serve media files such as images or videos associated with posts.

### Path Parameters

| Name      | Type   | Required | Description              |
|-----------|--------|----------|--------------------------|
| filename  | string | Yes      | Name of the media file   |

### Responses

- `200 OK` – Returns the media file.
- `404 Not Found` – File does not exist or is a directory.

### Example Request

```
GET /api/pictures/uploads/img_42.jpg
```

---
