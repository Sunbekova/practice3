CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    balance NUMERIC(10,2) DEFAULT 0
);

INSERT INTO users (name, email, balance) VALUES
('Mike Vizovsky', 'mike@example.com', 200.00),
('Lightning MCqueen', 'kchau@example.com', 300.00);
