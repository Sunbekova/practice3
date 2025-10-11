CREATE TABLE IF NOT EXISTS expenses (
    id           BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id      BIGINT        NOT NULL,
    category_id  BIGINT        NOT NULL,
    amount       NUMERIC(14,2) NOT NULL,
    currency     TEXT          NOT NULL,  
    spent_at     TIMESTAMPTZ   NOT NULL,
    created_at   TIMESTAMPTZ   NOT NULL DEFAULT NOW(),
    note         TEXT,

    CONSTRAINT expenses_user_fk
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT expenses_category_fk
        FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE RESTRICT,

    CONSTRAINT expenses_amount_positive CHECK (amount > 0),
    CONSTRAINT expenses_currency_len CHECK (char_length(currency) >= 3)
);
