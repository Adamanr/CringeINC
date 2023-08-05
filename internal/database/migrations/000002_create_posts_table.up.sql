CREATE TABLE Posts (
    post_id INT PRIMARY KEY,
    user_id INT,
    content TEXT,
    image_url VARCHAR(255),
    date_posted TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (user_id)
);