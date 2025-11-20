create table refresh_tokens
(
    id         serial primary key,
    user_id    int,
    jti        text,
    token_hash text,
    expires_at TIMESTAMP,
    revoked    bool,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    foreign key (user_id) references users (id)
)