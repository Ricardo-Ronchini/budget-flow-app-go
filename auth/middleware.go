package auth

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ricardo-ronchini/budget-flow-app-go/common"
)

var jwtSecret = []byte("sua-chave-secreta-super-segura")

func Middleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"erro": "Token inválido ou ausente",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenMalformed
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"erro": "Token inválido ou ausente",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID, _ := claims["user_id"].(string)
			ctx := WithUserID(c.Request().Context(), userID)
			req := c.Request().WithContext(ctx)
			c.SetRequest(req)
		}

		return next(c)
	}
}

func ConfigCORS() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: common.GetEnvArray("ALLOWED_ORIGINS", []string{}),
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			"X-API-Key",
		},
		AllowCredentials: true,
		MaxAge:           86400, // 24 hours
	})
}
