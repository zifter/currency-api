-- +goose Up
CREATE TABLE rate_post
(
    id SERIAL PRIMARY KEY,
    postDate text,
    msg text
);

-- +goose Down
DROP TABLE rate_post;