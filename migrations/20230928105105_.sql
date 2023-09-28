-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                       id bigserial primary key,
                       name text not null,
                       email text not null,
                       created_at timestamp default now() not null,
                       updated_at timestamp default now() not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
