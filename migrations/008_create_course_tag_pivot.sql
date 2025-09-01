CREATE TABLE IF NOT EXISTS course_tags (
    course_id INT NOT NULL REFERENCES courses(id) ON DELETE CASCADE,
    tag_id INT NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
    PRIMARY KEY(course_id, tag_id)
);
