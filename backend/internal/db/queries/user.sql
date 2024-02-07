-- name: ListTodos :many
SELECT * FROM todos;

-- name: CreateTodo :one
INSERT INTO todos (text) VALUES ($1) RETURNING *;