package delivery

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/HosseinForouzan/E-Commerce-API/config"
	mw "github.com/HosseinForouzan/E-Commerce-API/delivery/middleware"
	"github.com/HosseinForouzan/E-Commerce-API/service/authservice"
	"github.com/HosseinForouzan/E-Commerce-API/service/productservice"
	"github.com/HosseinForouzan/E-Commerce-API/service/userservice"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config config.Config
	authSvc authservice.Service
	userSvc userservice.Service
	productSvc productservice.Service
}

func New(config config.Config ,authSvc authservice.Service, userSvc userservice.Service, productSvc productservice.Service) Server {
	return Server{
		config:config,
		 authSvc: authSvc,
		  userSvc:  userSvc,
		  productSvc: productSvc,
		
		}
}

func (s Server) SetRoutes() {
   e := echo.New()

	// Middleware
	e.Use(middleware.RequestLogger()) // use the default RequestLogger middleware with slog logger
	e.Use(middleware.Recover()) // recover panics as errors for proper error handling

	// Routes
	e.GET("/healthcheck", s.HealthCheck, mw.RequireRole("admin"))
	e.POST("/register", s.UserRegister)
	e.POST("/login", s.UserLogin)
	e.GET("/profile", s.UserProfile, mw.Auth(s.authSvc, s.config.Auth))

	e.GET("/products", s.ListProducts)
	e.GET("/products/:id", s.GetProductByID)






	// Start server
	if err := e.Start(fmt.Sprintf(":%d", s.config.HttpServer.Port)); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)

  }
}