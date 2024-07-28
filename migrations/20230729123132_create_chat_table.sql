-- +goose Up
create table chat
(
    chat_id    serial primary key,
    name       text      not null,
    user_ids   INTEGER[],
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table chat;

