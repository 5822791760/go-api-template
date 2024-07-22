-- +goose Up
-- +goose StatementBegin
CREATE TABLE book_slips (
  id SERIAL PRIMARY KEY,
  paid_amount DECIMAL(10,2) NOT NULL,
  book_id INTEGER,
  created_by_id INTEGER,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  buyer text NOT NULL,
  book_name text NOT NULL,
  CONSTRAINT fk_book_slips_created_by_id_users FOREIGN KEY(created_by_id) REFERENCES users(id) ON DELETE SET NULL,
  CONSTRAINT fk_book_slips_book_id_books FOREIGN KEY(book_id) REFERENCES books(id) ON DELETE SET NULL
);

CREATE INDEX idx_book_slips_created_by_id ON book_slips(created_by_id);
CREATE INDEX idx_book_slips_book_id ON book_slips(book_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_book_slips_created_by_id;
DROP INDEX IF EXISTS idx_book_slips_book_id;
DROP TABLE book_slips;
-- +goose StatementEnd