CREATE TABLE IF NOT EXISTS discussion_subscriptions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    thread_id INT NOT NULL REFERENCES discussion_threads(id),
    notify_on_reply BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT NOW()
);
