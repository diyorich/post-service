-- +goose Up
-- +goose StatementBegin
CREATE TABLE post(
    id SERIAL PRIMARY KEY,
    external_id TEXT NOT NULL UNIQUE,
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    ip_address TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE post IF exists;
-- +goose StatementEnd
