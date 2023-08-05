CREATE TABLE Users (
    user_id SERIAL PRIMARY KEY ,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(100),
    descriptions varchar(250),
    avatar_url VARCHAR(255),
    date_joined TIMESTAMP NOT NULL
);