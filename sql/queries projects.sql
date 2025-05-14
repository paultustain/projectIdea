-- name: CreateProject :one
INSERT INTO projects (id, created_at, updated_at, title, description)
VALUES (
	gen_random_uuid(), 
	NOW(), 
	NOW(), 
	$1, 
	$2
)
RETURNING *;

