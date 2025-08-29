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
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var V1Login = &contexts.WebRoute{
	Path:   "/login",
	Method: contexts.POST,
	Handler: func(c *contexts.Context) (int, any) {
		var creds LoginCredentials

		// Faz o bind do JSON para a struct
		if err := c.EchoContext.Bind(&creds); err != nil {
			return http.StatusBadRequest, echo.Map{
				"erro": "Formato do JSON inválido",
			}
		}

		userData := service.User{
			UserName: &creds.UserName,
			Email:    &creds.Email,
		}

		user, err := userData.GetUserForLogin(c, nil)
		if err != nil {
			return http.StatusUnauthorized, echo.Map{
				"error": "unauthorized access",
			}
		}

		if user.Password == nil {
			return http.StatusInternalServerError, echo.Map{
				"error": "problema na validação de usuário",
			}
		}

		if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(creds.Password)); err != nil {
			return http.StatusUnauthorized, echo.Map{
				"error": "Credenciais inválidas",
			}
		}

		// Criação do token JWT | mock
		claims := jwt.MapClaims{
			"user_name": user.UserName,
			"email":     user.Email,
			"exp":       time.Now().Add(2 * time.Hour).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		tokenStr, err := token.SignedString(jwtSecret)
		if err != nil {
			return http.StatusInternalServerError, echo.Map{
				"erro": "Erro ao gerar o token",
			}
		}

		resp := echo.Map{
			"token": tokenStr,
		}

		return http.StatusOK, c.API().Ok(resp)
	},
}
