BEGIN;
    CREATE TABLE IF NOT EXISTS users_code (
        id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
        user_id uuid REFERENCES users(id),
        code_buffer TEXT NOT NULL,
        input_buffer TEXT,
        ts TIMESTAMPTZ DEFAULT current_timestamp,
        ts_mod TIMESTAMPTZ DEFAULT current_timestamp
    );
COMMIT;