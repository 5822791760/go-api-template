-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
  id   SERIAL PRIMARY KEY,
  name text      NOT NULL,
  email  text      NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd
