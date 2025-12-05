# blogging-platform-API

A simple RESTful API for creating, reading, updating, deleting, and searching blog posts.
Built with **Go** + **MySQL** following a clean, layered architecture.

---

## **Features**

* CRUD blog posts
* Tags (many-to-many)
* Search (`?term=`)
* MySQL storage
* Clean separation: handler → service → repository → DB

---

## **Project Structure**

```
cmd/server        → entrypoint
config            → config loading
internal/blog     → handlers, services, repository, models
internal/database → MySQL connection
scripts           → SQL schema
.env.example      → environment variables
```

---

## **Environment Variables**

Create a `.env` file:

```
DB_USER=root
DB_PASSWORD=yourpassword
DB_ADDR=localhost:3306
DB_NAME=blog_api
PORT=8080
```

A matching `.env.example` is recommended.

---

## **Setup**

### 1. Install deps

```sh
go mod tidy
```

### 2. Load env vars

Your `config.go` should read from `.env` (e.g., using `github.com/joho/godotenv`).

### 3. Init DB

```sh
mysql -u root -p blog_api < scripts/data.sql
```

### 4. Run server

```sh
go run cmd/server/main.go
```

Runs on:
`http://localhost:${SERVER_PORT}`

---

## **API Endpoints**

### **POST /posts**

Create a post:

```json
{
  "title": "My Post",
  "content": "Hello world",
  "category": "Tech",
  "tags": ["Go", "Backend"]
}
```

### **GET /posts**

`GET /posts`
`GET /posts?term=tech`

### **GET /posts/{id}**

### **PUT /posts/{id}**

### **DELETE /posts/{id}**

---
https://roadmap.sh/projects/blogging-platform-api
