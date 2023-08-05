package controller

import (
	"context"
	"cringeinc_server/internal/database/model"
	"cringeinc_server/internal/database/postgres"
	"errors"
	"log"
	"log/slog"
	"strconv"
	"time"
)

func Registration(storage *model.Storage, user *model.User) error {
	if !postgres.CheckTables(storage, "users") {
		return errors.New("tables not found")
	}

	if err := storage.DB.QueryRow(context.Background(), "INSERT INTO users(username, email, password, date_joined) VALUES ($1, $2, $3, $4)",
		user.Username, user.Email, user.Password, time.Now()).Scan(&user); err != nil {
		return errors.New("the user is already registered")
	}

	return nil
}

func Authorization(storage *model.Storage, username, password string) (*model.User, error) {
	if !postgres.CheckTables(storage, "users") {
		return nil, errors.New("tables not found")
	}

	var user *model.User
	if err := storage.DB.QueryRow(context.Background(), `SELECT * FROM users WHERE username = ? AND password = ?`, username, password).Scan(&user); err != nil {
		slog.Warn("users not select!", slog.String("error", err.Error()))
		return nil, errors.New("users not found in the tables! Error: " + err.Error())
	}

	return user, nil
}

func User(storage *model.Storage, userIdStr string) (*model.User, error) {
	if !postgres.CheckTables(storage, "users") {
		return nil, errors.New("tables not found")
	}

	_, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, err
	}
	user := &model.User{} // Initialize user as a pointer to model.User
	if err := storage.DB.QueryRow(context.Background(), "SELECT * FROM users WHERE user_id = 1").Scan(
		&user.UserId, &user.Username, &user.Email, &user.Password, &user.FullName, &user.Descriptions, &user.AvatarURL, &user.DateJoined); err != nil {
		return nil, err
	}
	log.Println(user)
	return user, nil
}

func SetUser(storage *model.Storage, user *model.User) error {
	if !postgres.CheckTables(storage, "users") {
		return errors.New("tables not found")
	}

	if _, err := storage.DB.Exec(context.Background(), `UPDATE users SET username = $1, email = $2, password = $3, full_name = $4, descriptions = $5, avatar_url = $6`,
		&user.Username, &user.Email, &user.Password, &user.FullName, &user.Descriptions, &user.AvatarURL); err != nil {
		return errors.New("the user is already registered")
	}

	return nil
}
