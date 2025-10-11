CREATE TABLE IF NOT EXISTS categories (
    id          BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name        TEXT    NOT NULL,
    user_id     BIGINT  NULL,
    CONSTRAINT categories_user_fk
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL,
    CONSTRAINT categories_user_name_uniq
        UNIQUE (user_id, name)
);
