CREATE TABLE IF NOT EXISTS attachments (
    id SERIAL PRIMARY KEY,
    file_name TEXT NOT NULL,
    file_path TEXT NOT NULL,
    file_type TEXT,
    file_size BIGINT DEFAULT 0,
    user_id INT REFERENCES users(id) ON DELETE SET NULL,
    attachable_id INT NOT NULL,
    attachable_type TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
