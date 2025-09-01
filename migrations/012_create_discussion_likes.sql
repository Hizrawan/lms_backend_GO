CREATE TABLE IF NOT EXISTS discussion_likes (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    likeable_id INT NOT NULL,
    likeable_type TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
