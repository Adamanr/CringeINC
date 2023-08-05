CREATE TABLE Comments (
    comment_id INT PRIMARY KEY,
    user_id INT,
    post_id INT,
    content TEXT,
    date_posted TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (user_id),
    FOREIGN KEY (post_id) REFERENCES Posts (post_id)
);