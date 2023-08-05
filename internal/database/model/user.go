package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	UserId       int       `json:"user_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	FullName     *string   `json:"full_name"`
	Descriptions *string   `json:"descriptions"`
	AvatarURL    *string   `json:"avatar_url"`
	DateJoined   time.Time `json:"date_joined"`
}

func (u User) ToString() string {
	return fmt.Sprintf("User: {\n\t email: %s,\n\t login: %s,\n\t password: %s,\tdescription: %s,\n\t, created_at: %v\n}",
		u.Email, u.Username, u.Password, u.Descriptions, u.DateJoined)
}

func (u User) JSON() []byte {
	data, _ := json.Marshal(&u)
	return data
}
