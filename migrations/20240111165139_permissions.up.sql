BEGIN;

CREATE TABLE IF NOT EXISTS permissions
(
    id         SERIAL         NOT NULL,
    name       VARCHAR(128) NOT NULL,
    label      VARCHAR(128) NOT NULL,
    created_by VARCHAR(128),
    updated_by VARCHAR(128),
    deleted_by VARCHAR(128),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (id)
    );

COMMIT;
