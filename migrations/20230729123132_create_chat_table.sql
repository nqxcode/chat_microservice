-- +goose Up
create table chat
(
    chat_id    bigserial primary key,
    name       text      not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

create table message
(
    message_id bigserial primary key,
    chat_id    bigint    not null REFERENCES chat (chat_id),
    "from"     text      not null,
    message    text      not null,
    sent_at    timestamp not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

create table chat_to_user
(
    chat_to_user_id bigserial primary key,
    chat_id         BIGINT    not null references chat (chat_id),
    user_id         BIGINT    not null,
    created_at      timestamp not null default now(),
    unique (chat_id, user_id)
);

-- +goose Down
drop table chat_to_user;
drop table message;
drop table chat;

