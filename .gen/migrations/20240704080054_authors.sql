-- +goose Up
-- +goose StatementBegin
CREATE TABLE authors (
  id   SERIAL PRIMARY KEY,
  name text      NOT NULL,
  bio  text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE authors;
-- +goose StatementEnd
