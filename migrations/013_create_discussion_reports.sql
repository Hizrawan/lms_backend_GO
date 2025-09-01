CREATE TABLE IF NOT EXISTS discussion_reports (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    reportable_id INT NOT NULL,
    reportable_type TEXT NOT NULL,
    reason TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);
