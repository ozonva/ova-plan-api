-- +goose Up
-- +goose StatementBegin
create table plans
(
    id          bigserial unique primary key,
    user_id     bigint,
    title       text,
    description text,
    created_at  timestamp with time zone,
    deadline_at timestamp with time zone
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table plans;
-- +goose StatementEnd
