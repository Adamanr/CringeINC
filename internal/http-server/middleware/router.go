package middleware

import (
	"cringeinc_server/internal/database/model"
	"cringeinc_server/internal/http-server/middleware/user"
	"github.com/go-chi/chi"
)

func SetRouter(router *chi.Mux, storage *model.Storage) {
	router.Route("/api", func(r chi.Router) {
		r.Route("/users", func(rt chi.Router) {
			rt.Post("/sign_in", user.Authorization(storage))
			rt.Post("/sign_up", user.Registration(storage))
			rt.Post("/logout", user.Logout())

			// /api/users/:id - GET-запрос для получения профиля конкретного пользователя.
			rt.Get("/{user_id}", user.User(storage))
			// /api/users/:id - PUT-запрос для обновления профиля пользователя
			rt.Put("/{user_id}", user.SetUser(storage))

			// /api/users/search - GET-запрос для поиска пользователей по заданным параметрам (например, по имени).
			rt.Get("/search", nil)

			// /api/users/{user_id}/friends - GET-запрос для получения списка друзей пользователя.
			rt.Get("/{user_id}/friends", nil)
			// /api/users/{user_id}/followers - GET-запрос для получения списка подписчиков пользователя.
			rt.Get("/{user_id}/followers", nil)
			// /api/users/{user_id}/follow - POST-запрос для подписки на пользователя.
			rt.Post("/{user_id}/follow", nil)
			// /api/users/{user_id}/unfollow - POST-запрос для отписки от пользователя.
			rt.Post("/{user_id}/unfollow", nil)
		})

		r.Route("/posts", func(rt chi.Router) {
			// /api/posts - GET-запрос для получения списка всех постов.
			rt.Get("/posts", nil)
			// /api/posts/{post_id} - GET-запрос для получения информации о конкретном посте.
			rt.Get("/posts/{post_id}", nil)
			// /api/posts - POST-запрос для создания нового поста.
			rt.Post("/posts", nil)
			// /api/posts/{post_id} - PUT-запрос для обновления информации о посте.
			rt.Put("/posts/{post_id}", nil)
			// /api/posts/{post_id} - DELETE-запрос для удаления поста.
			rt.Delete("/posts/{post_id}", nil)

			// /api/posts/{post_id}/comments - GET-запрос для получения комментариев к определенному посту.
			rt.Get("/{post_id}/comments", nil)
			// /api/posts/{post_id}/comments - POST-запрос для создания нового комментария к посту.
			rt.Post("/{post_id}/comments", nil)
			// /api/posts/{post_id}/comments/{comment_id} - DELETE-запрос для удаления комментария.
			rt.Delete("/{post_id}/comments/{comment_id}", nil)

			// /api/posts/{post_id}/like - POST-запрос для добавления лайка к посту.
			r.Post("/{post_id}/like", nil)
			// /api/posts/{post_id}/unlike - POST-запрос для удаления лайка у поста.
			r.Post("/{post_id}/unlike", nil)
			// /api/posts/{post_id}/rating - GET-запрос для получения рейтинга поста.
			r.Get("/{post_id}/rating", nil)

			// /api/posts/search - GET-запрос для поиска постов по заданным параметрам (например, по ключевым словам).
			r.Get("/search", nil)
		})
	})
}
