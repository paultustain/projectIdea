-- +goose Up 
CREATE TABLE projects ( 
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL, 
	updated_at TIMESTAMP NOT NULL, 
	title TEXT UNIQUE NOT NULL,
	description TEXT
);

-- +goose Down
DROP TABLE projects;
