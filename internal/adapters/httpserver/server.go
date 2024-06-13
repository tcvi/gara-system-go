package httpserver

import (
	"garasystem/internal/adapters/httpserver/handler/auth"
	"garasystem/internal/adapters/httpserver/handler/category"
	"garasystem/internal/adapters/httpserver/handler/item"
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
	config          *config.Config
	Router          *echo.Echo
	userHandler     ports.UserHandler
	authHandler     ports.AuthHandler
	vehicleHandler  ports.VehicleOrderHandler
	categoryHandler ports.CategoryHandler
	itemHandler     ports.ItemHandler
}

func NewServer(
	config *config.Config,
	userService ports.UserService,
	vehicleService ports.VehicleOrderService,
	categoryService ports.CategoryService,
	itemService ports.ItemService,
	snsService ports.SNSService,
) (*Server, error) {
	s := &Server{
		config:          config,
		Router:          echo.New(),
		userHandler:     user.NewHandler(userService),
		authHandler:     auth.NewHandler(config, userService, snsService),
		vehicleHandler:  vehicleorder.NewHandler(vehicleService),
		categoryHandler: category.NewHandler(categoryService),
		itemHandler:     item.NewHandler(itemService),
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

	categoryGroup := apiGroup.Group("/categories")
	categoryGroup.GET("", s.categoryHandler.GetList)
	categoryGroup.POST("", s.categoryHandler.Create)
	categoryGroup.PUT("/:id", s.categoryHandler.Update)
	categoryGroup.DELETE("/:id", s.categoryHandler.Delete)

	itemGroup := apiGroup.Group("/items")
	itemGroup.GET("", s.itemHandler.GetItems)
	itemGroup.POST("", s.itemHandler.CreateItem)
	itemGroup.PUT("/:id", s.itemHandler.UpdateItem)

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
