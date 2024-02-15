package driver

import (
	"log"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/tusmasoma/clean-architecture-campfinder/adapter/controller"
	"github.com/tusmasoma/clean-architecture-campfinder/adapter/gateway"
	"github.com/tusmasoma/clean-architecture-campfinder/adapter/presenter"
	"github.com/tusmasoma/clean-architecture-campfinder/config"
	"github.com/tusmasoma/clean-architecture-campfinder/usecase/interactor"
)

func InitRoute() *chi.Mux {
	conn, err := NewDB()
	if err != nil {
		log.Printf("Database connection failed: %s\n", err)
		return nil
	}

	user := controller.User{
		OutputFactory: presenter.NewUserOutputPort,
		InputFactory:  interactor.NewUserInputPort,
		RepoFactory:   gateway.NewUserRepository,
		Conn:          conn,
	}

	spot := controller.Spot{
		OutputFactory: presenter.NewSpotOutputPort,
		InputFactory:  interactor.NewSpotInputPort,
		RepoFactory:   gateway.NewSpotRepository,
		Conn:          conn,
	}

	comment := controller.Comment{
		OutputFactory:   presenter.NewCommentOutputPort,
		InputFactory:    interactor.NewCommentInputPort,
		RepoFactory:     gateway.NewCommentRepository,
		UserRepoFactory: gateway.NewUserRepository,
		Conn:            conn,
	}

	img := controller.Image{
		OutputFactory:   presenter.NewImageOutputPort,
		InputFactory:    interactor.NewImageInputPort,
		RepoFactory:     gateway.NewImageRepository,
		UserRepoFactory: gateway.NewUserRepository,
		Conn:            conn,
	}

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Origin"},
		ExposedHeaders:   []string{"Link", "Authorization"},
		AllowCredentials: false,
		MaxAge:           config.PreflightCacheDurationSeconds,
	}))

	//r.Use(middleware.Logging)

	r.Route("/api", func(r chi.Router) {
		r.Route("/user", func(r chi.Router) {
			r.Post("/create", user.HandleUserCreate)
			//r.Post("/login", userHandler.HandleUserLogin)
			r.Group(func(r chi.Router) {
				//r.Use(authMiddleware.Authenticate)
				//r.Get("/api/user/logout", userHandler.HandleUserLogout)
			})
		})

		r.Route("/spot", func(r chi.Router) {
			r.Get("/", spot.HandleSpotGet)
			r.Post("/create", spot.HandleSpotCreate)
		})

		r.Route("/comment", func(r chi.Router) {
			r.Get("/", comment.HandleCommentGet)
			r.Group(func(r chi.Router) {
				//r.Use(authMiddleware.Authenticate)
				r.Post("/create", comment.HandleCommentCreate)
				r.Post("/update", comment.HandleCommentUpdate)
				r.Delete("/delete", comment.HandleCommentDelete)
			})
		})

		r.Route("/img", func(r chi.Router) {
			r.Get("/", img.HandleImageGet)
			r.Group(func(r chi.Router) {
				//r.Use(authMiddleware.Authenticate)
				r.Post("/create", img.HandleImageCreate)
				r.Post("/delete", img.HandleImageDelete)
			})
		})
	})

	return r
}
