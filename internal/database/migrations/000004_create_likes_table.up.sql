CREATE TABLE Likes (
    like_id INT PRIMARY KEY,
    user_id INT,
    post_id INT,
    date_liked TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (user_id),
    FOREIGN KEY (post_id) REFERENCES Posts (post_id)
);