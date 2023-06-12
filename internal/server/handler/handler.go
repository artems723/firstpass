package handler

import (
	"firstpass/internal/server/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	userService    service.UserService
	noteService    service.NoteService
	cardService    service.CardService
	binService     service.BinService
	accountService service.AccountService
}

func New(u service.UserService, n service.NoteService, c service.CardService, b service.BinService, a service.AccountService) *Handler {
	return &Handler{
		userService:    u,
		noteService:    n,
		cardService:    c,
		binService:     b,
		accountService: a,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	// Create new chi router
	r := chi.NewRouter()

	// Using built-in middleware
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentEncoding("gzip"))
	r.Use(middleware.Compress(5))
	r.Use(middleware.Recoverer)

	// Public routes
	r.Group(func(r chi.Router) {
		r.Post("/api/user/register", h.RegisterHandler)
		r.Post("/api/user/login", h.LoginHandler)
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(Authenticator)
		r.Get("/api/user/sync", h.SyncHandler)
		r.Post("/api/user/create", h.CreateHandler)
		r.Post("/api/user/update", h.UpdateHandler)
		r.Post("/api/user/delete", h.DeleteHandler)
	})
	return r
}
