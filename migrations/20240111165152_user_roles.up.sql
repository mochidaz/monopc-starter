BEGIN;

CREATE TABLE IF NOT EXISTS user_roles
(
    id         UUID         NOT NULL,
    user_id    UUID         NOT NULL,
    role_id    int         NOT NULL,
    created_by VARCHAR(128),
    updated_by VARCHAR(128),
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
    );

COMMIT;
