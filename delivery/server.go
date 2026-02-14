package delivery

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/HosseinForouzan/E-Commerce-API/config"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice/authservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config config.Config
	authSvc authservice.Service
	userSvc userservice.Service
}

func New(config config.Config ,authSvc authservice.Service, userSvc userservice.Service) Server {
	return Server{
		config:config,
		 authSvc: authSvc,
		  userSvc:  userSvc}
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
	e.POST("/profile", s.UserProfile)

	// Start server
	if err := e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)

  }
}