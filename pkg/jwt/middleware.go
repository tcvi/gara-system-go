package jwt

import (
	"garasystem/internal/core/myerror"
	"garasystem/internal/logger"
	"garasystem/pkg/util"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"regexp"
)

type SkipPattern = string

var skipPatterns = []SkipPattern{
	"api/auths/.*",
}

func NewAuthMiddleware(secretKey string) echo.MiddlewareFunc {
	contextKey := "User"
	jwtConfig := echojwt.Config{
		TokenLookup: "header:Authorization:Bearer ",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(MyClaims)
		},
		SigningKey: []byte(secretKey),
		ContextKey: contextKey,
		Skipper: func(c echo.Context) bool {
			for _, pattern := range skipPatterns {
				re, err := regexp.Compile(pattern)
				if err != nil {
					logger.Log.Error(err)
					return false
				}

				if re.MatchString(c.Path()) {
					return true
				}
			}
			return false
		},
		SuccessHandler: func(c echo.Context) {
			token := c.Get(contextKey).(*jwt.Token)
			user := token.Claims.(*MyClaims)
			c.Set(util.JWT_USER_ID_CONTEXT_KEY, user.UserId)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return util.Response.Error(c, myerror.ErrAuthUnauthorized(err))
		},
	}
	return echojwt.WithConfig(jwtConfig)
}
