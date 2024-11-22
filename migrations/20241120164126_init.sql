-- +goose Up
-- +goose StatementBegin
CREATE TYPE gender_type AS ENUM ('Male', 'Female', 'Non-binary');

CREATE TABLE post(
    id INT PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    gender gender_type,
    ip_address TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE post IF exists;
-- +goose StatementEnd
