package httpserver

import (
	"garasystem/internal/adapters/httpserver/handler/auth"
	"garasystem/internal/adapters/httpserver/handler/user"
	"garasystem/internal/core/ports"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Server struct {
	Router      *echo.Echo
	userHandler ports.UserHandler
	authHandler ports.AuthHandler
}

func NewServer(userService ports.UserService, snsService ports.SNSService) (*Server, error) {
	s := &Server{
		Router:      echo.New(),
		userHandler: user.NewHandler(userService),
		authHandler: auth.NewHandler(userService, snsService),
	}

	// Middleware
	s.Router.Use(middleware.Logger())
	s.Router.Use(middleware.Recover())

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

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
