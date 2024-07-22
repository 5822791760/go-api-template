-- +goose Up
-- +goose StatementBegin
CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    bookno VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2),
    summary TEXT,
    author_id INTEGER,
    amount INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_books_author_id_authors FOREIGN KEY(author_id) REFERENCES authors(id) ON DELETE SET NULL
);

CREATE INDEX idx_books_author_id ON books(author_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_books_author_id;
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
