package httpserver

import (
	"garasystem/internal/adapters/httpserver/handler/auth"
	"garasystem/internal/adapters/httpserver/handler/user"
	"garasystem/internal/adapters/httpserver/handler/vehicleorder"
	"garasystem/internal/core/ports"
	"garasystem/pkg/config"
	"garasystem/pkg/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Server struct {
	config         *config.Config
	Router         *echo.Echo
	userHandler    ports.UserHandler
	authHandler    ports.AuthHandler
	vehicleHandler ports.VehicleOrderHandler
}

func NewServer(
	config *config.Config,
	userService ports.UserService,
	vehicleService ports.VehicleOrderService,
	snsService ports.SNSService,
) (*Server, error) {
	s := &Server{
		config:         config,
		Router:         echo.New(),
		userHandler:    user.NewHandler(userService),
		authHandler:    auth.NewHandler(config, userService, snsService),
		vehicleHandler: vehicleorder.NewHandler(vehicleService),
	}

	// Middleware
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())
	s.Router.Use(jwt.NewAuthMiddleware(config.JwtSecretKey))

	s.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	apiGroup := s.Router.Group("api")

	userGroup := apiGroup.Group("/users")
	userGroup.PUT("", s.userHandler.Update)

	authGroup := apiGroup.Group("/auths")
	authGroup.POST("/register", s.authHandler.Register)
	authGroup.POST("/verify", s.authHandler.Verify)
	authGroup.POST("/resend-code", s.authHandler.ResendCode)
	authGroup.POST("/login", s.authHandler.Login)

	vehicleGroup := apiGroup.Group("/vehicle-orders")
	vehicleGroup.GET("/:id", s.vehicleHandler.GetByID)
	vehicleGroup.GET("", s.vehicleHandler.GetList)
	vehicleGroup.POST("", s.vehicleHandler.Create)
	vehicleGroup.PUT("/:id", s.vehicleHandler.Update)

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
