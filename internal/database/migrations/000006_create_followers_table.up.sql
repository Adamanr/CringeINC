CREATE TABLE Followers (
    follower_id INT PRIMARY KEY,
    user_id INT,
    follower_user_id INT,
    date_followed TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (user_id),
    FOREIGN KEY (follower_user_id) REFERENCES Users (user_id)
);