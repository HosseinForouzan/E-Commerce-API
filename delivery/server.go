package delivery

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/HosseinForouzan/E-Commerce-API/service/userservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	UserSvc userservice.Service
}

func New(userSvc userservice.Service) Server {
	return Server{UserSvc: userSvc}
}

func (s Server) SetRoutes() {
   e := echo.New()



	// Middleware
	e.Use(middleware.RequestLogger()) // use the default RequestLogger middleware with slog logger
	e.Use(middleware.Recover()) // recover panics as errors for proper error handling

	// Routes
	e.GET("/healthcheck", s.HealthCheck)
	e.POST("register", s.UserRegister)
	e.POST("/login", s.UserLogin)

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)

  }
}