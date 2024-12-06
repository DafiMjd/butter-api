CREATE TABLE IF NOT EXISTS butter.posts (
    id UUID PRIMARY KEY,
    user_id UUID,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES butter.users(id)
);