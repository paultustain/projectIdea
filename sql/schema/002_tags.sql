-- +goose Up 
CREATE TABLE tags (
	tag TEXT PRIMARY KEY, 
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE tags;
