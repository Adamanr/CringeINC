CREATE TABLE Friends (
    friendship_id INT PRIMARY KEY,
    user_id INT,
    friend_id INT,
    date_connected TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users (user_id),
    FOREIGN KEY (friend_id) REFERENCES Users (user_id)
);