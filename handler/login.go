package handler

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/ricardo-ronchini/budget-flow-app-go/contexts"
	"github.com/ricardo-ronchini/budget-flow-app-go/service"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("sua-chave-secreta-super-segura") // ideal usar env

type LoginCredentials struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var V1Login = &contexts.WebRoute{
	Path:   "/login",
	Method: contexts.POST,
	Handler: func(c *contexts.Context) (int, any) {
		c.Logs().Info("[LOGIN] initiating a request to access the services")

		var creds LoginCredentials

		if err := c.EchoContext.Bind(&creds); err != nil {
			return http.StatusBadRequest, echo.Map{
				"erro": "invalid JSON format",
			}
		}

		userData := service.LoginUser{
			UserName: &creds.UserName,
			Password: &creds.Password,
		}

		db := c.Database().Connect()
		defer db.Close()

		user, err := userData.GetUserForLogin(c, db)
		if err != nil {
			return http.StatusUnauthorized, echo.Map{
				"error":   "unauthorized access",
				"message": err.Error(),
			}
		}

		if user.PasswordHash == nil {
			return http.StatusInternalServerError, echo.Map{
				"error": "user validation was not possible",
			}
		}

		if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(creds.Password)); err != nil {
			return http.StatusUnauthorized, echo.Map{
				"error": "invalid credentials",
			}
		}

		// Criação do token JWT | mock
		claims := jwt.MapClaims{
			"username": *user.UserName,
			"email":    *user.Email,
			"exp":      time.Now().Add(2 * time.Hour).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenStr, err := token.SignedString(jwtSecret)
		if err != nil {
			return http.StatusInternalServerError, echo.Map{
				"erro": "Erro ao gerar o token",
			}
		}

		c.Logs().Logger.Debug("token str: ", tokenStr)

		resp := echo.Map{
			"token": tokenStr,
		}

		return http.StatusOK, c.API().Ok(resp)
	},
	Authenticate: false,
}
