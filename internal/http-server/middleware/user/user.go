package user

import (
	"cringeinc_server/internal/database/model"
	"cringeinc_server/internal/database/postgres/controller"
	"cringeinc_server/internal/http-server/middleware/assets"
	"encoding/json"
	"github.com/go-chi/chi"
	"log/slog"
	"net/http"
)

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func Authorization(db *model.Storage) http.HandlerFunc {
	var userRequest *model.User

	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
			slog.Error("Error parse user")
			assets.ErrorResponse[string](w, "Error decode user", 500)
			return
		}

		user, err := controller.Authorization(db, userRequest.Username, userRequest.Password)
		if err != nil {
			assets.ErrorResponse[string](w, err.Error(), 500)
			return
		}

		assets.ErrorResponse[string](w, user.ToString(), 200)
	}
}

func Registration(storage *model.Storage) http.HandlerFunc {
	var userRequest *model.User
	return func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
			assets.ErrorResponse[string](w, "Error decode user", 500)
			return
		}

		if err := controller.Registration(storage, userRequest); err != nil {
			assets.ErrorResponse(w, err.Error(), 500)
			return
		}

		assets.ErrorResponse[bool](w, true, 200)
	}
}

func User(storage *model.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "user_id")
		user, err := controller.User(storage, userId)
		if err != nil {
			assets.ErrorResponse[string](w, err.Error(), 500)
			return
		}

		assets.ErrorResponse[*model.User](w, user, 200)
	}
}

func SetUser(storage *model.Storage) http.HandlerFunc {
	var userRequest *model.User
	return func(w http.ResponseWriter, r *http.Request) {
		userId := chi.URLParam(r, "user_id")

		if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
			assets.ErrorResponse[string](w, "Error decode user", 500)
			return
		}

		user, err := controller.User(storage, userId)
		if err != nil {
			assets.ErrorResponse[string](w, err.Error(), 500)
			return
		}

		assets.ErrorResponse[*model.User](w, user, 500)

	}
}
